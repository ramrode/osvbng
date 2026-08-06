# L2GW (Layer 2 Wholesale Gateway)

The layer 2 wholesale gateway cross-connects subscriber circuits between
access networks and retail ISP handoff ports without terminating DHCP or L3.
osvbng acts as the wholesale aggregator: the first DHCPv4 DISCOVER or DHCPv6
SOLICIT on a wholesale circuit triggers AAA, the auth response selects a named
**handoff group** (plus optional explicit egress VLANs), and osvbng installs a
bidirectional L2 cross-connect with S/C-VLAN rewrite. From then on every frame
(ARP, ND, DHCP renewals, IGMP) is switched in the dataplane, and the retail
ISP's BNG terminates the subscriber. It is the IPoE analogue of LAC/LNS
wholesale.

Two operating modes share one mechanism:

- **Static maps.** An entire access S-VLAN (all C-VLANs) is cross-connected
  to an ISP by configuration. No trigger, no RADIUS, no per-subscriber state.
- **Dynamic circuits.** Per-(S,C) circuits authorized by AAA, with egress
  VLANs allocated by osvbng or returned by RADIUS.

## How the exit interface is chosen

The exit (handoff) interface is never named directly by RADIUS or by the
subscriber group. Both sides only ever reference a **handoff group label**,
and the label resolves to an interface in `l2gw.handoff-groups`. This keeps
the OSS/RADIUS integration stable when the physical wiring changes: re-point
the label at a new interface and every circuit follows.

Resolution order for a dynamic circuit:

1. `l2gw.handoff-group` attribute in the Access-Accept, if present.
2. Otherwise the subscriber group's `l2gw.handoff-group` default.
3. If neither exists, the circuit is rejected and logged.

The chosen group's `interface` field is the exit interface. The egress VLANs
on that interface come from the `l2gw.svlan` / `l2gw.cvlan` attributes when
RADIUS supplies them, and from the group's `svlan` / `svlan-range` /
`cvlan-range` allocator when it does not.

Static maps name their handoff group in configuration, so the exit interface
is fixed at config time.

## `l2gw.handoff-groups`

Bond/LACP interfaces are supported as handoff or access ports.

| Field | Type | Description | Example |
|-------|------|-------------|---------|
| `interface` | string | Exit port toward the retail ISP (physical or bond). | `bond1` |
| `vlan-tpid` | string | Outer TPID emitted toward the handoff: `dot1ad` (default) or `dot1q`. | `dot1ad` |
| `svlan` | uint16 | Pin every circuit of this group to one outer VLAN (VLAN-per-ISP model). Mutually exclusive with `svlan-range`. | `200` |
| `svlan-range` | string | Outer VLAN allocator range for dynamic circuits. | `"200-299"` |
| `cvlan-range` | string | Inner VLAN allocator range for dynamic circuits. | `"1-4000"` |

## `l2gw.static-maps`

| Field | Type | Description | Example |
|-------|------|-------------|---------|
| `access-interface` | string | Access port the S-VLANs arrive on. | `eth1` |
| `svlan` | string | Single VLAN or range. | `"10-99"` |
| `handoff-group` | string | Target handoff group. | `isp-green` |
| `handoff-svlan` | uint16 | Translate the outer tag toward the ISP. Only valid for a single-SVLAN map, because a range sharing one egress S-VLAN would collide on the handoff side. | `210` |
| `transparent` | bool | Pass all tags untouched. Implied when neither `handoff-svlan` nor the group's fixed `svlan` is set. | `true` |

C-VLANs always pass through untouched on static maps (wildcard circuits).

## Subscriber group binding (dynamic mode)

Dynamic circuits are armed per VLAN range with the `l2gw` access type, which
is mutually exclusive with all other access types on that range:

```yaml
subscriber-groups:
  groups:
    wholesale-access-a:
      vlans:
        - svlan: "100-199"
          cvlan: any
          parent-interface: eth1
          access-types: [l2gw]
      l2gw:
        handoff-group: isp-blue   # default when AAA returns no label
      aaa-policy: circuit-policy  # same policy language as IPoE (MAC, VLANs, option-82)
```

## AAA attributes

| Internal attribute | Direction | Meaning |
|---|---|---|
| `l2gw.handoff-group` | Access-Accept | Selects the handoff group by label. Falls back to the subscriber group's `l2gw.handoff-group`. |
| `l2gw.svlan` | Access-Accept | Explicit egress outer VLAN, overriding the allocator. |
| `l2gw.cvlan` | Access-Accept | Explicit egress inner VLAN, overriding the allocator. |
| `l2gw.handoff-group`, `l2gw.svlan`, `l2gw.cvlan` | Accounting | The resolved values are reported back so the OSS learns what was allocated. |

The trigger's Access-Request carries the usual circuit identity: MAC, S/C
VLANs, access interface, and DHCPv4 option-82 `circuit_id` / `remote_id`
when present. `aaa-policy` username formats work exactly as for IPoE.

## Lifecycle

- **Install.** Trigger, AAA accept, circuit programmed. The held trigger
  frame is then replayed out the handoff with the subscriber's own source
  MAC so the retail BNG learns the subscriber.
- **Accounting.** Start on install, interim on the standard cadence with
  per-circuit upstream and downstream counters from the dataplane, Stop on
  teardown. RADIUS and HTTP accounting providers work unchanged.
- **Persistence.** Dynamic circuits survive restarts. They are re-installed
  from the operational DB with no re-authentication and no duplicate
  Accounting-Start.
- **Teardown.** RADIUS Disconnect-Message or operator termination. Rejected
  circuits back off for 30 seconds before a retransmit can re-trigger.

CoA policy push is deliberately not supported. Subscriber policy belongs to
the retail ISP's BNG; osvbng is a layer 2 gateway in this role.

## Observability

- `show` path `l2gw.circuits` lists all circuits with access tuple, handoff
  resolution, static or dynamic origin, and state.
- Per-circuit packet/byte counters per direction live in the VPP stats
  segment under `/osvbng/l2gw`.
- Dataplane CLI: `show osvbng l2gw circuits` (vppctl), including counters.

## Full example

```yaml
l2gw:
  handoff-groups:
    isp-blue:
      interface: bond1
      vlan-tpid: dot1ad
      svlan-range: "200-299"
      cvlan-range: "1-4000"
    isp-green:
      interface: eth3
      svlan: 400
      cvlan-range: "1-4000"
  static-maps:
    - access-interface: eth1
      svlan: "10-99"
      handoff-group: isp-green
      transparent: true

subscriber-groups:
  groups:
    wholesale-dynamic:
      vlans:
        - svlan: "100-199"
          cvlan: any
          parent-interface: eth1
          access-types: [l2gw]
      l2gw:
        handoff-group: isp-blue
      aaa-policy: circuit-policy
```
