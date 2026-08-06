// Copyright 2026 The osvbng Authors
// Licensed under the GNU General Public License v3.0 or later.
// SPDX-License-Identifier: GPL-3.0-or-later

package l2gw

import (
	"encoding/binary"
	"testing"

	"github.com/google/gopacket/layers"
	"github.com/veesix-networks/osvbng/pkg/dataplane"
	"github.com/veesix-networks/osvbng/pkg/dhcp6"
)

func appendV6Option(buf []byte, code uint16, data []byte) []byte {
	var hdr [4]byte
	binary.BigEndian.PutUint16(hdr[0:2], code)
	binary.BigEndian.PutUint16(hdr[2:4], uint16(len(data)))
	buf = append(buf, hdr[:]...)
	return append(buf, data...)
}

func buildRelayForwardSolicit(interfaceID, remoteID []byte) []byte {
	inner := []byte{byte(dhcp6.MsgTypeSolicit), 0x01, 0x02, 0x03}

	relay := make([]byte, 0, 128)
	relay = append(relay, byte(dhcp6.MsgTypeRelayForward))
	relay = append(relay, 0)
	relay = append(relay, make([]byte, 32)...)
	relay = appendV6Option(relay, dhcp6.OptInterfaceID, interfaceID)
	remoteOpt := append([]byte{0, 0, 0x0B, 0xBC}, remoteID...)
	relay = appendV6Option(relay, dhcp6.OptRemoteID, remoteOpt)
	relay = appendV6Option(relay, dhcp6.OptRelayMsg, inner)
	return relay
}

func TestDHCPv6TriggerIdentityDirect(t *testing.T) {
	pkt := &dataplane.ParsedPacket{
		DHCPv6: &layers.DHCPv6{MsgType: layers.DHCPv6MsgTypeSolicit},
	}
	msgType, info := dhcpv6TriggerIdentity(pkt)
	if msgType != layers.DHCPv6MsgTypeSolicit {
		t.Fatalf("direct solicit: got %v", msgType)
	}
	if info != nil {
		t.Fatal("direct solicit must have no relay info")
	}
}

func TestDHCPv6TriggerIdentityRelayed(t *testing.T) {
	raw := buildRelayForwardSolicit([]byte("olt1 0/1/7:100.10"), []byte("SUB-42"))
	d := &layers.DHCPv6{MsgType: layers.DHCPv6MsgTypeRelayForward}
	d.Contents = raw

	pkt := &dataplane.ParsedPacket{DHCPv6: d}
	msgType, info := dhcpv6TriggerIdentity(pkt)
	if msgType != layers.DHCPv6MsgTypeSolicit {
		t.Fatalf("relayed solicit: got %v", msgType)
	}
	if info == nil {
		t.Fatal("expected relay info")
	}
	if string(info.InterfaceID) != "olt1 0/1/7:100.10" {
		t.Fatalf("interface-id: got %q", info.InterfaceID)
	}
	if string(info.RemoteID) != "SUB-42" {
		t.Fatalf("remote-id (enterprise prefix stripped): got %q", info.RemoteID)
	}
}
