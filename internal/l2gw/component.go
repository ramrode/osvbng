// Copyright 2026 The osvbng Authors
// Licensed under the GNU General Public License v3.0 or later.
// SPDX-License-Identifier: GPL-3.0-or-later

// Package l2gw implements the layer 2 wholesale gateway control plane:
// DHCP-triggered, AAA-authorized cross-connects of subscriber circuits
// between access networks and retail ISP handoff ports. osvbng never
// terminates DHCP or L3 for these subscribers — the retail ISP's BNG
// does; osvbng owns circuit steering, egress VLAN allocation, and
// wholesale accounting.
package l2gw

import (
	"context"
	"sync"

	"github.com/veesix-networks/osvbng/pkg/component"
	"github.com/veesix-networks/osvbng/pkg/dataplane"
	"github.com/veesix-networks/osvbng/pkg/events"
	"github.com/veesix-networks/osvbng/pkg/ha"
	"github.com/veesix-networks/osvbng/pkg/ifmgr"
	"github.com/veesix-networks/osvbng/pkg/logger"
	"github.com/veesix-networks/osvbng/pkg/opdb"
	"github.com/veesix-networks/osvbng/pkg/southbound"
)

const opdbNamespace = "l2gw:circuits"

type Component struct {
	*component.Base

	logger   *logger.Logger
	eventBus events.Bus
	cfgMgr   component.ConfigManager
	vpp      southbound.Southbound
	ifMgr    *ifmgr.Manager
	opdb     opdb.Store
	srgMgr   ha.SRGProvider

	// circuits keyed by access tuple "port:svlan:cvlan"; sessionIndex
	// keyed by session ID for terminate/accounting correlation.
	circuits     sync.Map
	sessionIndex sync.Map

	allocMu    sync.Mutex
	allocators map[string]*vlanAllocator

	armedMu    sync.Mutex
	armedPorts map[uint32]bool

	triggerChan chan *dataplane.ParsedPacket

	aaaSub       events.Subscription
	terminateSub events.Subscription
}

func New(deps component.Dependencies, srgMgr ha.SRGProvider, ifMgr *ifmgr.Manager) (*Component, error) {
	c := &Component{
		Base:        component.NewBase("l2gw"),
		logger:      logger.Get(logger.L2GW),
		eventBus:    deps.EventBus,
		cfgMgr:      deps.ConfigManager,
		vpp:         deps.Southbound,
		ifMgr:       ifMgr,
		opdb:        deps.OpDB,
		srgMgr:      srgMgr,
		allocators:  make(map[string]*vlanAllocator),
		armedPorts:  make(map[uint32]bool),
		triggerChan: make(chan *dataplane.ParsedPacket, 1024),
	}
	return c, nil
}

// TriggerChan is the write side handed to the IPoE component's DHCP
// dispatch: packets on subscriber groups with access-type l2gw are
// forwarded here instead of being terminated.
func (c *Component) TriggerChan() chan<- *dataplane.ParsedPacket {
	return c.triggerChan
}

func (c *Component) Start(ctx context.Context) error {
	c.StartContext(ctx)
	c.logger.Info("Starting L2GW component")

	c.SetReadyState(component.StateRestoring)

	if err := c.restoreCircuits(ctx); err != nil {
		c.logger.Warn("Failed to restore l2gw circuits from OpDB", "error", err)
	}

	if err := c.applyStaticMaps(); err != nil {
		c.logger.Error("Failed to apply l2gw static maps", "error", err)
	}

	c.aaaSub = c.eventBus.Subscribe(events.TopicAAAResponseL2GW, c.handleAAAResponse)
	c.terminateSub = c.eventBus.Subscribe(events.TopicSubscriberTerminate, c.handleSubscriberTerminate)

	c.Go(c.consumeTriggers)
	c.Go(c.janitor)

	c.SetReadyState(component.StateReady)
	c.eventBus.Publish(events.TopicComponentReady, events.Event{
		Source: c.Name(),
		Data:   &events.ComponentReadyEvent{Component: c.Name(), State: c.ReadyState().String()},
	})

	return nil
}

func (c *Component) Stop(ctx context.Context) error {
	c.logger.Info("Stopping L2GW component")
	c.SetReadyState(component.StateDraining)

	if c.aaaSub != nil {
		c.aaaSub.Unsubscribe()
	}
	if c.terminateSub != nil {
		c.terminateSub.Unsubscribe()
	}

	c.StopContext()
	return nil
}

// armPort enables the l2gw-input feature on a port exactly once.
func (c *Component) armPort(swIfIndex uint32, name string) error {
	c.armedMu.Lock()
	defer c.armedMu.Unlock()
	if c.armedPorts[swIfIndex] {
		return nil
	}
	if err := c.vpp.L2GWEnableInput(name, true); err != nil {
		return err
	}
	c.armedPorts[swIfIndex] = true
	c.logger.Info("Armed l2gw on port", "interface", name, "sw_if_index", swIfIndex)
	return nil
}

// resolvePort maps a (possibly sub-)interface index to its parent port
// index and name — the l2gw plugin keys circuits on the port, tags live
// in the packet.
func (c *Component) resolvePort(swIfIndex uint32) (uint32, string) {
	if c.ifMgr == nil {
		return swIfIndex, ""
	}
	iface := c.ifMgr.Get(swIfIndex)
	if iface == nil {
		return swIfIndex, ""
	}
	portIdx := swIfIndex
	if iface.HasParent() {
		portIdx = iface.SupSwIfIndex
	}
	name := ""
	if parent := c.ifMgr.Get(portIdx); parent != nil {
		name = parent.Name
	}
	return portIdx, name
}
