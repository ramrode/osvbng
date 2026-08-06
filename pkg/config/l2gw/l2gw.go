// Copyright 2026 The osvbng Authors
// Licensed under the GNU General Public License v3.0 or later.
// SPDX-License-Identifier: GPL-3.0-or-later

package l2gw

import (
	"fmt"

	"github.com/veesix-networks/osvbng/pkg/config/vlan"
)

// L2GWConfig is the top-level layer 2 wholesale gateway configuration.
// Handoff groups give the OSS/RADIUS side a stable label for an egress
// port; static maps cross-connect whole S-VLANs without any per-subscriber
// signalling.
type L2GWConfig struct {
	HandoffGroups map[string]*HandoffGroup `json:"handoff-groups,omitempty" yaml:"handoff-groups,omitempty"`
	StaticMaps    []*StaticMap             `json:"static-maps,omitempty" yaml:"static-maps,omitempty"`
}

// HandoffGroup names an egress port toward a retail ISP. RADIUS returns
// the group label; osvbng resolves it to the interface and allocates
// egress VLANs from the configured ranges unless RADIUS supplies
// explicit values.
type HandoffGroup struct {
	Interface string `json:"interface" yaml:"interface"`
	VlanTpid  string `json:"vlan-tpid,omitempty" yaml:"vlan-tpid,omitempty"`

	// SVLAN pins every circuit of this group to one outer VLAN
	// (VLAN-per-ISP model). Mutually exclusive with SVLANRange.
	SVLAN uint16 `json:"svlan,omitempty" yaml:"svlan,omitempty"`

	// SVLANRange / CVLANRange bound the per-circuit egress VLAN
	// allocator (e.g. "200-299", "1-4000").
	SVLANRange string `json:"svlan-range,omitempty" yaml:"svlan-range,omitempty"`
	CVLANRange string `json:"cvlan-range,omitempty" yaml:"cvlan-range,omitempty"`
}

// StaticMap cross-connects an entire access S-VLAN (all C-VLANs,
// wildcard) to a handoff group with no DHCP trigger and no RADIUS.
type StaticMap struct {
	AccessInterface string `json:"access-interface" yaml:"access-interface"`
	SVLAN           string `json:"svlan" yaml:"svlan"`
	HandoffGroup    string `json:"handoff-group" yaml:"handoff-group"`

	// HandoffSVLAN translates the outer tag toward the ISP; 0 with
	// Transparent unset falls back to the handoff group's fixed SVLAN,
	// or (if none) transparent pass-through.
	HandoffSVLAN uint16 `json:"handoff-svlan,omitempty" yaml:"handoff-svlan,omitempty"`
	Transparent  bool   `json:"transparent,omitempty" yaml:"transparent,omitempty"`
}

func (hg *HandoffGroup) GetOuterTPID() uint16 {
	switch hg.VlanTpid {
	case "dot1q":
		return 0x8100
	case "dot1ad", "":
		return 0x88A8
	}
	return 0x88A8
}

func (sm *StaticMap) GetSVLANs() ([]uint16, error) {
	return vlan.ParseVLANRange(sm.SVLAN)
}

// Validate performs structural validation of the l2gw section.
func (c *L2GWConfig) Validate() error {
	if c == nil {
		return nil
	}
	for name, hg := range c.HandoffGroups {
		if hg == nil || hg.Interface == "" {
			return fmt.Errorf("l2gw handoff-group %q: interface is required", name)
		}
		if hg.SVLAN != 0 && hg.SVLANRange != "" {
			return fmt.Errorf("l2gw handoff-group %q: svlan and svlan-range are mutually exclusive", name)
		}
		if hg.SVLAN > 4094 {
			return fmt.Errorf("l2gw handoff-group %q: svlan %d out of range", name, hg.SVLAN)
		}
		if hg.SVLANRange != "" {
			if _, err := vlan.ParseVLANRange(hg.SVLANRange); err != nil {
				return fmt.Errorf("l2gw handoff-group %q: svlan-range: %w", name, err)
			}
		}
		if hg.CVLANRange != "" {
			if _, err := vlan.ParseVLANRange(hg.CVLANRange); err != nil {
				return fmt.Errorf("l2gw handoff-group %q: cvlan-range: %w", name, err)
			}
		}
		switch hg.VlanTpid {
		case "", "dot1ad", "dot1q":
		default:
			return fmt.Errorf("l2gw handoff-group %q: vlan-tpid must be dot1ad or dot1q", name)
		}
	}
	for i, sm := range c.StaticMaps {
		if sm == nil {
			continue
		}
		if sm.AccessInterface == "" {
			return fmt.Errorf("l2gw static-map %d: access-interface is required", i)
		}
		if sm.SVLAN == "" {
			return fmt.Errorf("l2gw static-map %d: svlan is required", i)
		}
		if _, err := vlan.ParseVLANRange(sm.SVLAN); err != nil {
			return fmt.Errorf("l2gw static-map %d: svlan: %w", i, err)
		}
		if _, ok := c.HandoffGroups[sm.HandoffGroup]; !ok {
			return fmt.Errorf("l2gw static-map %d: unknown handoff-group %q", i, sm.HandoffGroup)
		}
		if sm.Transparent && sm.HandoffSVLAN != 0 {
			return fmt.Errorf("l2gw static-map %d: transparent excludes handoff-svlan", i)
		}
	}
	return nil
}
