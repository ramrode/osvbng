// Copyright 2026 The osvbng Authors
// Licensed under the GNU General Public License v3.0 or later.
// SPDX-License-Identifier: GPL-3.0-or-later

package l2gw

import (
	"context"
	"time"

	hapb "github.com/veesix-networks/osvbng/api/proto/ha"
	"github.com/veesix-networks/osvbng/pkg/events"
	"github.com/veesix-networks/osvbng/pkg/ha"
	"github.com/veesix-networks/osvbng/pkg/models"
	"github.com/veesix-networks/osvbng/pkg/opdb"
	"google.golang.org/protobuf/proto"
)

// restoreSyncedStandby re-installs peer-synced circuits (disabled) after
// a standby restart: the HA synced namespace survives in opdb while the
// component's own checkpoints only cover circuits this node owned.
func (c *Component) restoreSyncedStandby(ctx context.Context) error {
	if c.opdb == nil || c.srgMgr == nil {
		return nil
	}
	return c.opdb.Load(ctx, opdb.NamespaceHASyncedL2GW, func(key string, value []byte) error {
		cp := &hapb.SessionCheckpoint{}
		if err := proto.Unmarshal(value, cp); err != nil {
			c.logger.Warn("Corrupt synced l2gw checkpoint; dropping", "key", key, "error", err)
			_ = c.opdb.Delete(ctx, opdb.NamespaceHASyncedL2GW, key)
			return nil
		}
		c.ApplySyncedCircuit(hapb.SyncAction_SYNC_ACTION_UPDATE, cp)
		return nil
	})
}

// ForEachSession feeds the HA BulkSync stream: every dynamic circuit
// with an SRG binding is a syncable session.
func (c *Component) ForEachSession(fn func(models.SubscriberSession) bool) {
	c.circuits.Range(func(_, v any) bool {
		ct := v.(*Circuit)
		ct.mu.Lock()
		static, srg := ct.Static, ct.SRGName
		ct.mu.Unlock()
		if static || srg == "" {
			return true
		}
		return fn(ct.buildSessionModel(models.SessionStateActive))
	})
}

// ApplySyncedCircuit eagerly installs (or removes) a peer-synced circuit
// on the standby with forwarding disabled, so promotion is a batch
// state flip rather than a re-install storm. Interface indexes are
// re-resolved locally by name, they never match across peers.
func (c *Component) ApplySyncedCircuit(action hapb.SyncAction, cp *hapb.SessionCheckpoint) {
	if cp == nil || cp.AccessType != "l2gw" {
		return
	}
	if c.srgMgr != nil && cp.SrgName != "" && c.srgMgr.IsActive(cp.SrgName) {
		return
	}

	accessIdx, ok := c.ifMgr.GetSwIfIndex(cp.AccessInterface)
	if !ok {
		c.logger.Warn("Synced l2gw circuit references unknown access interface",
			"session_id", cp.SessionId, "interface", cp.AccessInterface)
		return
	}
	key := circuitKey(accessIdx, uint16(cp.OuterVlan), uint16(cp.InnerVlan))

	if action == hapb.SyncAction_SYNC_ACTION_DELETE {
		if val, exists := c.circuits.Load(key); exists {
			ct := val.(*Circuit)
			c.removeCircuit(ct)
			c.circuits.Delete(key)
			if ct.SessionID != "" {
				c.sessionIndex.Delete(ct.SessionID)
			}
		}
		return
	}

	handoffIdx, ok := c.ifMgr.GetSwIfIndex(cp.HandoffInterface)
	if !ok {
		c.logger.Warn("Synced l2gw circuit references unknown handoff interface",
			"session_id", cp.SessionId, "interface", cp.HandoffInterface)
		return
	}

	if val, exists := c.circuits.Load(key); exists {
		existing := val.(*Circuit)
		existing.mu.Lock()
		same := existing.SessionID == cp.SessionId
		existing.mu.Unlock()
		if same {
			return
		}
		c.removeCircuit(existing)
		c.circuits.Delete(key)
		if existing.SessionID != "" {
			c.sessionIndex.Delete(existing.SessionID)
		}
	}

	ct := &Circuit{
		SessionID:        cp.SessionId,
		AcctSessionID:    cp.AaaSessionId,
		Username:         cp.Username,
		MAC:              string(cp.Mac),
		AccessInterface:  cp.AccessInterface,
		AccessIfIndex:    accessIdx,
		AccessSVLAN:      uint16(cp.OuterVlan),
		AccessCVLAN:      uint16(cp.InnerVlan),
		AccessTPID:       uint16(cp.AccessTpid),
		HandoffGroup:     cp.HandoffGroup,
		HandoffInterface: cp.HandoffInterface,
		HandoffIfIndex:   handoffIdx,
		HandoffSVLAN:     uint16(cp.HandoffSvlan),
		HandoffCVLAN:     uint16(cp.HandoffCvlan),
		HandoffTPID:      uint16(cp.HandoffTpid),
		Transparent:      cp.Transparent,
		SRGName:          cp.SrgName,
		Standby:          true,
		Attributes:       cp.AaaAttributes,
		Protocol:         string(models.ProtocolDHCPv4),
	}
	if cp.BoundAtNs > 0 {
		ct.CreatedAt = time.Unix(0, cp.BoundAtNs)
	}

	if err := c.armPort(accessIdx, cp.AccessInterface); err != nil {
		c.logger.Warn("Failed to arm access port for synced circuit", "error", err)
	}
	if err := c.armPort(handoffIdx, cp.HandoffInterface); err != nil {
		c.logger.Warn("Failed to arm handoff port for synced circuit", "error", err)
	}

	if err := c.installCircuit(ct); err != nil {
		c.logger.Error("Failed to install synced l2gw circuit",
			"session_id", cp.SessionId, "error", err)
		return
	}

	c.circuits.Store(key, ct)
	if ct.SessionID != "" {
		c.sessionIndex.Store(ct.SessionID, ct)
	}

	c.logger.Debug("Installed synced l2gw circuit (standby)",
		"session_id", ct.SessionID, "mac", ct.MAC,
		"svlan", ct.AccessSVLAN, "cvlan", ct.AccessCVLAN)
}

