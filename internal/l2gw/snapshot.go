// Copyright 2026 The osvbng Authors
// Licensed under the GNU General Public License v3.0 or later.
// SPDX-License-Identifier: GPL-3.0-or-later

package l2gw

import (
	"sort"
	"time"
)

// CircuitSummary is the show/API projection of one circuit.
type CircuitSummary struct {
	SessionID       string    `json:"session_id,omitempty"`
	MAC             string    `json:"mac,omitempty"`
	Username        string    `json:"username,omitempty"`
	AccessInterface string    `json:"access_interface"`
	AccessSVLAN     uint16    `json:"access_svlan"`
	AccessCVLAN     uint16    `json:"access_cvlan,omitempty"`
	AccessCVLANAny  bool      `json:"access_cvlan_any,omitempty"`
	HandoffGroup    string    `json:"handoff_group"`
	HandoffSVLAN    uint16    `json:"handoff_svlan,omitempty"`
	HandoffCVLAN    uint16    `json:"handoff_cvlan,omitempty"`
	Transparent     bool      `json:"transparent,omitempty"`
	Static          bool      `json:"static,omitempty"`
	State           string    `json:"state"`
	CircuitID       uint32    `json:"circuit_id"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
}

// SnapshotCircuits returns all circuits sorted by access tuple.
func (c *Component) SnapshotCircuits() []CircuitSummary {
	var out []CircuitSummary
	c.circuits.Range(func(_, v any) bool {
		ct := v.(*Circuit)
		ct.mu.Lock()
		s := CircuitSummary{
			SessionID:       ct.SessionID,
			MAC:             ct.MAC,
			Username:        ct.Username,
			AccessInterface: ct.AccessInterface,
			AccessSVLAN:     ct.AccessSVLAN,
			HandoffGroup:    ct.HandoffGroup,
			HandoffSVLAN:    ct.HandoffSVLAN,
			HandoffCVLAN:    ct.HandoffCVLAN,
			Transparent:     ct.Transparent,
			Static:          ct.Static,
			State:           ct.State,
			CircuitID:       ct.CircuitID,
			CreatedAt:       ct.CreatedAt,
		}
		if ct.AccessCVLAN == 0xFFFF {
			s.AccessCVLANAny = true
		} else {
			s.AccessCVLAN = ct.AccessCVLAN
		}
		ct.mu.Unlock()
		out = append(out, s)
		return true
	})
	sort.Slice(out, func(i, j int) bool {
		if out[i].AccessInterface != out[j].AccessInterface {
			return out[i].AccessInterface < out[j].AccessInterface
		}
		if out[i].AccessSVLAN != out[j].AccessSVLAN {
			return out[i].AccessSVLAN < out[j].AccessSVLAN
		}
		return out[i].AccessCVLAN < out[j].AccessCVLAN
	})
	return out
}
