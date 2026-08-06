// Copyright 2026 The osvbng Authors
// Licensed under the GNU General Public License v3.0 or later.
// SPDX-License-Identifier: GPL-3.0-or-later

package southbound

// L2GWCvlanAny selects the per-S-VLAN wildcard on a circuit side: the
// circuit matches any (or no) inner VLAN and inner tags pass through
// untouched.
const L2GWCvlanAny = 0xFFFF

// L2GWCircuit describes one bidirectional wholesale circuit between an
// access port and a handoff port. VLAN values are the on-wire tags on
// each side; TPIDs are the outer TPID emitted toward that side (0 =
// dot1ad default).
type L2GWCircuit struct {
	AccessIfIndex  uint32
	AccessSVLAN    uint16
	AccessCVLAN    uint16
	AccessTPID     uint16
	HandoffIfIndex uint32
	HandoffSVLAN   uint16
	HandoffCVLAN   uint16
	HandoffTPID    uint16
	Transparent    bool
	Enabled        bool
}

type L2GW interface {
	L2GWEnableInput(ifaceName string, enable bool) error
	AddL2GWCircuit(circuit L2GWCircuit) (uint32, error)
	DelL2GWCircuit(circuit L2GWCircuit) error
	SetL2GWCircuitState(circuitID uint32, enabled bool) error
	DumpL2GWCircuits() ([]L2GWCircuitDetails, error)
}

// L2GWCircuitDetails is one circuit from the dataplane dump. The entry
// indices are the per-direction counter indices in the /osvbng/l2gw
// stats segment.
type L2GWCircuitDetails struct {
	CircuitID         uint32
	AccessEntryIndex  uint32
	HandoffEntryIndex uint32
	L2GWCircuit
}
