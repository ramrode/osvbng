// Copyright 2026 The osvbng Authors
// Licensed under the GNU General Public License v3.0 or later.
// SPDX-License-Identifier: GPL-3.0-or-later

package l2gw

import (
	"fmt"

	l2gwcfg "github.com/veesix-networks/osvbng/pkg/config/l2gw"
	"github.com/veesix-networks/osvbng/pkg/southbound"
)

// applyStaticMaps installs config-driven wildcard circuits: whole access
// S-VLANs cross-connected to a handoff group with no trigger, no RADIUS,
// no per-subscriber control-plane work.
func (c *Component) applyStaticMaps() error {
	cfg, err := c.cfgMgr.GetRunning()
	if err != nil || cfg == nil {
		return nil
	}
	return c.applyStaticMapsFrom(cfg.L2GW)
}

func (c *Component) applyStaticMapsFrom(l2gwCfg *l2gwcfg.L2GWConfig) error {
	if l2gwCfg == nil {
		return nil
	}

	for i, sm := range l2gwCfg.StaticMaps {
		if sm == nil {
			continue
		}
		if err := c.applyStaticMap(l2gwCfg, sm); err != nil {
			c.logger.Error("Failed to apply l2gw static map",
				"index", i, "access_interface", sm.AccessInterface,
				"svlan", sm.SVLAN, "handoff_group", sm.HandoffGroup,
				"error", err)
		}
	}
	return nil
}

func (c *Component) applyStaticMap(l2gwCfg *l2gwcfg.L2GWConfig, sm *l2gwcfg.StaticMap) error {
	hg, ok := l2gwCfg.HandoffGroups[sm.HandoffGroup]
	if !ok {
		return fmt.Errorf("unknown handoff-group %q", sm.HandoffGroup)
	}

	accessIdx, ok := c.ifMgr.GetSwIfIndex(sm.AccessInterface)
	if !ok {
		return fmt.Errorf("access interface %q not found", sm.AccessInterface)
	}
	handoffIdx, ok := c.ifMgr.GetSwIfIndex(hg.Interface)
	if !ok {
		return fmt.Errorf("handoff interface %q not found", hg.Interface)
	}

	svlans, err := sm.GetSVLANs()
	if err != nil {
		return fmt.Errorf("svlan: %w", err)
	}

	// A fixed egress S-VLAN cannot fan in more than one access S-VLAN:
	// the reverse-direction wildcard key would collide.
	transparent := sm.Transparent
	handoffSVLAN := sm.HandoffSVLAN
	if !transparent && handoffSVLAN == 0 {
		handoffSVLAN = hg.SVLAN
	}
	if !transparent && handoffSVLAN == 0 {
		transparent = true
	}
	if !transparent && len(svlans) > 1 {
		return fmt.Errorf("svlan range with a fixed handoff-svlan would collide on the handoff side; use transparent")
	}

	if err := c.armPort(accessIdx, sm.AccessInterface); err != nil {
		return fmt.Errorf("arm access port: %w", err)
	}
	if err := c.armPort(handoffIdx, hg.Interface); err != nil {
		return fmt.Errorf("arm handoff port: %w", err)
	}

	for _, svlan := range svlans {
		hs := handoffSVLAN
		if transparent {
			hs = svlan
		}

		key := circuitKey(accessIdx, svlan, southbound.L2GWCvlanAny)
		if _, exists := c.circuits.Load(key); exists {
			continue
		}

		ct := &Circuit{
			AccessInterface:  sm.AccessInterface,
			AccessIfIndex:    accessIdx,
			AccessSVLAN:      svlan,
			AccessCVLAN:      southbound.L2GWCvlanAny,
			HandoffGroup:     sm.HandoffGroup,
			HandoffInterface: hg.Interface,
			HandoffIfIndex:   handoffIdx,
			HandoffSVLAN:     hs,
			HandoffCVLAN:     southbound.L2GWCvlanAny,
			HandoffTPID:      hg.GetOuterTPID(),
			Transparent:      transparent,
			Static:           true,
			State:            circuitStateAuthenticating,
		}

		if err := c.installCircuit(ct); err != nil {
			c.logger.Error("Failed to install static l2gw circuit",
				"access_interface", sm.AccessInterface, "svlan", svlan,
				"handoff_group", sm.HandoffGroup, "error", err)
			continue
		}
		c.circuits.Store(key, ct)
		c.logger.Info("Installed static l2gw circuit",
			"access_interface", sm.AccessInterface, "svlan", svlan,
			"handoff_group", sm.HandoffGroup, "handoff_svlan", hs,
			"transparent", transparent, "circuit_id", ct.CircuitID)
	}

	return nil
}