func (c *Component) handleHAStateChange(event events.Event) {
	data, ok := event.Data.(events.HAStateChangeEvent)
	if !ok {
		return
	}

	wasActive := data.OldState == string(ha.SRGStateActive) || data.OldState == string(ha.SRGStateActiveSolo)
	isActive := data.NewState == string(ha.SRGStateActive) || data.NewState == string(ha.SRGStateActiveSolo)
	isStandby := data.NewState == string(ha.SRGStateStandby) || data.NewState == string(ha.SRGStateStandbyAlone)

	switch {
	case isActive && !wasActive:
		c.promoteSRGCircuits(data.SRGName)
	case isStandby && wasActive:
		c.demoteSRGCircuits(data.SRGName)
	}
}

// promoteSRGCircuits batch-enables the SRG's standby circuits and hands
// their accounting over to this node via the Restored topic (no
// duplicate Accounting-Start).
func (c *Component) promoteSRGCircuits(srgName string) {
	cfg, _ := c.cfgMgr.GetRunning()

	var enabled int
	c.circuits.Range(func(_, v any) bool {
		ct := v.(*Circuit)
		ct.mu.Lock()
		match := ct.SRGName == srgName && ct.Standby
		ct.mu.Unlock()
		if !match {
			return true
		}

		if err := c.vpp.SetL2GWCircuitState(ct.CircuitID, true); err != nil {
			c.logger.Error("Failed to enable l2gw circuit on promotion",
				"circuit_id", ct.CircuitID, "error", err)
			return true
		}

		ct.mu.Lock()
		ct.Standby = false
		ct.State = circuitStateInstalled
		ct.mu.Unlock()

		if cfg != nil && cfg.L2GW != nil {
			if hg, ok := cfg.L2GW.HandoffGroups[ct.HandoffGroup]; ok {
				if alloc, err := c.getAllocator(ct.HandoffGroup, hg); err == nil {
					c.allocMu.Lock()
					alloc.mark(ct.HandoffSVLAN, ct.HandoffCVLAN)
					c.allocMu.Unlock()
				}
			}
		}

		c.checkpointCircuit(ct)

		c.eventBus.Publish(events.TopicSessionRestored, events.Event{
			Source: c.Name(),
			Data: &events.SessionRestoredEvent{
				AccessType: models.AccessTypeL2GW,
				Protocol:   models.Protocol(ct.Protocol),
				SessionID:  ct.SessionID,
				Session:    ct.buildSessionModel(models.SessionStateActive),
			},
		})
		enabled++
		return true
	})

	if enabled > 0 {
		c.logger.Info("Enabled l2gw circuits on SRG promotion",
			"srg", srgName, "circuits", enabled)
	}
}

// demoteSRGCircuits disables forwarding for the SRG's circuits; the
// peer's promotion owns the traffic from here. Session state is kept so
// a re-promotion is another batch flip.
func (c *Component) demoteSRGCircuits(srgName string) {
	var disabled int
	c.circuits.Range(func(_, v any) bool {
		ct := v.(*Circuit)
		ct.mu.Lock()
		match := ct.SRGName == srgName && !ct.Standby && !ct.Static
		ct.mu.Unlock()
		if !match {
			return true
		}

		if err := c.vpp.SetL2GWCircuitState(ct.CircuitID, false); err != nil {
			c.logger.Error("Failed to disable l2gw circuit on demotion",
				"circuit_id", ct.CircuitID, "error", err)
			return true
		}

		ct.mu.Lock()
		ct.Standby = true
		ct.State = circuitStateStandby
		ct.mu.Unlock()
		disabled++
		return true
	})

	if disabled > 0 {
		c.logger.Info("Disabled l2gw circuits on SRG demotion",
			"srg", srgName, "circuits", disabled)
	}
}
