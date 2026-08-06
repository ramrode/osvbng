// Copyright 2026 The osvbng Authors
// Licensed under the GNU General Public License v3.0 or later.
// SPDX-License-Identifier: GPL-3.0-or-later

package l2gw

import (
	"context"
	"encoding/json"

	"github.com/veesix-networks/osvbng/pkg/events"
	"github.com/veesix-networks/osvbng/pkg/models"
)

// restoreCircuits replays opdb-checkpointed dynamic circuits into the
// dataplane on startup. No RADIUS re-auth, no Accounting-Start (the
// Restored topic rebuilds AAA's in-memory state without emitting Start),
// and no trigger storm: 20k circuits restore as 20k idempotent
// dataplane adds.
func (c *Component) restoreCircuits(ctx context.Context) error {
	if c.opdb == nil {
		return nil
	}

	var restored, failed int
	err := c.opdb.Load(ctx, opdbNamespace, func(key string, value []byte) error {
		ct := &Circuit{}
		if err := json.Unmarshal(value, ct); err != nil {
			c.logger.Warn("Corrupt l2gw circuit checkpoint; dropping", "key", key, "error", err)
			_ = c.opdb.Delete(ctx, opdbNamespace, key)
			return nil
		}

		// Interface indexes may have drifted across a VPP restart,
		// re-resolve by name before touching the dataplane.
		if idx, ok := c.ifMgr.GetSwIfIndex(ct.AccessInterface); ok {
			ct.AccessIfIndex = idx
		}
		if idx, ok := c.ifMgr.GetSwIfIndex(ct.HandoffInterface); ok {
			ct.HandoffIfIndex = idx
		}

		// A restart while this node's SRG is standby must not resume
		// forwarding; promotion batch-enables.
		if c.srgMgr != nil && ct.SRGName != "" && !c.srgMgr.IsActive(ct.SRGName) {
			ct.Standby = true
		}

		if err := c.armPort(ct.AccessIfIndex, ct.AccessInterface); err != nil {
			c.logger.Warn("Failed to arm access port during restore", "key", key, "error", err)
		}
		if err := c.armPort(ct.HandoffIfIndex, ct.HandoffInterface); err != nil {
			c.logger.Warn("Failed to arm handoff port during restore", "key", key, "error", err)
		}

		if err := c.installCircuit(ct); err != nil {
			c.logger.Error("Failed to restore l2gw circuit", "key", key, "error", err)
			failed++
			return nil
		}

		if ct.AllocatedEgress {
			cfg, _ := c.cfgMgr.GetRunning()
			if cfg != nil && cfg.L2GW != nil {
				if hg, ok := cfg.L2GW.HandoffGroups[ct.HandoffGroup]; ok {
					if alloc, aErr := c.getAllocator(ct.HandoffGroup, hg); aErr == nil {
						c.allocMu.Lock()
						alloc.mark(ct.HandoffSVLAN, ct.HandoffCVLAN)
						c.allocMu.Unlock()
					}
				}
			}
		}

		c.circuits.Store(ct.key(), ct)
		if ct.SessionID != "" {
			c.sessionIndex.Store(ct.SessionID, ct)
		}
		restored++

		// Refresh the checkpoint with re-resolved indexes.
		c.checkpointCircuit(ct)

		// Standby circuits carry no accounting here, the active peer
		// accounts; promotion publishes Restored when this node takes
		// over.
		if !ct.Standby {
			c.eventBus.Publish(events.TopicSessionRestored, events.Event{
				Source: c.Name(),
				Data: &events.SessionRestoredEvent{
					AccessType: models.AccessTypeL2GW,
					Protocol:   models.Protocol(ct.Protocol),
					SessionID:  ct.SessionID,
					Session:    ct.buildSessionModel(models.SessionStateActive),
				},
			})
		}
		return nil
	})

	if restored > 0 || failed > 0 {
		c.logger.Info("l2gw circuit restore complete", "restored", restored, "failed", failed)
	}
	return err
}
