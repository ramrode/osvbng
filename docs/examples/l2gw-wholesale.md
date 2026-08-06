# L2 Wholesale Aggregation (L2GW)

osvbng as a wholesale aggregator between multiple access networks and retail
ISPs: subscribers' QinQ circuits are cross-connected at layer 2 to the ISP
that ordered them, DHCP-triggered and RADIUS-driven, while the retail ISP's
own BNG terminates DHCP, addressing, and policy.

```
                       +--------------------------+
  access network A ----+ eth1                bond1+---- ISP blue (BNG)
  (S 100-199, C per    |        osvbng            |     dynamic circuits
   subscriber)         |        (l2gw)        eth3+---- ISP green (BNG)
  access network B ----+ eth2                     |     static S-VLAN map
  (S 10-99)            +--------------------------+
```

## Scenario

- **ISP green** bought access network B wholesale: every S-VLAN 10 to 99 on
  `eth2` belongs to them. A static map passes the whole range transparently,
  with no per-subscriber signalling.
- **ISP blue** sells on access network A, where circuits are per-subscriber
  (one S-VLAN per OLT, one C-VLAN per subscriber). Each new subscriber's
  first DHCP packet triggers RADIUS. The wholesaler's OSS answers with the
  `isp-blue` label, or with explicit egress VLANs from its inventory, and
  osvbng splices the circuit through to `bond1`.

The exit interface is always resolved from the handoff group label, so the
RADIUS/OSS integration never references interface names. See
[L2GW configuration](../configuration/l2gw.md) for the resolution order.

## Configuration

```yaml
interfaces:
  eth1: {}
  eth2: {}
  eth3: {}
  bond1:
    bond:
      mode: lacp
      members: [eth4, eth5]

l2gw:
  handoff-groups:
    isp-blue:
      interface: bond1
      vlan-tpid: dot1ad
      svlan-range: "200-299"
      cvlan-range: "1-4000"
    isp-green:
      interface: eth3
  static-maps:
    - access-interface: eth2
      svlan: "10-99"
      handoff-group: isp-green
      transparent: true

subscriber-groups:
  groups:
    wholesale-a:
      vlans:
        - svlan: "100-199"
          cvlan: any
          parent-interface: eth1
          access-types: [l2gw]
      l2gw:
        handoff-group: isp-blue
      aaa-policy: wholesale-circuit

aaa:
  provider: radius
  policies:
    wholesale-circuit:
      type: dhcp
      format: "$circuit-id$"

plugins:
  subscriber.auth.radius:
    servers:
      - address: 192.0.2.10
        secret: wholesale-secret
```

## RADIUS integration

Access-Request, sent once per new circuit:

```
User-Name          = <circuit-id from option 82>
Calling-Station-Id = <subscriber MAC>
NAS-Port-Id        = eth1 svlan 150 cvlan 7
```

Minimal Access-Accept. osvbng allocates egress VLANs from the group ranges
and reports them in accounting:

```
l2gw.handoff-group = isp-blue
```

BSS-integrated Access-Accept, where the OSS pins the egress circuit:

```
l2gw.handoff-group = isp-blue
l2gw.svlan         = 250
l2gw.cvlan         = 1234
```

Accounting-Start/Interim/Stop then carry the resolved
`l2gw.handoff-group`, `l2gw.svlan`, and `l2gw.cvlan` plus per-circuit
upstream and downstream octet and packet counters. This is the wholesale
billing feed.

## What the subscriber experiences

Nothing osvbng-specific. Their DHCP DISCOVER reaches ISP blue's BNG (osvbng
replays the trigger frame and switches everything after it), and the lease,
ARP, IPv6 ND/DHCPv6, and all data flow between subscriber and ISP at line
rate. PPPoE subscribers on statically mapped ranges also pass transparently.
For PPPoE termination wholesale, use [L2TP LAC](l2tp-lac.md) instead.

## Verification

```
osvbngcli show l2gw circuits
sudo docker exec <container> vppctl -s /var/run/osvbng/cli.sock show osvbng l2gw circuits
```
