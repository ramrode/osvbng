// Copyright 2026 The osvbng Authors
// Licensed under the GNU General Public License v3.0 or later.
// SPDX-License-Identifier: GPL-3.0-or-later

package l2gw

import (
	"fmt"

	l2gwcfg "github.com/veesix-networks/osvbng/pkg/config/l2gw"
	"github.com/veesix-networks/osvbng/pkg/config/vlan"
)

// vlanAllocator hands out egress (svlan, cvlan) pairs for one handoff
// group from its configured ranges. Static-SVLAN groups allocate only
// the inner tag; range-SVLAN groups fill each S-VLAN's C-space before
// moving to the next.
type vlanAllocator struct {
	svlans []uint16
	cvlans []uint16
	used   map[uint32]bool // svlan<<16|cvlan
}

func pairKey(svlan, cvlan uint16) uint32 {
	return uint32(svlan)<<16 | uint32(cvlan)
}

func newVlanAllocator(hg *l2gwcfg.HandoffGroup) (*vlanAllocator, error) {
	a := &vlanAllocator{used: make(map[uint32]bool)}

	switch {
	case hg.SVLAN != 0:
		a.svlans = []uint16{hg.SVLAN}
	case hg.SVLANRange != "":
		svlans, err := vlan.ParseVLANRange(hg.SVLANRange)
		if err != nil {
			return nil, fmt.Errorf("svlan-range: %w", err)
		}
		a.svlans = svlans
	}

	if hg.CVLANRange != "" {
		cvlans, err := vlan.ParseVLANRange(hg.CVLANRange)
		if err != nil {
			return nil, fmt.Errorf("cvlan-range: %w", err)
		}
		a.cvlans = cvlans
	}

	return a, nil
}

// allocate returns the next free (svlan, cvlan) pair. cvlan is 0 when
// the group has no C-VLAN range (single-tagged handoff).
func (a *vlanAllocator) allocate() (uint16, uint16, error) {
	if len(a.svlans) == 0 {
		return 0, 0, fmt.Errorf("no egress vlan ranges configured")
	}
	for _, s := range a.svlans {
		if len(a.cvlans) == 0 {
			if !a.used[pairKey(s, 0)] {
				a.used[pairKey(s, 0)] = true
				return s, 0, nil
			}
			continue
		}
		for _, cv := range a.cvlans {
			if !a.used[pairKey(s, cv)] {
				a.used[pairKey(s, cv)] = true
				return s, cv, nil
			}
		}
	}
	return 0, 0, fmt.Errorf("egress vlan space exhausted")
}

// mark reserves an explicit pair (RADIUS override or restore).
func (a *vlanAllocator) mark(svlan, cvlan uint16) {
	a.used[pairKey(svlan, cvlan)] = true
}

func (a *vlanAllocator) free(svlan, cvlan uint16) {
	delete(a.used, pairKey(svlan, cvlan))
}

// getAllocator lazily builds the allocator for a handoff group.
func (c *Component) getAllocator(name string, hg *l2gwcfg.HandoffGroup) (*vlanAllocator, error) {
	c.allocMu.Lock()
	defer c.allocMu.Unlock()
	if a, ok := c.allocators[name]; ok {
		return a, nil
	}
	a, err := newVlanAllocator(hg)
	if err != nil {
		return nil, fmt.Errorf("handoff-group %q: %w", name, err)
	}
	c.allocators[name] = a
	return a, nil
}

// freeEgress releases a dynamic circuit's allocated egress pair.
func (c *Component) freeEgress(ct *Circuit) {
	if !ct.AllocatedEgress {
		return
	}
	c.allocMu.Lock()
	if a, ok := c.allocators[ct.HandoffGroup]; ok {
		a.free(ct.HandoffSVLAN, ct.HandoffCVLAN)
	}
	c.allocMu.Unlock()
}
