# Changelog

## [0.17.0](https://github.com/ramrode/osvbng/compare/v0.16.0...v0.17.0) (2026-09-05)


### Features

* **aaa:** add policy-based authentication mode ([#124](https://github.com/ramrode/osvbng/issues/124)) ([3ee887d](https://github.com/ramrode/osvbng/commit/3ee887dbcfdc1ebcf2d4fa16268356c987696455))
* **aaa:** add pool and service group attribute mappings ([#109](https://github.com/ramrode/osvbng/issues/109)) ([4640917](https://github.com/ramrode/osvbng/commit/4640917329d3428a07a1360639153c799b2badb9))
* **aaa:** add RADIUS auth provider with server failover and accounting ([#169](https://github.com/ramrode/osvbng/issues/169)) ([9217a5a](https://github.com/ramrode/osvbng/commit/9217a5a2b07b553d647cb30f4952e12699bef2d0))
* **aaa:** built-in osvbng vendor dictionary for l2gw radius attributes ([43ed1e5](https://github.com/ramrode/osvbng/commit/43ed1e55cf80d0cf42539363e0ed08c0b85a4d8d))
* **aaa:** hard-fail when policy-expanded RADIUS User-Name is empty ([#360](https://github.com/ramrode/osvbng/issues/360)) ([1db1230](https://github.com/ramrode/osvbng/commit/1db12303029c9003d9d2d721c15656224d2626fe))
* **aaa:** log returned attributes in authentication response ([#116](https://github.com/ramrode/osvbng/issues/116)) ([e83eef0](https://github.com/ramrode/osvbng/commit/e83eef08510ad2e28fc1aa60c6439e2822835e75))
* **aaa:** per-policy placeholder password for DHCP/IPoE RADIUS Access-Request ([#353](https://github.com/ramrode/osvbng/issues/353)) ([96484da](https://github.com/ramrode/osvbng/commit/96484daf7aa44653c2c1aaf38b692cf716645c89))
* **aaa:** populate RADIUS Acct octets/packets via cached interface stats ([#368](https://github.com/ramrode/osvbng/issues/368)) ([8d76b83](https://github.com/ramrode/osvbng/commit/8d76b83c33af96e385c7e5bf4be78701c99585bd))
* **aaa:** standardize radius nas-port/nas-port-id and extend accounting attributes ([#380](https://github.com/ramrode/osvbng/issues/380)) ([426970e](https://github.com/ramrode/osvbng/commit/426970e277d6d15dccf640427043a4d27c5e9003))
* **api:** complete typed metadata for all show, oper, and conf handlers ([#275](https://github.com/ramrode/osvbng/issues/275)) ([a66a7a3](https://github.com/ramrode/osvbng/commit/a66a7a377b4a15b9d7091cfc63f64bd0f129eb09))
* **api:** osvbngcli reaches the daemon via Unix domain socket ([#371](https://github.com/ramrode/osvbng/issues/371)) ([3dea4b5](https://github.com/ramrode/osvbng/commit/3dea4b5fe59857f88a6c7f1b0d13aa8c5d157ba7))
* **api:** paginate list-returning northbound show endpoints ([#292](https://github.com/ramrode/osvbng/issues/292)) ([e0e8fed](https://github.com/ramrode/osvbng/commit/e0e8fed8decd778023312ac75aad16eafd53564b))
* **api:** per-VRF multi-listener northbound API ([#372](https://github.com/ramrode/osvbng/issues/372)) ([ca37793](https://github.com/ramrode/osvbng/commit/ca377936b402ce074fa4c6b3f3883f5e363f06f2))
* **bgp:** add BGP L3VPN show command coverage ([#316](https://github.com/ramrode/osvbng/issues/316)) ([576aa26](https://github.com/ramrode/osvbng/commit/576aa26a384e081684ade3a1e836285c41c8045a))
* **bgp:** add BGP unicast show command coverage ([#315](https://github.com/ramrode/osvbng/issues/315)) ([6b7462e](https://github.com/ramrode/osvbng/commit/6b7462e9bc44d1b26a3bfa852f7714e6913107e6))
* **bgp:** add VPNv4/VPNv6 address family config model and templates ([#97](https://github.com/ramrode/osvbng/issues/97)) ([95d2c87](https://github.com/ramrode/osvbng/commit/95d2c875aebb58030c7c65cad7e16d9b73af7b60))
* **bgp:** add VPNv4/VPNv6 and L3VPN configuration and show handlers ([#98](https://github.com/ramrode/osvbng/issues/98)) ([f609909](https://github.com/ramrode/osvbng/commit/f6099094919655cb0043c2e7f166b8b48dc65d73))
* **cgnat:** add Carrier-Grade NAT with PBA mode for IPoE and PPPoE subscribers ([#183](https://github.com/ramrode/osvbng/issues/183)) ([978b4f5](https://github.com/ramrode/osvbng/commit/978b4f5fd12706b1d100a97e26de9bd277fd8b7f))
* **cgnat:** add CGNAT HA mapping sync with incremental and bulk sync ([#188](https://github.com/ramrode/osvbng/issues/188)) ([f04e43c](https://github.com/ramrode/osvbng/commit/f04e43ce5fef393af0cc16198d0f9039b69bffb6))
* **cgnat:** multi-worker datapath ([#362](https://github.com/ramrode/osvbng/issues/362)) ([05ed620](https://github.com/ramrode/osvbng/commit/05ed620799ff7f943bb301ee6fc8d83510423112))
* **cgnat:** restart-idempotent reconciler with active-mapping preflight gate ([#341](https://github.com/ramrode/osvbng/issues/341)) ([d4ff1f6](https://github.com/ramrode/osvbng/commit/d4ff1f69f3aaf056fe08a36d692168c39effd2e7))
* **cgnat:** session dump API with inside/outside/remote ip+port+proto filters ([#361](https://github.com/ramrode/osvbng/issues/361)) ([fabc43c](https://github.com/ramrode/osvbng/commit/fabc43c5418c4b03c9edb2461d3525c8f49a7956))
* **component:** add readiness signaling for async plugin startup ([#221](https://github.com/ramrode/osvbng/issues/221)) ([fe7f31d](https://github.com/ramrode/osvbng/commit/fe7f31de01ffccc9eaf64e8f662d9ace74d6bf67))
* **config:** implement auto cpu layout tiers and cp-cores intent knob ([#435](https://github.com/ramrode/osvbng/issues/435)) ([64d15d0](https://github.com/ramrode/osvbng/commit/64d15d0d9ac48e8ce6e5e9169d20546c82ab58d0))
* **config:** rename vlan-protocol to vlan-tpid with dot1ad default for QinQ ([#266](https://github.com/ramrode/osvbng/issues/266)) ([f8f1620](https://github.com/ramrode/osvbng/commit/f8f16203e8b216af6479f4742c400f19d00a02c4))
* **dataplane:** add LCP namespace support with routing protocol fixes ([#99](https://github.com/ramrode/osvbng/issues/99)) ([24351d8](https://github.com/ramrode/osvbng/commit/24351d85850cf5477c409084692b6f7f8e4b1259))
* **dataplane:** cgroup-aware CPU detection with conservative defaults ([#237](https://github.com/ramrode/osvbng/issues/237)) ([9d4cec5](https://github.com/ramrode/osvbng/commit/9d4cec5b14cd6c516937a80b8f9db57f4919066a))
* **dataplane:** support bond/LACP interfaces ([#32](https://github.com/ramrode/osvbng/issues/32)) ([#252](https://github.com/ramrode/osvbng/issues/252)) ([4b72576](https://github.com/ramrode/osvbng/commit/4b72576d515ec28c5add851e870f8e06d836cc15))
* **dataplane:** various CGNAT, QoS and VPP tweaks ([#394](https://github.com/ramrode/osvbng/issues/394)) ([07c5017](https://github.com/ramrode/osvbng/commit/07c50173494499570897d20db0b47782f86c8c13))
* **deploy:** rewrite KVM deploy script for VPP host tuning and NUMA-aware pinning ([#365](https://github.com/ramrode/osvbng/issues/365)) ([88064b7](https://github.com/ramrode/osvbng/commit/88064b73d6b335f3fc85fef0f8f38651bcef3d75))
* **dev:** one-shot QEMU/KVM development environment ([#251](https://github.com/ramrode/osvbng/issues/251)) ([0b016d4](https://github.com/ramrode/osvbng/commit/0b016d4e0a38fc100ed7a79da39d2100c99ef3e7))
* **dhcp:** add DHCP profile types and shared allocator ([#106](https://github.com/ramrode/osvbng/issues/106)) ([6191faf](https://github.com/ramrode/osvbng/commit/6191faf62ab333903a6cb7e15e9932da248c3566))
* **dhcp:** add per-VRF pool isolation to allocator registry ([#112](https://github.com/ramrode/osvbng/issues/112)) ([716f18e](https://github.com/ramrode/osvbng/commit/716f18e3edb9335ea831028d08fbdd1ffdcf7cb8))
* **dhcp:** add relay and proxy providers with Kea dev environment, smoke tests, and CI integration ([#172](https://github.com/ramrode/osvbng/issues/172)) ([1c2e5a1](https://github.com/ramrode/osvbng/commit/1c2e5a1f45ba6063b664249122fd1e2418a8a35d))
* **dhcp:** add typed AAA attributes and wire DHCPv4 provisioning context ([#107](https://github.com/ramrode/osvbng/issues/107)) ([7cbacbf](https://github.com/ramrode/osvbng/commit/7cbacbf757c461b88bf45f1a43572cb5a5c04403))
* **dhcp:** add VRF-aware pool overflow for IPv4, IANA, and PD ([#118](https://github.com/ramrode/osvbng/issues/118)) ([28c6025](https://github.com/ramrode/osvbng/commit/28c6025d8f5174b72418600d59d0be459bc35ac4))
* **dhcp:** centralize IP allocation in resolve layer ([#110](https://github.com/ramrode/osvbng/issues/110)) ([e72c44e](https://github.com/ramrode/osvbng/commit/e72c44eafd04d99efa4b5cdb8ef0e8dff9fb7def))
* **dhcp:** per-pool DHCPv4 + DHCPv6 vendor options ([#347](https://github.com/ramrode/osvbng/issues/347)) ([db1a5fb](https://github.com/ramrode/osvbng/commit/db1a5fb6d46378a8534661b2de4d0547254bc5d3))
* **dhcp:** support RFC 6221 LDRA termination in local DHCPv6 provider ([#288](https://github.com/ramrode/osvbng/issues/288)) ([bb929d8](https://github.com/ramrode/osvbng/commit/bb929d81d04a9db189dcea794109dfc480e8290d))
* **dhcpv6:** wire provisioning context through DHCPv6 provider ([#108](https://github.com/ramrode/osvbng/issues/108)) ([dff18f1](https://github.com/ramrode/osvbng/commit/dff18f16264374bc55c9922c50cde84abfe85792))
* **dhcp:** VRF support for DHCP relay/proxy ([#296](https://github.com/ramrode/osvbng/issues/296)) ([ed9e597](https://github.com/ramrode/osvbng/commit/ed9e597f8f9f44a4ee1b761def76acddf7f8332d))
* **ha:** add GARP flood on SRG promotion with batching and rate limiting ([#225](https://github.com/ramrode/osvbng/issues/225)) ([7a4f947](https://github.com/ramrode/osvbng/commit/7a4f947cf915488c77bc00e9c9cec5fcdc8f412d))
* **ha:** add HA foundation with SRG state machine, gRPC peer, and component integration ([#137](https://github.com/ramrode/osvbng/issues/137)) ([3a46bfa](https://github.com/ramrode/osvbng/commit/3a46bfa65cf84e0615ee531f502a974206268352))
* **ha:** add interface tracking, SRG counters handler, and split-brain resolution ([#141](https://github.com/ramrode/osvbng/issues/141)) ([f8b9d95](https://github.com/ramrode/osvbng/commit/f8b9d9537d49660da0e35029fcce9efffc9cc518))
* **ha:** add pool-targeted sync and full bulk sync from live sessions ([#165](https://github.com/ramrode/osvbng/issues/165)) ([9e3328b](https://github.com/ramrode/osvbng/commit/9e3328b383988015a6638a14aee91a3d1b608e96))
* **ha:** add session sync for HA standby replication ([#164](https://github.com/ramrode/osvbng/issues/164)) ([e1fafa7](https://github.com/ramrode/osvbng/commit/e1fafa785f035fcfd33896370bdca6fa0e1f8efe))
* **ha:** add SRG BGP route advertisement and withdrawal on failover ([#142](https://github.com/ramrode/osvbng/issues/142)) ([25460a4](https://github.com/ramrode/osvbng/commit/25460a44f08c994e4bb1ff0646ba4510ba53b9c3))
* **ha:** add SRG dataplane abstraction with VPP implementation and no-op fallback ([#140](https://github.com/ramrode/osvbng/issues/140)) ([dbc391c](https://github.com/ramrode/osvbng/commit/dbc391c9e657f617819f01367ae08ba59989d554))
* **ha:** add tracker-driven promotion from STANDBY_ALONE ([#196](https://github.com/ramrode/osvbng/issues/196)) ([b489188](https://github.com/ramrode/osvbng/commit/b489188249758ff6c53b769db1b86d6505805109))
* **ha:** l2gw circuit sync with standby pre-install and promotion enable ([0e92a4b](https://github.com/ramrode/osvbng/commit/0e92a4bb43c8d5106c83ae4440b4aea525736d05))
* **ha:** per-component VRF binding for HA peer-sync + gateway via pkg/netbind ([#295](https://github.com/ramrode/osvbng/issues/295)) ([abeb154](https://github.com/ramrode/osvbng/commit/abeb154c0c83486099c27467ef03bec0a5fa61c4))
* **ha:** restore synced sessions on HA promotion with failover tests ([#190](https://github.com/ramrode/osvbng/issues/190)) ([0ceeee1](https://github.com/ramrode/osvbng/commit/0ceeee15e2745117eac484c40e48aef5e3862b1d))
* **ha:** sync AAA attributes across HA failover with RADIUS validation ([#192](https://github.com/ramrode/osvbng/issues/192)) ([8a0e08b](https://github.com/ramrode/osvbng/commit/8a0e08b91be9311d6dc49bf0ac35ee3cde72f741))
* **ifmgr:** track interface IP addresses and FIB table IDs ([#93](https://github.com/ramrode/osvbng/issues/93)) ([4a0e4f7](https://github.com/ramrode/osvbng/commit/4a0e4f769bb8fdc15cf18ebd95c109036fad79e4))
* **ipoe:** C-VLAN gating when non-any is used for QinQ configurations ([#351](https://github.com/ramrode/osvbng/issues/351)) ([bd0fdfb](https://github.com/ramrode/osvbng/commit/bd0fdfb9d06b0b48529b4ef8e354a4bb299059ec))
* **ipoe:** emit RA + DHCPv6 + NA from link-local source per RFC 4861 ([#325](https://github.com/ramrode/osvbng/issues/325)) ([35424d8](https://github.com/ramrode/osvbng/commit/35424d845cfd8b20a70893603d60a89c72164396))
* **ipoe:** gate v4/v6 ingress on subscriber-group profile presence ([#355](https://github.com/ramrode/osvbng/issues/355)) ([3fb4692](https://github.com/ramrode/osvbng/commit/3fb469214c976f97a10d0a0911002bb2a1c0184b))
* **ipoe:** periodic unsolicited RAs so subscriber default routes don't expire ([#384](https://github.com/ramrode/osvbng/issues/384)) ([83b0330](https://github.com/ramrode/osvbng/commit/83b033028c5b9471d0cc565327160718ec465348))
* **ipoe:** punt IPv6 RS to control plane for per-subscriber RA handling ([#73](https://github.com/ramrode/osvbng/issues/73)) ([355eff7](https://github.com/ramrode/osvbng/commit/355eff7a5d9b7259a861be846e47a3a056cbd4df))
* **l2gw:** config schema, session model, and AAA attribute plumbing ([167f8f3](https://github.com/ramrode/osvbng/commit/167f8f37fd5b30ad10964c09a5d8fdb550a40b31))
* **l2gw:** dataplane dhcp trigger snoop arming and multi-nni access ([582051e](https://github.com/ramrode/osvbng/commit/582051e0e3ec9a4d785e55d5b3b10bd6bacd0591))
* **l2gw:** dhcpv6 relay identity options for aaa policies ([7444d92](https://github.com/ramrode/osvbng/commit/7444d9236e139eb678017c9c5afd87f95945c2c3))
* **l2gw:** export per-circuit counters via telemetry sdk ([9827251](https://github.com/ramrode/osvbng/commit/982725120b78130eb207574b4504069df0753187))
* **l2gw:** layer 2 wholesale gateway with dhcp-triggered circuits ([#402](https://github.com/ramrode/osvbng/issues/402)) ([4d772e1](https://github.com/ramrode/osvbng/commit/4d772e1ea7411d982f7931063ba39a86aad8c31f))
* **l2gw:** packet-triggered circuits with tuple identity and idle-timeout ([#403](https://github.com/ramrode/osvbng/issues/403)) ([c7f4bcb](https://github.com/ramrode/osvbng/commit/c7f4bcb0f3bf3c540373608d8bb2745921676d6e))
* **l2gw:** reconcile static maps and allocators on config commit ([5f5a710](https://github.com/ramrode/osvbng/commit/5f5a7104f5118e09d19730ba3021d5773e9a2c6d))
* **l2gw:** vxlan overlay nnis via generic tunnel rx dispatch ([#404](https://github.com/ramrode/osvbng/issues/404)) ([d316718](https://github.com/ramrode/osvbng/commit/d316718b85e97cc1dd48f121f6348de2b2e3947e))
* **l2gw:** wholesale circuit control plane ([c6b3801](https://github.com/ramrode/osvbng/commit/c6b38018c2250d7d3c12c5a7b9768001c0fb9ecb))
* **l2tp:** L2TPv2 LAC and LNS (RFC 2661) ([#305](https://github.com/ramrode/osvbng/issues/305)) ([dae651d](https://github.com/ramrode/osvbng/commit/dae651d2629f2a1386341bc96485aba07feeb14a))
* **l3vpn:** add L3VPN dev environment with loopback-based peering ([#103](https://github.com/ramrode/osvbng/issues/103)) ([7f715f5](https://github.com/ramrode/osvbng/commit/7f715f5a2ee997eab2741f1275883c63c26334ac))
* **ldp:** finalize LDP show command coverage ([#317](https://github.com/ramrode/osvbng/issues/317)) ([21db3fa](https://github.com/ramrode/osvbng/commit/21db3fa3e31792b358e78d38ebe8dcf9cf237bec))
* **logger:** async zerolog migration for non-blocking logging ([#217](https://github.com/ramrode/osvbng/issues/217)) ([c980a89](https://github.com/ramrode/osvbng/commit/c980a891c4d28ffb21dabd2833629660cb5dc25c))
* **logger:** per-key storm sampler, sampled auth-failure logging ([#436](https://github.com/ramrode/osvbng/issues/436)) ([bcae7d7](https://github.com/ramrode/osvbng/commit/bcae7d761225e188ef15d14287980966dcb5aac6))
* **models:** add username to subscriber session model ([#76](https://github.com/ramrode/osvbng/issues/76)) ([d864322](https://github.com/ramrode/osvbng/commit/d8643222817fbd064d3b127f1b8b73db5521fd11))
* **monitoring:** add subscriber session prometheus metrics and grafana dashboard ([#78](https://github.com/ramrode/osvbng/issues/78)) ([bbca6c5](https://github.com/ramrode/osvbng/commit/bbca6c5c79068bb5e98ba0cd727c70cc3f920302))
* **monitoring:** introduce typed telemetry SDK and migrated VPP exporter ([#299](https://github.com/ramrode/osvbng/issues/299)) ([40124e3](https://github.com/ramrode/osvbng/commit/40124e3a98a58d2ba426f5f2d17ab473413d5334))
* **monitoring:** migrate state.RegisterMetric callers to telemetry.RegisterMetric[T] ([d4f53cb](https://github.com/ramrode/osvbng/commit/d4f53cb95db08f560d305e8380e007e28bd19870))
* **monitoring:** retire legacy state.RegisterMetric + typed FRR surfaces ([#304](https://github.com/ramrode/osvbng/issues/304)) ([fac9a66](https://github.com/ramrode/osvbng/commit/fac9a667bc64730604a7c98db954fe2633d4664d))
* **monitoring:** show-driven metric registration ([#300](https://github.com/ramrode/osvbng/issues/300)) ([262f68f](https://github.com/ramrode/osvbng/commit/262f68ff643da2fcd9320e539f211c1929c8001e))
* **monitoring:** typed FRR JSON surfaces for BGP and LDP ([f0fdc42](https://github.com/ramrode/osvbng/commit/f0fdc422eb6c3f9b91b4891af9c02dd15dab1fd8))
* **monitoring:** unify RegisterMetric across show handler return shapes ([#302](https://github.com/ramrode/osvbng/issues/302)) ([6173696](https://github.com/ramrode/osvbng/commit/6173696587bfd7c18c6dd8322d5dbec901fc3516))
* **mpls:** add MPLS/LDP southbound API, config model, and FRR templates ([#96](https://github.com/ramrode/osvbng/issues/96)) ([795bc27](https://github.com/ramrode/osvbng/commit/795bc2767b317d6d70c4970e392e9d8890013722))
* **netbind:** plugin listener + HTTP client VRF binding with TLS ([#293](https://github.com/ramrode/osvbng/issues/293)) ([9c6f794](https://github.com/ramrode/osvbng/commit/9c6f794a413e7a13ae0df0f5530785eb2c696055))
* **ospf6:** add OSPFv3 show command coverage ([#313](https://github.com/ramrode/osvbng/issues/313)) ([f510c93](https://github.com/ramrode/osvbng/commit/f510c93b5f13671863cbee1405590d6b2925d324))
* **ospf:** add instance, interface, neighbor, lsdb, mpls-te show handlers ([#311](https://github.com/ramrode/osvbng/issues/311)) ([b9bc3e5](https://github.com/ramrode/osvbng/commit/b9bc3e5b57ebb04463a1867f550f33c2901e570a))
* **plugin:** add community cgnat-http-exporter plugin ([#284](https://github.com/ramrode/osvbng/issues/284)) ([fe2e7fb](https://github.com/ramrode/osvbng/commit/fe2e7fb82625180fa04f92690a139fa6c83e9942))
* **pppoe,ipoe:** unify session recovery via setupSession SDK ([#324](https://github.com/ramrode/osvbng/issues/324)) ([f7bad7b](https://github.com/ramrode/osvbng/commit/f7bad7bf983bb7b6b6a9de2a6fb7c4516acf5373))
* **pppoe:** add TCP MSS clamping with PPP MRU configuration (RFC 4638) ([#279](https://github.com/ramrode/osvbng/issues/279)) ([07a9f85](https://github.com/ramrode/osvbng/commit/07a9f8594e14dd1271b5f9c3c03731883a4e95c3))
* **pppoe:** ipv6 ra/nd + dhcpv6 over ppp; derive af_packet rx-queues from workers ([#385](https://github.com/ramrode/osvbng/issues/385)) ([84dc880](https://github.com/ramrode/osvbng/commit/84dc8807268712419cffe2355d34b3cb7b500f9b))
* **pppoe:** subscriber-group C-VLAN matching (parity with IPoE) ([#352](https://github.com/ramrode/osvbng/issues/352)) ([d6fe72c](https://github.com/ramrode/osvbng/commit/d6fe72cc389047de5db2d1bee265e3bcb6fc6a2b))
* **qa:** vrnetlab + Robot harness gain QEMU-mode parity for clab integration suites ([#378](https://github.com/ramrode/osvbng/issues/378)) ([8b9d54f](https://github.com/ramrode/osvbng/commit/8b9d54f1eb2d4f10767ef01ceb07640b4fe05860))
* **qos:** hierarchical qos control plane with port and s-vlan aggregates ([#422](https://github.com/ramrode/osvbng/issues/422)) ([95e35f0](https://github.com/ramrode/osvbng/commit/95e35f00117a9ce2ab93cfa18ab611a51d22b622))
* **qos:** hqos observability - session and hierarchy views, CLI renderings, prometheus fixes ([#434](https://github.com/ramrode/osvbng/issues/434)) ([b8f47a8](https://github.com/ramrode/osvbng/commit/b8f47a8eaa568e36d61558481e0a84f2beae93be))
* **qos:** implement per-subscriber policer lifecycle ([#120](https://github.com/ramrode/osvbng/issues/120)) ([a1244a9](https://github.com/ramrode/osvbng/commit/a1244a9136a068732cc5e8c5b7680c8811bf46ae))
* **qos:** integrate CAKE scheduler plugin into subscriber lifecycle ([#206](https://github.com/ramrode/osvbng/issues/206)) ([1ec9f30](https://github.com/ramrode/osvbng/commit/1ec9f30c544c54ae84b377737bb935a507a07a4a))
* **radius:** add RADIUS CoA/Disconnect-Message with subscriber runtime mutation ([#282](https://github.com/ramrode/osvbng/issues/282)) ([a8724bd](https://github.com/ramrode/osvbng/commit/a8724bda1e167510802c32246c627e20a053c278))
* **radius:** per-server VRF + source IP for auth/acct/CoA via netbind ([bcc3b22](https://github.com/ramrode/osvbng/commit/bcc3b222f5b013d7106bc523425583e175b52a3f))
* **radius:** per-server VRF + source IP for auth/acct/CoA via pkg/netbind ([#294](https://github.com/ramrode/osvbng/issues/294)) ([d2caa14](https://github.com/ramrode/osvbng/commit/d2caa14317e72fb1110a266650518ea2ee379f5a))
* **routing:** add routing policy framework with prefix-sets, community-sets, AS-path-sets, and route-policies ([#263](https://github.com/ramrode/osvbng/issues/263)) ([80f6690](https://github.com/ramrode/osvbng/commit/80f669042e9f331f19e8178c0cf79466024873e0))
* **routing:** add VRF assignment to IPoE and PPPoE subscriber sessions ([84ab258](https://github.com/ramrode/osvbng/commit/84ab258703c8ca595272826464d18eca696912e4))
* **routing:** add VRF manager with Linux VRF and VPP table lifecycle ([afc8c37](https://github.com/ramrode/osvbng/commit/afc8c37cd02e230b8797bd588e68c26bb91f82ab))
* **routing:** add VRF manager with Linux VRF and VPP table lifecycle ([#89](https://github.com/ramrode/osvbng/issues/89)) ([59f28db](https://github.com/ramrode/osvbng/commit/59f28db876a6fa1cafbba2d0e69c18891ae18674))
* **routing:** bind infrastructure interfaces to VRF during creation ([fb37a96](https://github.com/ramrode/osvbng/commit/fb37a968035cc85f63a01199ee329b596e8ca846))
* **routing:** consolidate label rename, .all telemetry, and ipv6 lab ([#318](https://github.com/ramrode/osvbng/issues/318)) ([815d4b2](https://github.com/ramrode/osvbng/commit/815d4b255d5995f5b6208ee42151af08216acd63))
* **routing:** expose authentication on bgp, ospfv2, ospfv3 ([#345](https://github.com/ramrode/osvbng/issues/345)) ([15b2641](https://github.com/ramrode/osvbng/commit/15b26411f09812a0aa3e523b8aa6dae2362d0fd0))
* **routing:** expose VPP FIB + FRR zebra RIB via protocols.fib.* and protocols.zebra.* paths ([#321](https://github.com/ramrode/osvbng/issues/321)) ([c37ecf1](https://github.com/ramrode/osvbng/commit/c37ecf12492eab111796cda54d5801d316c284a4))
* **routing:** per-VRF OSPFv2 + OSPFv3 instances ([#348](https://github.com/ramrode/osvbng/issues/348)) ([6975cf4](https://github.com/ramrode/osvbng/commit/6975cf44ca7174a2f694d833229ed4f819779e43))
* **routing:** wire VRF manager into application startup and config loading ([6e30a5b](https://github.com/ramrode/osvbng/commit/6e30a5b667d4ad620286f76de7e54f2942697c65))
* **show:** add show interfaces framework ([#254](https://github.com/ramrode/osvbng/issues/254)) ([601708a](https://github.com/ramrode/osvbng/commit/601708aa98e8000bfd2f44312acfbd0e3433f851))
* **southbound:** add explicit sub-interface support ([#259](https://github.com/ramrode/osvbng/issues/259)) ([45003b5](https://github.com/ramrode/osvbng/commit/45003b5f1834eda57d5a97225b59fc81c2ee32d1))
* **southbound:** evpn-signaled vxlan tunnels with dynamic vtep discovery ([#409](https://github.com/ramrode/osvbng/issues/409)) ([6e21c0f](https://github.com/ramrode/osvbng/commit/6e21c0fae9c92e1171549ef3bf13e61e191a11b6))
* **southbound:** generate osvbng_l2gw binapi bindings ([31b7b36](https://github.com/ramrode/osvbng/commit/31b7b3617d0cc64d173db164ea94f81cbb60c5b0))
* **southbound:** l2gw circuit stats and paired counter indices ([dba1065](https://github.com/ramrode/osvbng/commit/dba1065a0fc6ec35a64a2591a2d61147043ff3df))
* **southbound:** l2gw circuit wrappers and interface ([6ffcb53](https://github.com/ramrode/osvbng/commit/6ffcb53976d806daeda114aa3d2606634be6f8f2))
* **southbound:** pseudowire headend termination prep and vxlan transport termination ([#405](https://github.com/ramrode/osvbng/issues/405)) ([c668a58](https://github.com/ramrode/osvbng/commit/c668a5870397b4238cda977a3a29b2e946a277c0))
* **southbound:** regenerate osvbng_l2gw binapi for trigger svlan range ([7910fa6](https://github.com/ramrode/osvbng/commit/7910fa6240f1f647a214d17f2d72ca9b18cd238c))
* **subscriber:** add plugin-agnostic subscriber runtime mutation API ([#281](https://github.com/ramrode/osvbng/issues/281)) ([ca249ca](https://github.com/ramrode/osvbng/commit/ca249ca8f8bf169c4060bc3f490e31667dd125d4))
* **subscriber:** mixed IPoE+PPPoE on shared S-VLAN range ([#306](https://github.com/ramrode/osvbng/issues/306)) ([62beca1](https://github.com/ramrode/osvbng/commit/62beca193bebda0c06178bc1e3e6b21d450a4c6a))
* **subscriber:** subscriber clear session functionality ([#77](https://github.com/ramrode/osvbng/issues/77)) ([aa89965](https://github.com/ramrode/osvbng/commit/aa899652a212caad02f9cf15c3c6aae88e9ffe53))
* **subscriber:** unset unnumbered on session release ([#307](https://github.com/ramrode/osvbng/issues/307)) ([e86bbd9](https://github.com/ramrode/osvbng/commit/e86bbd9e257616e9cb950506652632b4fa505d91))
* **svcgroup:** add service group resolver with three-layer merge resolution ([199b979](https://github.com/ramrode/osvbng/commit/199b979c184b51395f49c10954a2cf46dcf248cb))
* **svcgroup:** added support for service groups ([abf753d](https://github.com/ramrode/osvbng/commit/abf753dd6746d09ac7558fb79be6fa28e31040d4))
* **svcgroup:** added support for service groups ([#91](https://github.com/ramrode/osvbng/issues/91)) ([abf753d](https://github.com/ramrode/osvbng/commit/abf753dd6746d09ac7558fb79be6fa28e31040d4))
* **upgrade:** add osvbngcli upgrade builtin (Tier A file-swap) ([#332](https://github.com/ramrode/osvbng/issues/332)) ([321e942](https://github.com/ramrode/osvbng/commit/321e9422b2af777de1b7a5df515dfcb9c29beb24))
* **upgrade:** tier-a v2 upgrade pipeline + QEMU test infrastructure ([#375](https://github.com/ramrode/osvbng/issues/375)) ([539d09b](https://github.com/ramrode/osvbng/commit/539d09b20259d344e739561b2c9935766f0e8fce))
* **vrf:** subscriber VRF cascade + VRF-lite / L3VPN integration suites ([#289](https://github.com/ramrode/osvbng/issues/289)) ([9fda8b9](https://github.com/ramrode/osvbng/commit/9fda8b918fffe1ecff78091b849735c2acb85221))
* **watchdog:** add VPP health monitoring and dataplane recovery ([#128](https://github.com/ramrode/osvbng/issues/128)) ([656d2ba](https://github.com/ramrode/osvbng/commit/656d2ba15c229204a96186a47cc35f655ded41c1))


### Bug Fixes

* **aaa:** add Message-Authenticator (attr 80) to Access-Request packets ([#181](https://github.com/ramrode/osvbng/issues/181)) ([068f967](https://github.com/ramrode/osvbng/commit/068f967483e1ce6b8a80ea086276218f911602ee))
* **aaa:** address RADIUS auth/accounting issues from code review ([#174](https://github.com/ramrode/osvbng/issues/174)) ([a829c53](https://github.com/ramrode/osvbng/commit/a829c53c54bb3b16738e746f6a3547d8a968292a))
* **aaa:** hash session-id to deterministic accounting bucket ([#358](https://github.com/ramrode/osvbng/issues/358)) ([34da0d4](https://github.com/ramrode/osvbng/commit/34da0d478c6ff40fd0323ec72e390d774b307a30))
* **aaa:** make the accounting bucket sleep interruptible ([#458](https://github.com/ramrode/osvbng/issues/458)) ([98cdd83](https://github.com/ramrode/osvbng/commit/98cdd83f7f355a0a56e8c5b8d80154871b01af67))
* **aaa:** stop policy username expansion from gating local auth ([aec07e3](https://github.com/ramrode/osvbng/commit/aec07e3bd94683cc72079035ccc61e10dd61eee5))
* **aaa:** stop policy username expansion from gating local auth ([#388](https://github.com/ramrode/osvbng/issues/388)) ([40eda4e](https://github.com/ramrode/osvbng/commit/40eda4e765206c4077a24ef69f70586e1a512ae7))
* **aaa:** use atomic pointer for global RADIUS provider reference ([#180](https://github.com/ramrode/osvbng/issues/180)) ([c30b869](https://github.com/ramrode/osvbng/commit/c30b869f2b149c662892f932deaf464eb3b7446a))
* **aaa:** wire up RadiusAttr name resolution for response mappings ([#179](https://github.com/ramrode/osvbng/issues/179)) ([e43d6e4](https://github.com/ramrode/osvbng/commit/e43d6e4b19952c5d079745f5a39c25f74f4fa850))
* **arp:** enforce VRF-aware ARP response filtering ([#94](https://github.com/ramrode/osvbng/issues/94)) ([1e7f7b4](https://github.com/ramrode/osvbng/commit/1e7f7b43943a4cd42adba615c79efeaf568fbc9c))
* **arp:** ignore DAD probe for client's own assigned IP ([#205](https://github.com/ramrode/osvbng/issues/205)) ([c623065](https://github.com/ramrode/osvbng/commit/c623065b3a55d974c0bb3d37cb8abd8b5561aef9))
* **arp:** use per-interface IP dump and ifmgr cache ([#95](https://github.com/ramrode/osvbng/issues/95)) ([77a88c4](https://github.com/ramrode/osvbng/commit/77a88c4b842c3cee6e9a17068f1eea2225e8e4b5))
* **autoconfig:** gate subscriber IPv6 derivation on VRF IPv6 unicast ([#314](https://github.com/ramrode/osvbng/issues/314)) ([96f011a](https://github.com/ramrode/osvbng/commit/96f011a8fccc80ea47b0e6eb02855e36f5903083))
* **bgp:** activate neighbors in unicast address families ([#121](https://github.com/ramrode/osvbng/issues/121)) ([9b9f0ad](https://github.com/ramrode/osvbng/commit/9b9f0ad87c7a4609233a66aabc93cbbf918ebce6))
* **bgp:** add blackhole routes for advertised pool networks ([#122](https://github.com/ramrode/osvbng/issues/122)) ([799e94a](https://github.com/ramrode/osvbng/commit/799e94afcb1f830c4ddb4099e323aac3affbdeac))
* **bgp:** add no bgp default ipv4-unicast to template ([#101](https://github.com/ramrode/osvbng/issues/101)) ([a3d9801](https://github.com/ramrode/osvbng/commit/a3d9801c7a26d36db152a1ec724d5f7c6f5e8d5a))
* **bgp:** drop direct vtysh from instance Apply so ASN changes survive osvbngd restart ([#377](https://github.com/ramrode/osvbng/issues/377)) ([580d240](https://github.com/ramrode/osvbng/commit/580d240253ff30cdf8d2bc0c36b9dd6487674f2a))
* **bgp:** guard nil values in instance Apply/Rollback to prevent commit-rollback panic ([215ae5a](https://github.com/ramrode/osvbng/commit/215ae5a0f1a9c6da9656edea92e17596e34b7235))
* **build:** copy template subdirectories in qemu image build ([44dd5e9](https://github.com/ramrode/osvbng/commit/44dd5e96f2e946497918410950e3a64c5bb7a6cd))
* **build:** copy template subdirectories in qemu image build ([#81](https://github.com/ramrode/osvbng/issues/81)) ([d598714](https://github.com/ramrode/osvbng/commit/d598714e5f41546802bab4d76e62e6eaf963f595))
* **cgnat:** key mappings and bypass by the session vrf on every programming path ([#505](https://github.com/ramrode/osvbng/issues/505)) ([86ef344](https://github.com/ramrode/osvbng/commit/86ef34433d6635bb7420c27a5ecda2633c66436a))
* **cgnat:** prevent gateway destination mapping via in2out vpp node ([#359](https://github.com/ramrode/osvbng/issues/359)) ([a0b8a5d](https://github.com/ramrode/osvbng/commit/a0b8a5d22cec131c204a851944a1011bc71fdb09))
* **cgnat:** rebuild plugin so with thread-safe shared session pool ([c7e96d7](https://github.com/ramrode/osvbng/commit/c7e96d7fbf2a6d6c01f5e230fc66b0e857f19b73))
* **cgnat:** rebuild plugin so with thread-safe shared session pool ([#386](https://github.com/ramrode/osvbng/issues/386)) ([2864846](https://github.com/ramrode/osvbng/commit/2864846de4e2452263d5f19feaa38637c9a8310a))
* **cgnat:** refuse pool config that crash-loops or silently misconfigures ([#509](https://github.com/ramrode/osvbng/issues/509)) ([9bd7938](https://github.com/ramrode/osvbng/commit/9bd7938b0a11be0a2d5bd10e957e0ee9a188ea7d))
* **cgnat:** skip orphan child cleanup for pools dropped via cascade-delete ([#343](https://github.com/ramrode/osvbng/issues/343)) ([99ae909](https://github.com/ramrode/osvbng/commit/99ae9091873cc892b6309fd0e6b47f9b269fe607))
* **ci:** add checkout step for Discord changelog notification ([#193](https://github.com/ramrode/osvbng/issues/193)) ([df8c331](https://github.com/ramrode/osvbng/commit/df8c331b4a71c747ebd8a8d161875bbd54640893))
* **ci:** add topology cleanup and diagnostics to test workflows ([#234](https://github.com/ramrode/osvbng/issues/234)) ([458bcc4](https://github.com/ramrode/osvbng/commit/458bcc445bd600b642ecbb6062cdbd367826ddd3))
* **ci:** extract PR number from release-please JSON output ([#199](https://github.com/ramrode/osvbng/issues/199)) ([018d9df](https://github.com/ramrode/osvbng/commit/018d9df74f827b510fc9593779902f2ee37fdad9))
* **ci:** handle non-zero exit codes in test setup version checks ([#235](https://github.com/ramrode/osvbng/issues/235)) ([1915d19](https://github.com/ramrode/osvbng/commit/1915d19b9acf5b124b7a023d5feaf3a792eb682f))
* **ci:** pass all github expressions as env vars to avoid shell parsing errors ([#195](https://github.com/ramrode/osvbng/issues/195)) ([1e98746](https://github.com/ramrode/osvbng/commit/1e987468978cc913c06268910d2de15676f945ea))
* **ci:** pre-create containerlab network to prevent parallel deploy race ([#236](https://github.com/ramrode/osvbng/issues/236)) ([b517363](https://github.com/ramrode/osvbng/commit/b5173631760a28601319c32e60cb02a1c4740a6a))
* **ci:** prevent shell parsing failures in Discord webhook notifications ([#197](https://github.com/ramrode/osvbng/issues/197)) ([eaa8be2](https://github.com/ramrode/osvbng/commit/eaa8be253e93ac8d5e9f791b7113e07869ca2240))
* **ci:** record the swept state on the rig, not in an actions variable ([#470](https://github.com/ramrode/osvbng/issues/470)) ([a8a4053](https://github.com/ramrode/osvbng/commit/a8a4053f50ff56a1cdc5b19339f57eb27f0f6219))
* **ci:** trigger discord webhook on release-please PR creation ([#212](https://github.com/ramrode/osvbng/issues/212)) ([a575bc7](https://github.com/ramrode/osvbng/commit/a575bc751a0edb67c9478348554d535528d5ada3))
* **ci:** use github context instead of git log for Discord notifications ([#194](https://github.com/ramrode/osvbng/issues/194)) ([91e1267](https://github.com/ramrode/osvbng/commit/91e1267af3508f50e20dd59c666667c1cbef5364))
* **cli:** stabilise show output column order and drop dead gRPC command framework ([#389](https://github.com/ramrode/osvbng/issues/389)) ([d694f98](https://github.com/ramrode/osvbng/commit/d694f98fba80e30946c3395aa7ba3126113e0edf))
* **component:** AllReady skips plugin components, matching WaitReady semantics ([#335](https://github.com/ramrode/osvbng/issues/335)) ([74cc66c](https://github.com/ramrode/osvbng/commit/74cc66ca61a8350efc3c80a3271cb0d9df1f576b))
* **configmgr:** skip no-op changes, suppress empty versions on restart, render values as JSON ([#320](https://github.com/ramrode/osvbng/issues/320)) ([616cffd](https://github.com/ramrode/osvbng/commit/616cffd084e7ce028ad1fb379f5ddaebc7b016f6))
* **config:** stabilize topological sort for deterministic change ordering ([e79affc](https://github.com/ramrode/osvbng/commit/e79affc2178e8be8fb638f10e4c970dcecbb2739))
* **dataplane:** bring up loopback in LCP namespace ([#102](https://github.com/ramrode/osvbng/issues/102)) ([a344d85](https://github.com/ramrode/osvbng/commit/a344d853ae32c1db50549fc1a1ce20d18b8f7774))
* **dataplane:** bring up loopback in LCP namespace and register in ifmgr ([#104](https://github.com/ramrode/osvbng/issues/104)) ([554e0ba](https://github.com/ramrode/osvbng/commit/554e0ba2e7f0e49acb60eb3618c07e67ad478189))
* **dataplane:** default AF_PACKET interfaces to interrupt rx-mode ([#209](https://github.com/ramrode/osvbng/issues/209)) ([76c83fa](https://github.com/ramrode/osvbng/commit/76c83faff678238a2cb015224a1d3e17bb74bef4))
* **dataplane:** default QinQ outer TPID to 802.1ad ([183d408](https://github.com/ramrode/osvbng/commit/183d408aa7c42dde7fa95601aeaee75ae403fc33))
* **dataplane:** default QinQ outer TPID to 802.1ad ([#86](https://github.com/ramrode/osvbng/issues/86)) ([183d408](https://github.com/ramrode/osvbng/commit/183d408aa7c42dde7fa95601aeaee75ae403fc33))
* **dataplane:** default QinQ outer TPID to 802.1ad with per-group override ([97eab03](https://github.com/ramrode/osvbng/commit/97eab034faf6a03a87507527a8d07e1289bcf924))
* **dataplane:** default to exactly 1 thread worker regardless of available cores ([#238](https://github.com/ramrode/osvbng/issues/238)) ([5883f15](https://github.com/ramrode/osvbng/commit/5883f15283318ae833240ef8d2c48ae72333e142))
* **dataplane:** default VPP heap to 1G and increase egress channel capacity for DPDK scale ([#250](https://github.com/ramrode/osvbng/issues/250)) ([f0bd38a](https://github.com/ramrode/osvbng/commit/f0bd38ac9cbc7a07d0277f3dad530a7605d0593e))
* **dataplane:** skip sw_interface_set_table when interface already in target VRF table ([b2f6f0f](https://github.com/ramrode/osvbng/commit/b2f6f0fd51c9f9419614a0c1773e812b0d9f61ad))
* **dataplane:** use poll instead of busy-read on punt eventfd ([#211](https://github.com/ramrode/osvbng/issues/211)) ([afaf0eb](https://github.com/ramrode/osvbng/commit/afaf0ebe49f08a673d48eab280d655e59932281d))
* **dhcp:** address relay and proxy issues ([#175](https://github.com/ramrode/osvbng/issues/175)) ([0daeaf9](https://github.com/ramrode/osvbng/commit/0daeaf9925386781dda98f9b78b68f5c6c33e0a8))
* **dhcp:** detect and reject static/dynamic IP reservation collisions ([#119](https://github.com/ramrode/osvbng/issues/119)) ([2948c35](https://github.com/ramrode/osvbng/commit/2948c356cf5303067a9af104d44d20d96fd5fe98))
* **dhcp:** resolve per-pool gateway and add service group pool selection ([#117](https://github.com/ramrode/osvbng/issues/117)) ([dc10ba9](https://github.com/ramrode/osvbng/commit/dc10ba986515b96daaa9862fdf8e20dc4b05600b))
* **dhcp:** use compound keys for pending map to prevent XID collisions ([#178](https://github.com/ramrode/osvbng/issues/178)) ([998eb9d](https://github.com/ramrode/osvbng/commit/998eb9d6c1bce354b958c64dd69ea9bb4b286b5a))
* **dhcpv6:** add enterprise-number prefix to Remote-ID option per RFC 4649 ([#177](https://github.com/ramrode/osvbng/issues/177)) ([6dd6a1d](https://github.com/ramrode/osvbng/commit/6dd6a1d6286c5ca0b6a2c11aa21e824bcce44f3c))
* **dhcpv6:** reject overflowing option lengths instead of panicking ([#431](https://github.com/ramrode/osvbng/issues/431)) ([6545ccd](https://github.com/ramrode/osvbng/commit/6545ccdf86ad76657ee2572326bf2e4a521163b0))
* **dhcpv6:** resolve session cache race between concurrent IPv4 and IPv6 lifecycle events ([#222](https://github.com/ramrode/osvbng/issues/222)) ([1144dcc](https://github.com/ramrode/osvbng/commit/1144dcce34091440ce10cc32198419637e7ac8de))
* **dhcpv6:** send LDRA Relay-Reply to port 547, not 546 ([#424](https://github.com/ramrode/osvbng/issues/424)) ([46b7101](https://github.com/ramrode/osvbng/commit/46b7101705f48183827d2797a9ace69c5db30c33))
* **dhcpv6:** use SRG virtual MAC or access interface MAC for proxy DUID ([#176](https://github.com/ramrode/osvbng/issues/176)) ([1cc6715](https://github.com/ramrode/osvbng/commit/1cc6715a0c28c3b27bb1e2406760f73c8e92887b))
* **docker:** create dataplane netns in published image entrypoint ([#135](https://github.com/ramrode/osvbng/issues/135)) ([a82ea71](https://github.com/ramrode/osvbng/commit/a82ea7144c76d63bcb8f80a6e94e34a6654c7416))
* **docker:** decide the plugin drop from a marker the tmp purge cannot delete ([#465](https://github.com/ramrode/osvbng/issues/465)) ([8475153](https://github.com/ramrode/osvbng/commit/8475153d62ed004d84d5889efec2eb657a0e7250))
* **docker:** grow the hugepage pool instead of resetting it ([#433](https://github.com/ramrode/osvbng/issues/433)) ([0925816](https://github.com/ramrode/osvbng/commit/0925816fb9a8097a46dc71f1e075523882e3fc41))
* **docker:** treat the auto core sentinel as unset before taskset ([#471](https://github.com/ramrode/osvbng/issues/471)) ([77f1a58](https://github.com/ramrode/osvbng/commit/77f1a5808ad3edab64f9daecd548667ceedebc66))
* **docker:** wait for the dataplane API socket before starting osvbngd ([#477](https://github.com/ramrode/osvbng/issues/477)) ([3e2d358](https://github.com/ramrode/osvbng/commit/3e2d3587f8acf7486333b903fff3a24ff6dab96b))
* **ha:** handle reverse event ordering for tracker promotion ([#198](https://github.com/ramrode/osvbng/issues/198)) ([4c8983f](https://github.com/ramrode/osvbng/commit/4c8983f8c00cba18be78f1800878eac49efa5540))
* **ha:** only restore synced sessions when promoted from STANDBY_ALONE ([#201](https://github.com/ramrode/osvbng/issues/201)) ([70a4999](https://github.com/ramrode/osvbng/commit/70a4999120d91013a0fd682095f7ec84b956499d))
* **ha:** prevent standby from responding to PPPoE discovery ([#227](https://github.com/ramrode/osvbng/issues/227)) ([33a630f](https://github.com/ramrode/osvbng/commit/33a630f5da68faf2438dc631147fa966e362e8da))
* **ha:** prevent standby from responding to subscriber ARP and IPv6 ND ([#226](https://github.com/ramrode/osvbng/issues/226)) ([0f0da43](https://github.com/ramrode/osvbng/commit/0f0da439a18f02d9ddd0157b7dc28ddaa42765cf))
* **ha:** promote WAITING SRGs to ACTIVE_SOLO when peer is unreachable ([#171](https://github.com/ramrode/osvbng/issues/171)) ([749f244](https://github.com/ramrode/osvbng/commit/749f2447536b5846a700a619671e1149c284f35d))
* **ha:** restore synced sessions on graceful switchover and resolve encap sub-interface ([ed9e3da](https://github.com/ramrode/osvbng/commit/ed9e3da6a584f2ff722117f8b0ae5b172b5d7545))
* **ha:** restore synced sessions on graceful switchover and resolve encap sub-interface ([#400](https://github.com/ramrode/osvbng/issues/400)) ([b76ee3d](https://github.com/ramrode/osvbng/commit/b76ee3dadaa7de2367bb8b8522049729e6bf3ba4))
* **ha:** send the srg garp flood on the access interface with per-session vlan encap ([#418](https://github.com/ramrode/osvbng/issues/418)) ([1150e78](https://github.com/ramrode/osvbng/commit/1150e78fa0473ea2c57962e082804eef1ad04aa9))
* **ha:** standby does not auto-promote on peer loss ([#163](https://github.com/ramrode/osvbng/issues/163)) ([0f51126](https://github.com/ramrode/osvbng/commit/0f51126bf770772a2ae0375ee3571a2c80df3d36))
* **ha:** update GetVirtualMAC test to reflect active-state guard ([#229](https://github.com/ramrode/osvbng/issues/229)) ([43c311e](https://github.com/ramrode/osvbng/commit/43c311e82c1be4470afc1e6b52cdc64bb15950e0))
* **install:** bind frr.service to the dataplane netns on the qemu image ([#369](https://github.com/ramrode/osvbng/issues/369)) ([0415aa5](https://github.com/ramrode/osvbng/commit/0415aa5d5f146f0b75dc27280162747fb053236a))
* **interfaces:** re-enabling a subinterface no longer leaves it down ([#376](https://github.com/ramrode/osvbng/issues/376)) ([321fb0c](https://github.com/ramrode/osvbng/commit/321fb0caebdea45970186c4e81ce8fb33ca1b1d2))
* **ipoe,pppoe:** apply service-group qos/acl/urpf bindings on fresh session bring-up ([e6bfb0b](https://github.com/ramrode/osvbng/commit/e6bfb0b0d0c30d1fe64759f154584f9600ee7191))
* **ipoe:** advertise subscriber RA prefix off-link by default ([#382](https://github.com/ramrode/osvbng/issues/382)) ([dc2e34d](https://github.com/ramrode/osvbng/commit/dc2e34dc8d81f01758645df66b24803ea6ecafbf))
* **ipoe:** always call ip_table_bind for new IPoE session interfaces ([#134](https://github.com/ramrode/osvbng/issues/134)) ([f12ddc5](https://github.com/ramrode/osvbng/commit/f12ddc5cdb40e6b19f175c63af8cc973a96aced2))
* **ipoe:** eliminate session index race causing 1-in-48k stuck session ([#248](https://github.com/ramrode/osvbng/issues/248)) ([3fbc24b](https://github.com/ramrode/osvbng/commit/3fbc24b4fc1e17d13c03c7c3bcbf7b7d315aef14))
* **ipoe:** fix VPP crash on session re-creation and release resource leaks ([#133](https://github.com/ramrode/osvbng/issues/133)) ([80c24ef](https://github.com/ramrode/osvbng/commit/80c24ef4a03be7d1dc144600480b1a1c84cd9ced))
* **ipoe:** ignore foreign DHCPv4 server responses in server mode ([#370](https://github.com/ramrode/osvbng/issues/370)) ([8fd58d4](https://github.com/ramrode/osvbng/commit/8fd58d4b0cf0bd1038ffb73500ba771d562081b5))
* **ipoe:** log error when address profile not found ([#125](https://github.com/ramrode/osvbng/issues/125)) ([9963d1a](https://github.com/ramrode/osvbng/commit/9963d1a0a34c75d7e1cf5325db3b42efde7db336))
* **ipoe:** make runtime session expiry lease-aware ([#354](https://github.com/ramrode/osvbng/issues/354)) ([7b2fea2](https://github.com/ramrode/osvbng/commit/7b2fea28c00b672744eea2ab01ef6872298fa61f))
* **ipoe:** preserve subscriber ActivatedAt across DHCP renewals ([#390](https://github.com/ramrode/osvbng/issues/390)) ([c4ba0bf](https://github.com/ramrode/osvbng/commit/c4ba0bf724d001fbaf7956913974f379ea56ddbb))
* **ipoe:** reserve restored session addresses in allocator on opdb restore ([#363](https://github.com/ramrode/osvbng/issues/363)) ([16cb21e](https://github.com/ramrode/osvbng/commit/16cb21ed31b54d3f6912271f69ded03ea2244dad))
* **ipoe:** reset stale AAA-approved sessions ([#82](https://github.com/ramrode/osvbng/issues/82)) ([268b7aa](https://github.com/ramrode/osvbng/commit/268b7aafc3e63edd93164e4e6c0429845c0a6979))
* **ipoe:** send unsolicited na for restored sessions ([#451](https://github.com/ramrode/osvbng/issues/451)) ([d66ff01](https://github.com/ramrode/osvbng/commit/d66ff01569174e3fc8d1fa84721e2d45377d0453))
* **ipoe:** subscriber-group matching ignores inner C-VLAN; gate via precomputed S/C-VLAN index ([#350](https://github.com/ramrode/osvbng/issues/350)) ([0e85a04](https://github.com/ramrode/osvbng/commit/0e85a0414a074a9927078aef408ddcfca993209d))
* **ipv6:** use separate IPv6 profile name in allocator and add PPPoE IANA allocation ([#159](https://github.com/ramrode/osvbng/issues/159)) ([7ae29f4](https://github.com/ramrode/osvbng/commit/7ae29f42426104a2a2406157cbb0062166684f2a))
* **monitoring:** drop hostname label from BGP neighbor/VPN peer ([3f5304a](https://github.com/ramrode/osvbng/commit/3f5304a25799dfa326fa862bc9c64aeedae6724f))
* **netbind:** open VRF-bound sockets in the LCP netns ([#346](https://github.com/ramrode/osvbng/issues/346)) ([ba6b1a3](https://github.com/ramrode/osvbng/commit/ba6b1a38949b0e59bf6e1ba4aa6314ec9e5268f2))
* **netbind:** pick LocalAddr family by dial network for UDP source pinning ([1bfbf0e](https://github.com/ramrode/osvbng/commit/1bfbf0e6c1ce5ae839ef8e75397a803448dd7bd0))
* **opdb:** apply SQLite pragmas via DSN to all pool connections + retry on busy ([#330](https://github.com/ramrode/osvbng/issues/330)) ([0e00fd3](https://github.com/ramrode/osvbng/commit/0e00fd30535193aa934c8315c0bc29eb026e7381))
* **ospf:** use accept-all-interfaces mfib flag ([#100](https://github.com/ramrode/osvbng/issues/100)) ([88a6fc5](https://github.com/ramrode/osvbng/commit/88a6fc5859b894a2408ebd18a1094f4ba8c98e1d))
* **osvbng:** startup and config re-apply resilience (radius dial, bgp rollback, vrf table) ([#357](https://github.com/ramrode/osvbng/issues/357)) ([424e2d0](https://github.com/ramrode/osvbng/commit/424e2d09ad80243a09917763b3b403fadb1c2670))
* poll for bngblaster exit instead of fixed sleep ([#157](https://github.com/ramrode/osvbng/issues/157)) ([5e33939](https://github.com/ramrode/osvbng/commit/5e3393931ac8c2b1a6d74acfe1b0cde04e7ebfae))
* poll for IPv6 readiness in session tests and enable IP6CP for PPPoE ([#158](https://github.com/ramrode/osvbng/issues/158)) ([6428dc4](https://github.com/ramrode/osvbng/commit/6428dc4f37f00f32b7808e0e382345cb2f8dbceb))
* **pppoe:** bind the session interface unnumbered on ha restore ([#510](https://github.com/ramrode/osvbng/issues/510)) ([1e7b0f7](https://github.com/ramrode/osvbng/commit/1e7b0f794b70fd5c28a0438836d3c7e963bcc314))
* **pppoe:** make dual-stack dhcpv6 reach the dataplane ([#446](https://github.com/ramrode/osvbng/issues/446)) ([f13730d](https://github.com/ramrode/osvbng/commit/f13730d3ef120ede8b041f2f82f15dc8e3ef2a13))
* **pppoe:** resolve PPPoE session egress and unicast packet handling ([#75](https://github.com/ramrode/osvbng/issues/75)) ([560079f](https://github.com/ramrode/osvbng/commit/560079f64335002b81875fe936e948826e1b154b))
* **pppoe:** size session interface mtu as mru plus pppoe overhead ([#406](https://github.com/ramrode/osvbng/issues/406)) ([2066cf5](https://github.com/ramrode/osvbng/commit/2066cf55fb2278712e7c489f709def43f721b0f8))
* **qos:** convert AAA download-rate from bps to kbps without truncation ([#426](https://github.com/ramrode/osvbng/issues/426)) ([8c6d452](https://github.com/ramrode/osvbng/commit/8c6d452511dd4e8135d92b7d5187bb97eff6642f))
* **radius:** defer server dial so an unreachable RADIUS cannot block startup ([2259ce6](https://github.com/ramrode/osvbng/commit/2259ce64eb9a56ed5ed39c49601c53c877f4f820))
* **routing:** use loaded config for FRR config generation ([7cdb6f2](https://github.com/ramrode/osvbng/commit/7cdb6f2d7471da28b8565996d7beac9c21d773fc))
* **routing:** use loaded config for FRR config generation ([e529fa2](https://github.com/ramrode/osvbng/commit/e529fa2bba34c9a906f3a329064a533d986fed9b))
* **routing:** use loaded config for FRR config generation ([#88](https://github.com/ramrode/osvbng/issues/88)) ([7cdb6f2](https://github.com/ramrode/osvbng/commit/7cdb6f2d7471da28b8565996d7beac9c21d773fc))
* **scripts:** avoid set -e death in deploy-vm.sh on hosts without SMT ([#413](https://github.com/ramrode/osvbng/issues/413)) ([a8835da](https://github.com/ramrode/osvbng/commit/a8835da96dd7d332f0eeb96cdf2a791ccaa84f0e))
* **shm:** drain one punt ring per VPP thread for the v2 punt protocol ([#415](https://github.com/ramrode/osvbng/issues/415)) ([52e23a5](https://github.com/ramrode/osvbng/commit/52e23a50a566c2cc55b8de6b73b64958f0f273ea))
* **show, osvbngcli:** typed BoolOption, fast exit on all paths, hide .all paths, scope flag help, render wildcard descriptions ([#319](https://github.com/ramrode/osvbng/issues/319)) ([5e29dfa](https://github.com/ramrode/osvbng/commit/5e29dfa71942f8c1f701b60ea65522eae660f664))
* **southbound:** auto-create LCP pairs, remove lcp config option ([#271](https://github.com/ramrode/osvbng/issues/271)) ([4a680b2](https://github.com/ramrode/osvbng/commit/4a680b29052d43565cf79625fc88c46c4daf96ae))
* **southbound:** BNG parent promisc when Linux link missing ([#301](https://github.com/ramrode/osvbng/issues/301)) ([ef550c4](https://github.com/ramrode/osvbng/commit/ef550c4ab4c15940ffc70568ce4d3fb7e399a85c))
* **southbound:** call VPP SwInterfaceSetMtu API in SetInterfaceMTU ([#258](https://github.com/ramrode/osvbng/issues/258)) ([9e55741](https://github.com/ramrode/osvbng/commit/9e55741d8e637a25a84e1c1dc5a95dd63ea5de29))
* **southbound:** disable kernel ipv6 on interfaces claimed for the dataplane ([#461](https://github.com/ramrode/osvbng/issues/461)) ([975d806](https://github.com/ramrode/osvbng/commit/975d806cc304b149560db312c4060d76faabaefa))
* **southbound:** hardcode LCP dataplane namespace and handle existing pairs on restart ([#127](https://github.com/ramrode/osvbng/issues/127)) ([5265f6c](https://github.com/ramrode/osvbng/commit/5265f6c3aa6ede4a3bba7bdb0cc7e9f548c49a62))
* **southbound:** make bond create idempotent on osvbngd restart ([#333](https://github.com/ramrode/osvbng/issues/333)) ([4b7301e](https://github.com/ramrode/osvbng/commit/4b7301e811bb6d723a8f47efd831d37968662796))
* **southbound:** populate ifmgr.Interface.OuterTPID on CreateSubinterface ([#328](https://github.com/ramrode/osvbng/issues/328)) ([bd7bbd5](https://github.com/ramrode/osvbng/commit/bd7bbd5d60b1ebdb500f31c6598859201c61ecb3))
* **southbound:** re-attach ipoe sessions on already-exists instead of failing restore ([#462](https://github.com/ramrode/osvbng/issues/462)) ([20619c4](https://github.com/ramrode/osvbng/commit/20619c4ac0c5a28b955d95c80d6916d375f0a77a))
* **southbound:** replace FIFO async worker with stream pool and parameterize VPP memory config ([#246](https://github.com/ramrode/osvbng/issues/246)) ([46e2bab](https://github.com/ramrode/osvbng/commit/46e2bab36d3f46e7de82dd3df4f5e663d91fc269))
* **southbound:** set af-packet MAC at creation instead of post-hoc sync ([#123](https://github.com/ramrode/osvbng/issues/123)) ([e5fde1c](https://github.com/ramrode/osvbng/commit/e5fde1c783f0e05ae93eb51a3975ea1111eb4f14))
* **southbound:** set HW MTU, host interface MTU, and QinQ sub-interface tag overhead ([#269](https://github.com/ramrode/osvbng/issues/269)) ([45a97e8](https://github.com/ramrode/osvbng/commit/45a97e800ee5d1d07f2773f9c07ac96019b95925))
* **southbound:** use correct VPP sub-interface flags for BNG SVLANs ([#264](https://github.com/ramrode/osvbng/issues/264)) ([1bb9c66](https://github.com/ramrode/osvbng/commit/1bb9c666b9ca143eeb2c02578d0b37652be063b9))
* **subscriber:** count dual-stack sessions by address presence and fix in-memory cache scan ([#79](https://github.com/ramrode/osvbng/issues/79)) ([1fb6733](https://github.com/ramrode/osvbng/commit/1fb67339cb3aa712cf2b4781056f765e46b06037))
* **test:** harden 23-radius-coa suite reliability ([#283](https://github.com/ramrode/osvbng/issues/283)) ([f16b388](https://github.com/ramrode/osvbng/commit/f16b388a8dd1e930734d0f855fda59c186917ee2))
* **upgrade:** restart vpp.service when apply or rollback swaps a plugin ([#334](https://github.com/ramrode/osvbng/issues/334)) ([976343f](https://github.com/ramrode/osvbng/commit/976343f99cbcfaa9889c3c2df90eb3d2e4133932))
* **upgrade:** stamp daemon version in state file + order shutdown so tracker can't clobber "stopping" ([#338](https://github.com/ramrode/osvbng/issues/338)) ([ad59986](https://github.com/ramrode/osvbng/commit/ad59986bb256e2742ee36ea07c6892af7638ce84))
* **upgrade:** tolerate "stopping" state and reset stall window on daemon restart ([#336](https://github.com/ramrode/osvbng/issues/336)) ([f0fc452](https://github.com/ramrode/osvbng/commit/f0fc452afd0862cb3d0b0b1549041ee7155b8703))


### Performance Improvements

* **dhcp:** replace gopacket with binary parsers, fix dataplane config generation ([#216](https://github.com/ramrode/osvbng/issues/216)) ([a3992d2](https://github.com/ramrode/osvbng/commit/a3992d25c1e38fe1a730038c836631bf541f9d1b))
* **dhcpv6:** revert IPoE mutex additions, use lightweight subscriber cache merge ([#224](https://github.com/ramrode/osvbng/issues/224)) ([ef38154](https://github.com/ramrode/osvbng/commit/ef38154d83cf16b13f46f3ddedcfb27f1a486b66))
* **ipoe:** fix DHCPv6 scale bottleneck at 32k dual-stack subscribers ([#244](https://github.com/ramrode/osvbng/issues/244)) ([b3e9ede](https://github.com/ramrode/osvbng/commit/b3e9ede72f3be0778d91d64c3ca065def8dc1bc9))
* **sessions:** improve IPoE session setup throughput ([#213](https://github.com/ramrode/osvbng/issues/213)) ([7b6cab7](https://github.com/ramrode/osvbng/commit/7b6cab7f27ae539a1bc83e21cee73bc3f2f82668))
* **sessions:** profile-guided session setup optimizations ([#215](https://github.com/ramrode/osvbng/issues/215)) ([a957aa8](https://github.com/ramrode/osvbng/commit/a957aa8db8fc8a372b181608a8d5e40ef27a3219))


### Reverts

* worker pin, identical cores across labs concentrate contention ([#467](https://github.com/ramrode/osvbng/issues/467)) ([ef3e87b](https://github.com/ramrode/osvbng/commit/ef3e87b00dd6c182f96da4874a6a532292d413c0))

## [0.16.0](https://github.com/veesix-networks/osvbng/compare/v0.15.0...v0.16.0) (2026-07-21)


### Features

* **dataplane:** various CGNAT, QoS and VPP tweaks ([#394](https://github.com/veesix-networks/osvbng/issues/394)) ([07c5017](https://github.com/veesix-networks/osvbng/commit/07c50173494499570897d20db0b47782f86c8c13))


### Bug Fixes

* **ipoe,pppoe:** apply service-group qos/acl/urpf bindings on fresh session bring-up ([e6bfb0b](https://github.com/veesix-networks/osvbng/commit/e6bfb0b0d0c30d1fe64759f154584f9600ee7191))
* **ipoe:** preserve subscriber ActivatedAt across DHCP renewals ([#390](https://github.com/veesix-networks/osvbng/issues/390)) ([c4ba0bf](https://github.com/veesix-networks/osvbng/commit/c4ba0bf724d001fbaf7956913974f379ea56ddbb))

## [0.15.0](https://github.com/veesix-networks/osvbng/compare/v0.14.0...v0.15.0) (2026-06-28)


### Features

* **aaa:** standardize radius nas-port/nas-port-id and extend accounting attributes ([#380](https://github.com/veesix-networks/osvbng/issues/380)) ([426970e](https://github.com/veesix-networks/osvbng/commit/426970e277d6d15dccf640427043a4d27c5e9003))
* **ipoe:** periodic unsolicited RAs so subscriber default routes don't expire ([#384](https://github.com/veesix-networks/osvbng/issues/384)) ([83b0330](https://github.com/veesix-networks/osvbng/commit/83b033028c5b9471d0cc565327160718ec465348))
* **pppoe:** ipv6 ra/nd + dhcpv6 over ppp; derive af_packet rx-queues from workers ([#385](https://github.com/veesix-networks/osvbng/issues/385)) ([84dc880](https://github.com/veesix-networks/osvbng/commit/84dc8807268712419cffe2355d34b3cb7b500f9b))


### Bug Fixes

* **aaa:** stop policy username expansion from gating local auth ([aec07e3](https://github.com/veesix-networks/osvbng/commit/aec07e3bd94683cc72079035ccc61e10dd61eee5))
* **aaa:** stop policy username expansion from gating local auth ([#388](https://github.com/veesix-networks/osvbng/issues/388)) ([40eda4e](https://github.com/veesix-networks/osvbng/commit/40eda4e765206c4077a24ef69f70586e1a512ae7))
* **cgnat:** rebuild plugin so with thread-safe shared session pool ([c7e96d7](https://github.com/veesix-networks/osvbng/commit/c7e96d7fbf2a6d6c01f5e230fc66b0e857f19b73))
* **cgnat:** rebuild plugin so with thread-safe shared session pool ([#386](https://github.com/veesix-networks/osvbng/issues/386)) ([2864846](https://github.com/veesix-networks/osvbng/commit/2864846de4e2452263d5f19feaa38637c9a8310a))
* **cli:** stabilise show output column order and drop dead gRPC command framework ([#389](https://github.com/veesix-networks/osvbng/issues/389)) ([d694f98](https://github.com/veesix-networks/osvbng/commit/d694f98fba80e30946c3395aa7ba3126113e0edf))
* **ipoe:** advertise subscriber RA prefix off-link by default ([#382](https://github.com/veesix-networks/osvbng/issues/382)) ([dc2e34d](https://github.com/veesix-networks/osvbng/commit/dc2e34dc8d81f01758645df66b24803ea6ecafbf))

## [0.14.0](https://github.com/veesix-networks/osvbng/compare/v0.13.0...v0.14.0) (2026-06-16)


### Features

* **aaa:** hard-fail when policy-expanded RADIUS User-Name is empty ([#360](https://github.com/veesix-networks/osvbng/issues/360)) ([1db1230](https://github.com/veesix-networks/osvbng/commit/1db12303029c9003d9d2d721c15656224d2626fe))
* **aaa:** per-policy placeholder password for DHCP/IPoE RADIUS Access-Request ([#353](https://github.com/veesix-networks/osvbng/issues/353)) ([96484da](https://github.com/veesix-networks/osvbng/commit/96484daf7aa44653c2c1aaf38b692cf716645c89))
* **aaa:** populate RADIUS Acct octets/packets via cached interface stats ([#368](https://github.com/veesix-networks/osvbng/issues/368)) ([8d76b83](https://github.com/veesix-networks/osvbng/commit/8d76b83c33af96e385c7e5bf4be78701c99585bd))
* **api:** osvbngcli reaches the daemon via Unix domain socket ([#371](https://github.com/veesix-networks/osvbng/issues/371)) ([3dea4b5](https://github.com/veesix-networks/osvbng/commit/3dea4b5fe59857f88a6c7f1b0d13aa8c5d157ba7))
* **api:** per-VRF multi-listener northbound API ([#372](https://github.com/veesix-networks/osvbng/issues/372)) ([ca37793](https://github.com/veesix-networks/osvbng/commit/ca377936b402ce074fa4c6b3f3883f5e363f06f2))
* **cgnat:** multi-worker datapath ([#362](https://github.com/veesix-networks/osvbng/issues/362)) ([05ed620](https://github.com/veesix-networks/osvbng/commit/05ed620799ff7f943bb301ee6fc8d83510423112))
* **cgnat:** session dump API with inside/outside/remote ip+port+proto filters ([#361](https://github.com/veesix-networks/osvbng/issues/361)) ([fabc43c](https://github.com/veesix-networks/osvbng/commit/fabc43c5418c4b03c9edb2461d3525c8f49a7956))
* **deploy:** rewrite KVM deploy script for VPP host tuning and NUMA-aware pinning ([#365](https://github.com/veesix-networks/osvbng/issues/365)) ([88064b7](https://github.com/veesix-networks/osvbng/commit/88064b73d6b335f3fc85fef0f8f38651bcef3d75))
* **ipoe:** C-VLAN gating when non-any is used for QinQ configurations ([#351](https://github.com/veesix-networks/osvbng/issues/351)) ([bd0fdfb](https://github.com/veesix-networks/osvbng/commit/bd0fdfb9d06b0b48529b4ef8e354a4bb299059ec))
* **ipoe:** gate v4/v6 ingress on subscriber-group profile presence ([#355](https://github.com/veesix-networks/osvbng/issues/355)) ([3fb4692](https://github.com/veesix-networks/osvbng/commit/3fb469214c976f97a10d0a0911002bb2a1c0184b))
* **pppoe:** subscriber-group C-VLAN matching (parity with IPoE) ([#352](https://github.com/veesix-networks/osvbng/issues/352)) ([d6fe72c](https://github.com/veesix-networks/osvbng/commit/d6fe72cc389047de5db2d1bee265e3bcb6fc6a2b))
* **qa:** vrnetlab + Robot harness gain QEMU-mode parity for clab integration suites ([#378](https://github.com/veesix-networks/osvbng/issues/378)) ([8b9d54f](https://github.com/veesix-networks/osvbng/commit/8b9d54f1eb2d4f10767ef01ceb07640b4fe05860))
* **routing:** per-VRF OSPFv2 + OSPFv3 instances ([#348](https://github.com/veesix-networks/osvbng/issues/348)) ([6975cf4](https://github.com/veesix-networks/osvbng/commit/6975cf44ca7174a2f694d833229ed4f819779e43))
* **upgrade:** tier-a v2 upgrade pipeline + QEMU test infrastructure ([#375](https://github.com/veesix-networks/osvbng/issues/375)) ([539d09b](https://github.com/veesix-networks/osvbng/commit/539d09b20259d344e739561b2c9935766f0e8fce))


### Bug Fixes

* **aaa:** hash session-id to deterministic accounting bucket ([#358](https://github.com/veesix-networks/osvbng/issues/358)) ([34da0d4](https://github.com/veesix-networks/osvbng/commit/34da0d478c6ff40fd0323ec72e390d774b307a30))
* **bgp:** drop direct vtysh from instance Apply so ASN changes survive osvbngd restart ([#377](https://github.com/veesix-networks/osvbng/issues/377)) ([580d240](https://github.com/veesix-networks/osvbng/commit/580d240253ff30cdf8d2bc0c36b9dd6487674f2a))
* **bgp:** guard nil values in instance Apply/Rollback to prevent commit-rollback panic ([215ae5a](https://github.com/veesix-networks/osvbng/commit/215ae5a0f1a9c6da9656edea92e17596e34b7235))
* **cgnat:** prevent gateway destination mapping via in2out vpp node ([#359](https://github.com/veesix-networks/osvbng/issues/359)) ([a0b8a5d](https://github.com/veesix-networks/osvbng/commit/a0b8a5d22cec131c204a851944a1011bc71fdb09))
* **dataplane:** skip sw_interface_set_table when interface already in target VRF table ([b2f6f0f](https://github.com/veesix-networks/osvbng/commit/b2f6f0fd51c9f9419614a0c1773e812b0d9f61ad))
* **install:** bind frr.service to the dataplane netns on the qemu image ([#369](https://github.com/veesix-networks/osvbng/issues/369)) ([0415aa5](https://github.com/veesix-networks/osvbng/commit/0415aa5d5f146f0b75dc27280162747fb053236a))
* **interfaces:** re-enabling a subinterface no longer leaves it down ([#376](https://github.com/veesix-networks/osvbng/issues/376)) ([321fb0c](https://github.com/veesix-networks/osvbng/commit/321fb0caebdea45970186c4e81ce8fb33ca1b1d2))
* **ipoe:** ignore foreign DHCPv4 server responses in server mode ([#370](https://github.com/veesix-networks/osvbng/issues/370)) ([8fd58d4](https://github.com/veesix-networks/osvbng/commit/8fd58d4b0cf0bd1038ffb73500ba771d562081b5))
* **ipoe:** make runtime session expiry lease-aware ([#354](https://github.com/veesix-networks/osvbng/issues/354)) ([7b2fea2](https://github.com/veesix-networks/osvbng/commit/7b2fea28c00b672744eea2ab01ef6872298fa61f))
* **ipoe:** reserve restored session addresses in allocator on opdb restore ([#363](https://github.com/veesix-networks/osvbng/issues/363)) ([16cb21e](https://github.com/veesix-networks/osvbng/commit/16cb21ed31b54d3f6912271f69ded03ea2244dad))
* **ipoe:** subscriber-group matching ignores inner C-VLAN; gate via precomputed S/C-VLAN index ([#350](https://github.com/veesix-networks/osvbng/issues/350)) ([0e85a04](https://github.com/veesix-networks/osvbng/commit/0e85a0414a074a9927078aef408ddcfca993209d))
* **osvbng:** startup and config re-apply resilience (radius dial, bgp rollback, vrf table) ([#357](https://github.com/veesix-networks/osvbng/issues/357)) ([424e2d0](https://github.com/veesix-networks/osvbng/commit/424e2d09ad80243a09917763b3b403fadb1c2670))
* **radius:** defer server dial so an unreachable RADIUS cannot block startup ([2259ce6](https://github.com/veesix-networks/osvbng/commit/2259ce64eb9a56ed5ed39c49601c53c877f4f820))

## [0.13.0](https://github.com/veesix-networks/osvbng/compare/v0.12.1...v0.13.0) (2026-05-28)


### Features

* **cgnat:** restart-idempotent reconciler with active-mapping preflight gate ([#341](https://github.com/veesix-networks/osvbng/issues/341)) ([d4ff1f6](https://github.com/veesix-networks/osvbng/commit/d4ff1f69f3aaf056fe08a36d692168c39effd2e7))
* **dhcp:** per-pool DHCPv4 + DHCPv6 vendor options ([#347](https://github.com/veesix-networks/osvbng/issues/347)) ([db1a5fb](https://github.com/veesix-networks/osvbng/commit/db1a5fb6d46378a8534661b2de4d0547254bc5d3))
* **routing:** expose authentication on bgp, ospfv2, ospfv3 ([#345](https://github.com/veesix-networks/osvbng/issues/345)) ([15b2641](https://github.com/veesix-networks/osvbng/commit/15b26411f09812a0aa3e523b8aa6dae2362d0fd0))
* **upgrade:** add osvbngcli upgrade builtin (Tier A file-swap) ([#332](https://github.com/veesix-networks/osvbng/issues/332)) ([321e942](https://github.com/veesix-networks/osvbng/commit/321e9422b2af777de1b7a5df515dfcb9c29beb24))


### Bug Fixes

* **cgnat:** skip orphan child cleanup for pools dropped via cascade-delete ([#343](https://github.com/veesix-networks/osvbng/issues/343)) ([99ae909](https://github.com/veesix-networks/osvbng/commit/99ae9091873cc892b6309fd0e6b47f9b269fe607))
* **component:** AllReady skips plugin components, matching WaitReady semantics ([#335](https://github.com/veesix-networks/osvbng/issues/335)) ([74cc66c](https://github.com/veesix-networks/osvbng/commit/74cc66ca61a8350efc3c80a3271cb0d9df1f576b))
* **netbind:** open VRF-bound sockets in the LCP netns ([#346](https://github.com/veesix-networks/osvbng/issues/346)) ([ba6b1a3](https://github.com/veesix-networks/osvbng/commit/ba6b1a38949b0e59bf6e1ba4aa6314ec9e5268f2))
* **opdb:** apply SQLite pragmas via DSN to all pool connections + retry on busy ([#330](https://github.com/veesix-networks/osvbng/issues/330)) ([0e00fd3](https://github.com/veesix-networks/osvbng/commit/0e00fd30535193aa934c8315c0bc29eb026e7381))
* **southbound:** make bond create idempotent on osvbngd restart ([#333](https://github.com/veesix-networks/osvbng/issues/333)) ([4b7301e](https://github.com/veesix-networks/osvbng/commit/4b7301e811bb6d723a8f47efd831d37968662796))
* **upgrade:** restart vpp.service when apply or rollback swaps a plugin ([#334](https://github.com/veesix-networks/osvbng/issues/334)) ([976343f](https://github.com/veesix-networks/osvbng/commit/976343f99cbcfaa9889c3c2df90eb3d2e4133932))
* **upgrade:** stamp daemon version in state file + order shutdown so tracker can't clobber "stopping" ([#338](https://github.com/veesix-networks/osvbng/issues/338)) ([ad59986](https://github.com/veesix-networks/osvbng/commit/ad59986bb256e2742ee36ea07c6892af7638ce84))
* **upgrade:** tolerate "stopping" state and reset stall window on daemon restart ([#336](https://github.com/veesix-networks/osvbng/issues/336)) ([f0fc452](https://github.com/veesix-networks/osvbng/commit/f0fc452afd0862cb3d0b0b1549041ee7155b8703))

## [0.12.1](https://github.com/veesix-networks/osvbng/compare/v0.12.0...v0.12.1) (2026-05-25)


### Bug Fixes

* **southbound:** populate ifmgr.Interface.OuterTPID on CreateSubinterface ([#328](https://github.com/veesix-networks/osvbng/issues/328)) ([5765716](https://github.com/veesix-networks/osvbng/commit/5765716a1ecc45c071514f084c91e19008fddada))

## [0.12.0](https://github.com/veesix-networks/osvbng/compare/v0.11.0...v0.12.0) (2026-05-25)


### Features

* **bgp:** add BGP L3VPN show command coverage ([#316](https://github.com/veesix-networks/osvbng/issues/316)) ([15bf36d](https://github.com/veesix-networks/osvbng/commit/15bf36d10e775ee8d46713c99883e81a5e6ab550))
* **bgp:** add BGP unicast show command coverage ([#315](https://github.com/veesix-networks/osvbng/issues/315)) ([844335c](https://github.com/veesix-networks/osvbng/commit/844335c832c87accd106b262dede6f9cd2dab896))
* **ipoe:** emit RA + DHCPv6 + NA from link-local source per RFC 4861 ([#325](https://github.com/veesix-networks/osvbng/issues/325)) ([6089d8a](https://github.com/veesix-networks/osvbng/commit/6089d8a546f7f12e64150871b67f6b564b97c484))
* **ldp:** finalize LDP show command coverage ([#317](https://github.com/veesix-networks/osvbng/issues/317)) ([15612cf](https://github.com/veesix-networks/osvbng/commit/15612cf16b4ad82a3a9f9331630b1c44fff4025a))
* **ospf6:** add OSPFv3 show command coverage ([#313](https://github.com/veesix-networks/osvbng/issues/313)) ([06dec0b](https://github.com/veesix-networks/osvbng/commit/06dec0be06927f22358fefeebef07bacce79825b))
* **ospf:** add instance, interface, neighbor, lsdb, mpls-te show handlers ([#311](https://github.com/veesix-networks/osvbng/issues/311)) ([d4c955d](https://github.com/veesix-networks/osvbng/commit/d4c955d7dccf6cb2e71f3141d5819e377ac2a734))
* **pppoe,ipoe:** unify session recovery via setupSession SDK ([#324](https://github.com/veesix-networks/osvbng/issues/324)) ([e5487e3](https://github.com/veesix-networks/osvbng/commit/e5487e380218d66eb7cb6f87f794239a54adf6bc))
* **routing:** consolidate label rename, .all telemetry, and ipv6 lab ([#318](https://github.com/veesix-networks/osvbng/issues/318)) ([7fb6983](https://github.com/veesix-networks/osvbng/commit/7fb6983f26f90cbe214a585a0b5fe73126c2a19d))
* **routing:** expose VPP FIB + FRR zebra RIB via protocols.fib.* and protocols.zebra.* paths ([#321](https://github.com/veesix-networks/osvbng/issues/321)) ([766d2b9](https://github.com/veesix-networks/osvbng/commit/766d2b9bb538382cd6b845bbcf2fe19fb0037362))


### Bug Fixes

* **autoconfig:** gate subscriber IPv6 derivation on VRF IPv6 unicast ([#314](https://github.com/veesix-networks/osvbng/issues/314)) ([5bb681b](https://github.com/veesix-networks/osvbng/commit/5bb681b1cc7aa356516dac3d86f23ab532664979))
* **configmgr:** skip no-op changes, suppress empty versions on restart, render values as JSON ([#320](https://github.com/veesix-networks/osvbng/issues/320)) ([b0046c5](https://github.com/veesix-networks/osvbng/commit/b0046c5c3fd4f9a9320c8043e85f9d5ac410e61a))
* **show, osvbngcli:** typed BoolOption, fast exit on all paths, hide .all paths, scope flag help, render wildcard descriptions ([#319](https://github.com/veesix-networks/osvbng/issues/319)) ([6c758ff](https://github.com/veesix-networks/osvbng/commit/6c758ffd24427cfcccc314a6ee848aef71d0dd5d))

## [0.11.0](https://github.com/veesix-networks/osvbng/compare/v0.10.0...v0.11.0) (2026-05-18)


### Features

* **api:** paginate list-returning northbound show endpoints ([#292](https://github.com/veesix-networks/osvbng/issues/292)) ([082fa32](https://github.com/veesix-networks/osvbng/commit/082fa32d3a21cc567d6912b2b9c0b00096f0d6e5))
* **dev:** one-shot QEMU/KVM development environment ([#251](https://github.com/veesix-networks/osvbng/issues/251)) ([3a90cf8](https://github.com/veesix-networks/osvbng/commit/3a90cf831c5fab18f3e07c6e3a6d8399cfbc9de7))
* **dhcp:** support RFC 6221 LDRA termination in local DHCPv6 provider ([#288](https://github.com/veesix-networks/osvbng/issues/288)) ([110e55e](https://github.com/veesix-networks/osvbng/commit/110e55e487066cd4197fc3c4ca9c9b7c8264f5e2))
* **dhcp:** VRF support for DHCP relay/proxy ([#296](https://github.com/veesix-networks/osvbng/issues/296)) ([1d6cade](https://github.com/veesix-networks/osvbng/commit/1d6cade928061cb1c4b387c830b51a1749df4710))
* **ha:** per-component VRF binding for HA peer-sync + gateway via pkg/netbind ([#295](https://github.com/veesix-networks/osvbng/issues/295)) ([def589c](https://github.com/veesix-networks/osvbng/commit/def589caec1d23aa5aca31d599d41bce31fff35d))
* **l2tp:** L2TPv2 LAC and LNS (RFC 2661) ([#305](https://github.com/veesix-networks/osvbng/issues/305)) ([5346ccf](https://github.com/veesix-networks/osvbng/commit/5346ccf02942f19b0b4be0d8db4a6e663662ea40))
* **monitoring:** introduce typed telemetry SDK and migrated VPP exporter ([#299](https://github.com/veesix-networks/osvbng/issues/299)) ([1a1d7c6](https://github.com/veesix-networks/osvbng/commit/1a1d7c636426bc2262c5d9c28f90db384d30791e))
* **monitoring:** migrate state.RegisterMetric callers to telemetry.RegisterMetric[T] ([ba5a5d0](https://github.com/veesix-networks/osvbng/commit/ba5a5d0ec4b38388aeff60b3b84997571449ff03))
* **monitoring:** retire legacy state.RegisterMetric + typed FRR surfaces ([#304](https://github.com/veesix-networks/osvbng/issues/304)) ([05b8f95](https://github.com/veesix-networks/osvbng/commit/05b8f95025daae673cb7c04737d16e111fb8b39e))
* **monitoring:** show-driven metric registration ([#300](https://github.com/veesix-networks/osvbng/issues/300)) ([6410e82](https://github.com/veesix-networks/osvbng/commit/6410e82615921700c9d9e6db288f861ef7052bc9))
* **monitoring:** typed FRR JSON surfaces for BGP and LDP ([eee7f97](https://github.com/veesix-networks/osvbng/commit/eee7f9794432901e44e6237f644d74591667b90d))
* **monitoring:** unify RegisterMetric across show handler return shapes ([#302](https://github.com/veesix-networks/osvbng/issues/302)) ([a8548cf](https://github.com/veesix-networks/osvbng/commit/a8548cf0595192bd3a227bb40d48c9955f3dc7d4))
* **netbind:** plugin listener + HTTP client VRF binding with TLS ([#293](https://github.com/veesix-networks/osvbng/issues/293)) ([44cc084](https://github.com/veesix-networks/osvbng/commit/44cc08448ffd7469f1d41c7707848679528ce4ab))
* **plugin:** add community cgnat-http-exporter plugin ([#284](https://github.com/veesix-networks/osvbng/issues/284)) ([82a8a0d](https://github.com/veesix-networks/osvbng/commit/82a8a0d6df520ce5ce1b557bbf9dd7c49f6ffbcc))
* **radius:** per-server VRF + source IP for auth/acct/CoA via netbind ([1f3424d](https://github.com/veesix-networks/osvbng/commit/1f3424d5860a8e310ee49024b59d881106c354e9))
* **radius:** per-server VRF + source IP for auth/acct/CoA via pkg/netbind ([#294](https://github.com/veesix-networks/osvbng/issues/294)) ([24265a9](https://github.com/veesix-networks/osvbng/commit/24265a99304b32c775091caf8d1c9668d1a9f6d2))
* **subscriber:** mixed IPoE+PPPoE on shared S-VLAN range ([#306](https://github.com/veesix-networks/osvbng/issues/306)) ([e4ed0ad](https://github.com/veesix-networks/osvbng/commit/e4ed0ad4fc683b6cfd0988737b7b93ad1ca85f99))
* **subscriber:** unset unnumbered on session release ([#307](https://github.com/veesix-networks/osvbng/issues/307)) ([798a264](https://github.com/veesix-networks/osvbng/commit/798a264f4affe4bdc438468de718b7765fd4aeff))
* **vrf:** subscriber VRF cascade + VRF-lite / L3VPN integration suites ([#289](https://github.com/veesix-networks/osvbng/issues/289)) ([05b5dda](https://github.com/veesix-networks/osvbng/commit/05b5ddaa105443dadcba89bd35f06d9c9d437289))


### Bug Fixes

* **monitoring:** drop hostname label from BGP neighbor/VPN peer ([918af4b](https://github.com/veesix-networks/osvbng/commit/918af4baaf4ec03c2cb9ae42141d7513b2f9c0b8))
* **netbind:** pick LocalAddr family by dial network for UDP source pinning ([c8ad5e4](https://github.com/veesix-networks/osvbng/commit/c8ad5e44ef6975a09fb1159cac3706743f211c55))
* **southbound:** BNG parent promisc when Linux link missing ([#301](https://github.com/veesix-networks/osvbng/issues/301)) ([fa947d6](https://github.com/veesix-networks/osvbng/commit/fa947d6dfedc2fe5c59d28d5ae33623d38e1d88e))

## [0.10.0](https://github.com/veesix-networks/osvbng/compare/v0.9.0...v0.10.0) (2026-04-13)


### Features

* **pppoe:** add TCP MSS clamping with PPP MRU configuration (RFC 4638) ([#279](https://github.com/veesix-networks/osvbng/issues/279)) ([4f7e041](https://github.com/veesix-networks/osvbng/commit/4f7e0418cbde2fd106dd27dc5b22c7b497b2eac9))
* **radius:** add RADIUS CoA/Disconnect-Message with subscriber runtime mutation ([#282](https://github.com/veesix-networks/osvbng/issues/282)) ([b2797b6](https://github.com/veesix-networks/osvbng/commit/b2797b6d5388187fed89d144ed0e4ee4324fd172))
* **subscriber:** add plugin-agnostic subscriber runtime mutation API ([#281](https://github.com/veesix-networks/osvbng/issues/281)) ([7bf5a0d](https://github.com/veesix-networks/osvbng/commit/7bf5a0d33e540334e5a6862adab6df3ac3d81899))


### Bug Fixes

* **test:** harden 23-radius-coa suite reliability ([#283](https://github.com/veesix-networks/osvbng/issues/283)) ([9e5c399](https://github.com/veesix-networks/osvbng/commit/9e5c3999d1ce5dae8ce730d8616bcc8011b40dd2))

## [0.9.0](https://github.com/veesix-networks/osvbng/compare/v0.8.0...v0.9.0) (2026-04-10)


### Features

* **api:** complete typed metadata for all show, oper, and conf handlers ([#275](https://github.com/veesix-networks/osvbng/issues/275)) ([127d80a](https://github.com/veesix-networks/osvbng/commit/127d80ab4ce4fe1756052d4dd18ba308e7551a2c))
* **config:** rename vlan-protocol to vlan-tpid with dot1ad default for QinQ ([#266](https://github.com/veesix-networks/osvbng/issues/266)) ([2aebf7d](https://github.com/veesix-networks/osvbng/commit/2aebf7deb8c15f3ae5f7592823f30ac6ca3ed160))


### Bug Fixes

* **southbound:** auto-create LCP pairs, remove lcp config option ([#271](https://github.com/veesix-networks/osvbng/issues/271)) ([2ec4c85](https://github.com/veesix-networks/osvbng/commit/2ec4c8597d475d82b7f28db8ce4e9addbcfd0024))
* **southbound:** set HW MTU, host interface MTU, and QinQ sub-interface tag overhead ([#269](https://github.com/veesix-networks/osvbng/issues/269)) ([16e0071](https://github.com/veesix-networks/osvbng/commit/16e007163ca8f459524b98206c26e904eb5cc793))

## [0.8.0](https://github.com/veesix-networks/osvbng/compare/v0.7.0...v0.8.0) (2026-04-04)


### Features

* **dataplane:** support bond/LACP interfaces ([#32](https://github.com/veesix-networks/osvbng/issues/32)) ([#252](https://github.com/veesix-networks/osvbng/issues/252)) ([ba91745](https://github.com/veesix-networks/osvbng/commit/ba91745b4b0c0a1722afe15becbb5fe1a7bad7e5))
* **routing:** add routing policy framework with prefix-sets, community-sets, AS-path-sets, and route-policies ([#263](https://github.com/veesix-networks/osvbng/issues/263)) ([65708f5](https://github.com/veesix-networks/osvbng/commit/65708f55a17363946685b13127ff1fa73f671fc3))
* **show:** add show interfaces framework ([#254](https://github.com/veesix-networks/osvbng/issues/254)) ([ca5fa64](https://github.com/veesix-networks/osvbng/commit/ca5fa64eb5f176c07884cb97c618440a8c3568a8))
* **southbound:** add explicit sub-interface support ([#259](https://github.com/veesix-networks/osvbng/issues/259)) ([3a4dbc1](https://github.com/veesix-networks/osvbng/commit/3a4dbc132f4d9945982a75d431aea9bf5ab560ab))


### Bug Fixes

* **southbound:** call VPP SwInterfaceSetMtu API in SetInterfaceMTU ([#258](https://github.com/veesix-networks/osvbng/issues/258)) ([9b22c91](https://github.com/veesix-networks/osvbng/commit/9b22c91380c0deba9eed2a080e75a19b17a6f95f))
* **southbound:** use correct VPP sub-interface flags for BNG SVLANs ([#264](https://github.com/veesix-networks/osvbng/issues/264)) ([d86235c](https://github.com/veesix-networks/osvbng/commit/d86235c9dc248619bb4513fb83b916e63bb6cb63))

## [0.7.0](https://github.com/veesix-networks/osvbng/compare/v0.6.1...v0.7.0) (2026-03-31)


### Features

* **component:** add readiness signaling for async plugin startup ([#221](https://github.com/veesix-networks/osvbng/issues/221)) ([920bee1](https://github.com/veesix-networks/osvbng/commit/920bee19f9e02d9b662975b13fbf9ef3eb83e8cf))
* **dataplane:** cgroup-aware CPU detection with conservative defaults ([#237](https://github.com/veesix-networks/osvbng/issues/237)) ([5ab23e2](https://github.com/veesix-networks/osvbng/commit/5ab23e215c72aea57e745efcb709d70199b18c01))
* **ha:** add GARP flood on SRG promotion with batching and rate limiting ([#225](https://github.com/veesix-networks/osvbng/issues/225)) ([f010a8b](https://github.com/veesix-networks/osvbng/commit/f010a8bc285a219c50b67ade0b2f1b5deb101205))
* **logger:** async zerolog migration for non-blocking logging ([#217](https://github.com/veesix-networks/osvbng/issues/217)) ([000a879](https://github.com/veesix-networks/osvbng/commit/000a87976eb2a98a7b98105358a5f540603925f1))


### Bug Fixes

* **ci:** add topology cleanup and diagnostics to test workflows ([#234](https://github.com/veesix-networks/osvbng/issues/234)) ([bb71adf](https://github.com/veesix-networks/osvbng/commit/bb71adf824ff2f765c8f779e4ad05abf7f4d6638))
* **ci:** handle non-zero exit codes in test setup version checks ([#235](https://github.com/veesix-networks/osvbng/issues/235)) ([00a2702](https://github.com/veesix-networks/osvbng/commit/00a2702709a0c42fe946284f9af1da73089d18b0))
* **ci:** pre-create containerlab network to prevent parallel deploy race ([#236](https://github.com/veesix-networks/osvbng/issues/236)) ([2303393](https://github.com/veesix-networks/osvbng/commit/23033932034dc04439412b2dc93896bc1d143cd7))
* **dataplane:** default to exactly 1 thread worker regardless of available cores ([#238](https://github.com/veesix-networks/osvbng/issues/238)) ([ed73678](https://github.com/veesix-networks/osvbng/commit/ed736780dd43677d9317d604307941d2fe72b1b8))
* **dataplane:** default VPP heap to 1G and increase egress channel capacity for DPDK scale ([#250](https://github.com/veesix-networks/osvbng/issues/250)) ([8cc78e5](https://github.com/veesix-networks/osvbng/commit/8cc78e56a917937bcf926d3408e8b0b8f0b1a309))
* **dhcpv6:** resolve session cache race between concurrent IPv4 and IPv6 lifecycle events ([#222](https://github.com/veesix-networks/osvbng/issues/222)) ([62c06a8](https://github.com/veesix-networks/osvbng/commit/62c06a89ad724836e26f0b006d64eed551b0ee55))
* **ha:** prevent standby from responding to PPPoE discovery ([#227](https://github.com/veesix-networks/osvbng/issues/227)) ([4d7c570](https://github.com/veesix-networks/osvbng/commit/4d7c5702366c7d83958bbf985966461dfb538115))
* **ha:** prevent standby from responding to subscriber ARP and IPv6 ND ([#226](https://github.com/veesix-networks/osvbng/issues/226)) ([89c84f4](https://github.com/veesix-networks/osvbng/commit/89c84f4f5086732a2776c93f94a4c8ba4a95e8d0))
* **ha:** update GetVirtualMAC test to reflect active-state guard ([#229](https://github.com/veesix-networks/osvbng/issues/229)) ([01a5024](https://github.com/veesix-networks/osvbng/commit/01a50243fdd9617edee377922a8538951fe9fb71))
* **ipoe:** eliminate session index race causing 1-in-48k stuck session ([#248](https://github.com/veesix-networks/osvbng/issues/248)) ([f79dd45](https://github.com/veesix-networks/osvbng/commit/f79dd45ec2b5ecda389c93fb8c84e2f1039a5b24))
* **southbound:** replace FIFO async worker with stream pool and parameterize VPP memory config ([#246](https://github.com/veesix-networks/osvbng/issues/246)) ([b51a6be](https://github.com/veesix-networks/osvbng/commit/b51a6be721ee19ed3e36ea556105974f78b9cd2a))


### Performance Improvements

* **dhcp:** replace gopacket with binary parsers, fix dataplane config generation ([#216](https://github.com/veesix-networks/osvbng/issues/216)) ([c68e95a](https://github.com/veesix-networks/osvbng/commit/c68e95a8afd7d88e07c96643eb16ccd30c5228dd))
* **dhcpv6:** revert IPoE mutex additions, use lightweight subscriber cache merge ([#224](https://github.com/veesix-networks/osvbng/issues/224)) ([5da7514](https://github.com/veesix-networks/osvbng/commit/5da751447ac8ad2bab37eabd4c1b688e0c7e983b))
* **ipoe:** fix DHCPv6 scale bottleneck at 32k dual-stack subscribers ([#244](https://github.com/veesix-networks/osvbng/issues/244)) ([abf37ed](https://github.com/veesix-networks/osvbng/commit/abf37edeec1ceb5bfe03ad84043a0f4d91fe3b87))
* **sessions:** improve IPoE session setup throughput ([#213](https://github.com/veesix-networks/osvbng/issues/213)) ([0a0b97d](https://github.com/veesix-networks/osvbng/commit/0a0b97d36686fdc5850f528f0d6022f2db190155))
* **sessions:** profile-guided session setup optimizations ([#215](https://github.com/veesix-networks/osvbng/issues/215)) ([20fa869](https://github.com/veesix-networks/osvbng/commit/20fa8690b75e17eb3d95c356e96387b9521a34a1))

## [0.6.1](https://github.com/veesix-networks/osvbng/compare/v0.6.0...v0.6.1) (2026-03-22)


### Bug Fixes

* **ci:** trigger discord webhook on release-please PR creation ([#212](https://github.com/veesix-networks/osvbng/issues/212)) ([81299b3](https://github.com/veesix-networks/osvbng/commit/81299b37649682ea6b5b4a18ea887b3f5ed2640d))
* **dataplane:** default AF_PACKET interfaces to interrupt rx-mode ([#209](https://github.com/veesix-networks/osvbng/issues/209)) ([47617ec](https://github.com/veesix-networks/osvbng/commit/47617ece50aa18e63a77aa2a70ea3e7edc7bbc5d))
* **dataplane:** use poll instead of busy-read on punt eventfd ([#211](https://github.com/veesix-networks/osvbng/issues/211)) ([55be33a](https://github.com/veesix-networks/osvbng/commit/55be33aa5daec67996c88672eb4d431d39b4c1ab))

## [0.6.0](https://github.com/veesix-networks/osvbng/compare/v0.5.0...v0.6.0) (2026-03-21)


### Features

* **cgnat:** add Carrier-Grade NAT with PBA mode for IPoE and PPPoE subscribers ([#183](https://github.com/veesix-networks/osvbng/issues/183)) ([c96cee1](https://github.com/veesix-networks/osvbng/commit/c96cee1fb449dd98188a41baf399e4725a5b0e3a))
* **cgnat:** add CGNAT HA mapping sync with incremental and bulk sync ([#188](https://github.com/veesix-networks/osvbng/issues/188)) ([95088e4](https://github.com/veesix-networks/osvbng/commit/95088e496284216dad406c6ad81b6fb30c7df26e))
* **ha:** add tracker-driven promotion from STANDBY_ALONE ([#196](https://github.com/veesix-networks/osvbng/issues/196)) ([7a8222c](https://github.com/veesix-networks/osvbng/commit/7a8222c26e2436bfab3783bc147f029e7be02f5f))
* **ha:** restore synced sessions on HA promotion with failover tests ([#190](https://github.com/veesix-networks/osvbng/issues/190)) ([3bc1dac](https://github.com/veesix-networks/osvbng/commit/3bc1dac5debaa3b0ab5f8894a1eb50361068d2fa))
* **ha:** sync AAA attributes across HA failover with RADIUS validation ([#192](https://github.com/veesix-networks/osvbng/issues/192)) ([504c88b](https://github.com/veesix-networks/osvbng/commit/504c88b4dfbd07f011fa5f3ed4361a99fd9db597))
* **qos:** integrate CAKE scheduler plugin into subscriber lifecycle ([#206](https://github.com/veesix-networks/osvbng/issues/206)) ([dd387a5](https://github.com/veesix-networks/osvbng/commit/dd387a55bf80e1ac9ffbb625cb8be3464b7f7d5e))


### Bug Fixes

* **arp:** ignore DAD probe for client's own assigned IP ([#205](https://github.com/veesix-networks/osvbng/issues/205)) ([678f6a0](https://github.com/veesix-networks/osvbng/commit/678f6a00d50ccb5170394b423bf15c5497c8242a))
* **ci:** add checkout step for Discord changelog notification ([#193](https://github.com/veesix-networks/osvbng/issues/193)) ([ab600b5](https://github.com/veesix-networks/osvbng/commit/ab600b5b89f44fd32c81ba990e398eb316e3dd07))
* **ci:** extract PR number from release-please JSON output ([#199](https://github.com/veesix-networks/osvbng/issues/199)) ([e4a1d95](https://github.com/veesix-networks/osvbng/commit/e4a1d9545cd5e8e23d5424afbf7cc5cff7ba631b))
* **ci:** pass all github expressions as env vars to avoid shell parsing errors ([#195](https://github.com/veesix-networks/osvbng/issues/195)) ([1f7f7b4](https://github.com/veesix-networks/osvbng/commit/1f7f7b4631a466a19dc9915ff951ace21c6fe869))
* **ci:** prevent shell parsing failures in Discord webhook notifications ([#197](https://github.com/veesix-networks/osvbng/issues/197)) ([c388837](https://github.com/veesix-networks/osvbng/commit/c388837aa76674e31b5c78440e5e97b727964021))
* **ci:** use github context instead of git log for Discord notifications ([#194](https://github.com/veesix-networks/osvbng/issues/194)) ([3f03898](https://github.com/veesix-networks/osvbng/commit/3f03898c1f4580682c39d77189eec23554e78bc7))
* **ha:** handle reverse event ordering for tracker promotion ([#198](https://github.com/veesix-networks/osvbng/issues/198)) ([a6071bd](https://github.com/veesix-networks/osvbng/commit/a6071bdac531c3725bd5f73d199429b9170e0694))
* **ha:** only restore synced sessions when promoted from STANDBY_ALONE ([#201](https://github.com/veesix-networks/osvbng/issues/201)) ([2189aee](https://github.com/veesix-networks/osvbng/commit/2189aee5e4afb8770a9726a269f93506776d8f92))

## [0.5.0](https://github.com/veesix-networks/osvbng/compare/v0.4.0...v0.5.0) (2026-03-14)


### Features

* **aaa:** add RADIUS auth provider with server failover and accounting ([#169](https://github.com/veesix-networks/osvbng/issues/169)) ([ad464a3](https://github.com/veesix-networks/osvbng/commit/ad464a37f64e8494ee2de3feaa55750addc3dde7))
* **dhcp:** add relay and proxy providers with Kea dev environment, smoke tests, and CI integration ([#172](https://github.com/veesix-networks/osvbng/issues/172)) ([5b6b794](https://github.com/veesix-networks/osvbng/commit/5b6b794529c0a7834ae3d4c43e39bc4f4a13c66c))


### Bug Fixes

* **aaa:** add Message-Authenticator (attr 80) to Access-Request packets ([#181](https://github.com/veesix-networks/osvbng/issues/181)) ([3f5796e](https://github.com/veesix-networks/osvbng/commit/3f5796e4db8c4c79d28a3ac4791cc72424430c8f))
* **aaa:** address RADIUS auth/accounting issues from code review ([#174](https://github.com/veesix-networks/osvbng/issues/174)) ([28625ae](https://github.com/veesix-networks/osvbng/commit/28625aeef5d19b7a22376efb2138f7c98aa42545))
* **aaa:** use atomic pointer for global RADIUS provider reference ([#180](https://github.com/veesix-networks/osvbng/issues/180)) ([612ff45](https://github.com/veesix-networks/osvbng/commit/612ff455b7e6ec4279c89fa954840b9de35bd7a2))
* **aaa:** wire up RadiusAttr name resolution for response mappings ([#179](https://github.com/veesix-networks/osvbng/issues/179)) ([c84b64e](https://github.com/veesix-networks/osvbng/commit/c84b64e1af4fda0d7152e22da4241b8d2d2b80c3))
* **dhcp:** address relay and proxy issues ([#175](https://github.com/veesix-networks/osvbng/issues/175)) ([a24133d](https://github.com/veesix-networks/osvbng/commit/a24133d1c29da9f6d7873bbccb4209c1613f41b7))
* **dhcp:** use compound keys for pending map to prevent XID collisions ([#178](https://github.com/veesix-networks/osvbng/issues/178)) ([82366cf](https://github.com/veesix-networks/osvbng/commit/82366cf3f505ac78926420fb4bb51461e32a2630))
* **dhcpv6:** add enterprise-number prefix to Remote-ID option per RFC 4649 ([#177](https://github.com/veesix-networks/osvbng/issues/177)) ([b914624](https://github.com/veesix-networks/osvbng/commit/b9146243a05774d1db7e76e7767ca62d5b14972b))
* **dhcpv6:** use SRG virtual MAC or access interface MAC for proxy DUID ([#176](https://github.com/veesix-networks/osvbng/issues/176)) ([785851f](https://github.com/veesix-networks/osvbng/commit/785851f9ee9553c9042fcc7eb452a6864803e034))
* **ha:** promote WAITING SRGs to ACTIVE_SOLO when peer is unreachable ([#171](https://github.com/veesix-networks/osvbng/issues/171)) ([06e1f15](https://github.com/veesix-networks/osvbng/commit/06e1f15b258e436ea00c8cc0e9cad44f3cb91bc8))

## [0.4.0](https://github.com/veesix-networks/osvbng/compare/v0.3.1...v0.4.0) (2026-03-03)


### Features

* **ha:** add HA foundation with SRG state machine, gRPC peer, and component integration ([#137](https://github.com/veesix-networks/osvbng/issues/137)) ([2df141b](https://github.com/veesix-networks/osvbng/commit/2df141b1d6aeae3a6744855a91dc0baab46750ed))
* **ha:** add interface tracking, SRG counters handler, and split-brain resolution ([#141](https://github.com/veesix-networks/osvbng/issues/141)) ([d6f5c5e](https://github.com/veesix-networks/osvbng/commit/d6f5c5ec6610a169961ba5aa52aed2bf64d93260))
* **ha:** add pool-targeted sync and full bulk sync from live sessions ([#165](https://github.com/veesix-networks/osvbng/issues/165)) ([573e145](https://github.com/veesix-networks/osvbng/commit/573e145e3c48be938ad897bea2ba44a38680c441))
* **ha:** add session sync for HA standby replication ([#164](https://github.com/veesix-networks/osvbng/issues/164)) ([0b2bf44](https://github.com/veesix-networks/osvbng/commit/0b2bf447fcca42849a51ccf50661a34932d8a566))
* **ha:** add SRG BGP route advertisement and withdrawal on failover ([#142](https://github.com/veesix-networks/osvbng/issues/142)) ([1e3613f](https://github.com/veesix-networks/osvbng/commit/1e3613faec9d5ea9483a03b7c1acb09d3d801cfd))
* **ha:** add SRG dataplane abstraction with VPP implementation and no-op fallback ([#140](https://github.com/veesix-networks/osvbng/issues/140)) ([89888e5](https://github.com/veesix-networks/osvbng/commit/89888e57e600bc5ac716ff5a44cb84432779f0d7))


### Bug Fixes

* **ha:** standby does not auto-promote on peer loss ([#163](https://github.com/veesix-networks/osvbng/issues/163)) ([0693571](https://github.com/veesix-networks/osvbng/commit/0693571e0f83cea5d3c54f3cabcfc969f3d03d5b))
* **ipv6:** use separate IPv6 profile name in allocator and add PPPoE IANA allocation ([#159](https://github.com/veesix-networks/osvbng/issues/159)) ([34a6dca](https://github.com/veesix-networks/osvbng/commit/34a6dcae44cd3d811135b22b66a131ba82190c7f))
* poll for bngblaster exit instead of fixed sleep ([#157](https://github.com/veesix-networks/osvbng/issues/157)) ([76d613d](https://github.com/veesix-networks/osvbng/commit/76d613dd88c501d54ccdbcc0fac49a85da16c71d))
* poll for IPv6 readiness in session tests and enable IP6CP for PPPoE ([#158](https://github.com/veesix-networks/osvbng/issues/158)) ([9bd142d](https://github.com/veesix-networks/osvbng/commit/9bd142d45dfbc9d121cf5e9344a3b7fe88860925))

## [0.3.1](https://github.com/veesix-networks/osvbng/compare/v0.3.0...v0.3.1) (2026-02-23)


### Bug Fixes

* **docker:** create dataplane netns in published image entrypoint ([#135](https://github.com/veesix-networks/osvbng/issues/135)) ([81f8530](https://github.com/veesix-networks/osvbng/commit/81f8530199d1e21bcbc496619ffc8f918c700def))

## [0.3.0](https://github.com/veesix-networks/osvbng/compare/v0.2.0...v0.3.0) (2026-02-22)


### Features

* **watchdog:** add VPP health monitoring and dataplane recovery ([#128](https://github.com/veesix-networks/osvbng/issues/128)) ([2bd4648](https://github.com/veesix-networks/osvbng/commit/2bd4648c47bd3daeefb06333b1887d475aaddb0d))


### Bug Fixes

* **ipoe:** always call ip_table_bind for new IPoE session interfaces ([#134](https://github.com/veesix-networks/osvbng/issues/134)) ([ea47a27](https://github.com/veesix-networks/osvbng/commit/ea47a276b3fea19ec5d0e3f3b0b09188d778f6ce))
* **ipoe:** fix VPP crash on session re-creation and release resource leaks ([#133](https://github.com/veesix-networks/osvbng/issues/133)) ([339665f](https://github.com/veesix-networks/osvbng/commit/339665f69b398e0e1b76c2a263723e938e93a240))

## [0.2.0](https://github.com/veesix-networks/osvbng/compare/v0.1.2...v0.2.0) (2026-02-21)


### Features

* **aaa:** add policy-based authentication mode ([#124](https://github.com/veesix-networks/osvbng/issues/124)) ([8a1758e](https://github.com/veesix-networks/osvbng/commit/8a1758e3aaca3a6a9ce21553eb249e5dc849f8c5))
* **aaa:** add pool and service group attribute mappings ([#109](https://github.com/veesix-networks/osvbng/issues/109)) ([75237b8](https://github.com/veesix-networks/osvbng/commit/75237b8dd3be13dfdf126040282212b2f945c4a3))
* **aaa:** log returned attributes in authentication response ([#116](https://github.com/veesix-networks/osvbng/issues/116)) ([dbdd00d](https://github.com/veesix-networks/osvbng/commit/dbdd00dce7139b09ae8f72a30723aeb7a6731671))
* **bgp:** add VPNv4/VPNv6 address family config model and templates ([#97](https://github.com/veesix-networks/osvbng/issues/97)) ([8ecbeb9](https://github.com/veesix-networks/osvbng/commit/8ecbeb9fe598893fa21446c8d0eef9d3d7cfda6b))
* **bgp:** add VPNv4/VPNv6 and L3VPN configuration and show handlers ([#98](https://github.com/veesix-networks/osvbng/issues/98)) ([958a7fe](https://github.com/veesix-networks/osvbng/commit/958a7fe7ab144be76c904a9c181728f88403696e))
* **dataplane:** add LCP namespace support with routing protocol fixes ([#99](https://github.com/veesix-networks/osvbng/issues/99)) ([9823ae7](https://github.com/veesix-networks/osvbng/commit/9823ae7faac8eba3cec98c6f2693e9262a68a2e5))
* **dhcp:** add DHCP profile types and shared allocator ([#106](https://github.com/veesix-networks/osvbng/issues/106)) ([82e8bb1](https://github.com/veesix-networks/osvbng/commit/82e8bb1d89a9d2a31d67ed7f81ead3c777256f6b))
* **dhcp:** add per-VRF pool isolation to allocator registry ([#112](https://github.com/veesix-networks/osvbng/issues/112)) ([d5a7281](https://github.com/veesix-networks/osvbng/commit/d5a728177f8e8e8d1d9d0ade1adbd5abe2db8787))
* **dhcp:** add typed AAA attributes and wire DHCPv4 provisioning context ([#107](https://github.com/veesix-networks/osvbng/issues/107)) ([5be6049](https://github.com/veesix-networks/osvbng/commit/5be60490e256e960abf977a8689555e0c1c77eef))
* **dhcp:** add VRF-aware pool overflow for IPv4, IANA, and PD ([#118](https://github.com/veesix-networks/osvbng/issues/118)) ([2652cf0](https://github.com/veesix-networks/osvbng/commit/2652cf06be6bc641ba196bbe5305637067c82a1e))
* **dhcp:** centralize IP allocation in resolve layer ([#110](https://github.com/veesix-networks/osvbng/issues/110)) ([597e77b](https://github.com/veesix-networks/osvbng/commit/597e77b9e5294874a33cbe428f8532e0d5d3d316))
* **dhcpv6:** wire provisioning context through DHCPv6 provider ([#108](https://github.com/veesix-networks/osvbng/issues/108)) ([c7ef3a6](https://github.com/veesix-networks/osvbng/commit/c7ef3a6f1312d0481d19342b3f874d24a395ba91))
* **ifmgr:** track interface IP addresses and FIB table IDs ([#93](https://github.com/veesix-networks/osvbng/issues/93)) ([8cfc5f2](https://github.com/veesix-networks/osvbng/commit/8cfc5f2d2176d296b4813dff5af2f88381f3a653))
* **l3vpn:** add L3VPN dev environment with loopback-based peering ([#103](https://github.com/veesix-networks/osvbng/issues/103)) ([816f2b1](https://github.com/veesix-networks/osvbng/commit/816f2b1aed83dc476678dc62d4c85743bda5e7c9))
* **mpls:** add MPLS/LDP southbound API, config model, and FRR templates ([#96](https://github.com/veesix-networks/osvbng/issues/96)) ([5f314a0](https://github.com/veesix-networks/osvbng/commit/5f314a09be790eef7339258003ff3025168e4796))
* **qos:** implement per-subscriber policer lifecycle ([#120](https://github.com/veesix-networks/osvbng/issues/120)) ([1b6f6ca](https://github.com/veesix-networks/osvbng/commit/1b6f6caa39274084b7d4047e49ab59a214ca92a2))
* **routing:** add VRF assignment to IPoE and PPPoE subscriber sessions ([bbbb6b7](https://github.com/veesix-networks/osvbng/commit/bbbb6b789ecf9d2ec218210ce392d531174217d6))
* **routing:** add VRF manager with Linux VRF and VPP table lifecycle ([1334211](https://github.com/veesix-networks/osvbng/commit/133421170a3995f1500cb0cb60a79b4956d0f7fc))
* **routing:** add VRF manager with Linux VRF and VPP table lifecycle ([#89](https://github.com/veesix-networks/osvbng/issues/89)) ([6c43cfe](https://github.com/veesix-networks/osvbng/commit/6c43cfe476e55ad73425410efd0e764a37e44b03))
* **routing:** bind infrastructure interfaces to VRF during creation ([3b838cd](https://github.com/veesix-networks/osvbng/commit/3b838cd332210d93445ebaf2a34e5d2cb838e688))
* **routing:** wire VRF manager into application startup and config loading ([c6b9546](https://github.com/veesix-networks/osvbng/commit/c6b95467b268dab3951765415f1045da9bd98002))
* **svcgroup:** add service group resolver with three-layer merge resolution ([33f679f](https://github.com/veesix-networks/osvbng/commit/33f679fb00fed4f76409d51f21ff2ebdf81d75c6))
* **svcgroup:** added support for service groups ([aa02eb8](https://github.com/veesix-networks/osvbng/commit/aa02eb8281ab051a3413f0401646fa7cdf7113de))
* **svcgroup:** added support for service groups ([#91](https://github.com/veesix-networks/osvbng/issues/91)) ([aa02eb8](https://github.com/veesix-networks/osvbng/commit/aa02eb8281ab051a3413f0401646fa7cdf7113de))


### Bug Fixes

* **arp:** enforce VRF-aware ARP response filtering ([#94](https://github.com/veesix-networks/osvbng/issues/94)) ([bf7bb78](https://github.com/veesix-networks/osvbng/commit/bf7bb78f91415089541b9f0f2d4a01fac2a0cfbe))
* **arp:** use per-interface IP dump and ifmgr cache ([#95](https://github.com/veesix-networks/osvbng/issues/95)) ([65dea39](https://github.com/veesix-networks/osvbng/commit/65dea39741bf490da6e43bf31b27d6ec9250385c))
* **bgp:** activate neighbors in unicast address families ([#121](https://github.com/veesix-networks/osvbng/issues/121)) ([2187724](https://github.com/veesix-networks/osvbng/commit/2187724bcb4e06d6bfa793f824a836d1a02ff768))
* **bgp:** add blackhole routes for advertised pool networks ([#122](https://github.com/veesix-networks/osvbng/issues/122)) ([8c17fdf](https://github.com/veesix-networks/osvbng/commit/8c17fdf60dac1c60ca46f0e2e70b9adc8caec6d3))
* **bgp:** add no bgp default ipv4-unicast to template ([#101](https://github.com/veesix-networks/osvbng/issues/101)) ([9574dcf](https://github.com/veesix-networks/osvbng/commit/9574dcfbf0428da99822b5ac54454c5e81ae1878))
* **config:** stabilize topological sort for deterministic change ordering ([5074cb3](https://github.com/veesix-networks/osvbng/commit/5074cb361e23518c8a0fd89fffe20e3f8eae2b05))
* **dataplane:** bring up loopback in LCP namespace ([#102](https://github.com/veesix-networks/osvbng/issues/102)) ([a75cce3](https://github.com/veesix-networks/osvbng/commit/a75cce3b5531f1db86e4df92f8dd578e1c1ed6c5))
* **dataplane:** bring up loopback in LCP namespace and register in ifmgr ([#104](https://github.com/veesix-networks/osvbng/issues/104)) ([7736aee](https://github.com/veesix-networks/osvbng/commit/7736aee81fd91db060189e96adf528b279c19a08))
* **dhcp:** detect and reject static/dynamic IP reservation collisions ([#119](https://github.com/veesix-networks/osvbng/issues/119)) ([2b69092](https://github.com/veesix-networks/osvbng/commit/2b6909219136a457092a91da87b7994e9d849fe2))
* **dhcp:** resolve per-pool gateway and add service group pool selection ([#117](https://github.com/veesix-networks/osvbng/issues/117)) ([4dcde9e](https://github.com/veesix-networks/osvbng/commit/4dcde9eaa4f882b5e4aedab16482bdc5f2844581))
* **ipoe:** log error when address profile not found ([#125](https://github.com/veesix-networks/osvbng/issues/125)) ([78c6f5e](https://github.com/veesix-networks/osvbng/commit/78c6f5e1fb69e81e25356f17fbba60ce6fc3d8d8))
* **ospf:** use accept-all-interfaces mfib flag ([#100](https://github.com/veesix-networks/osvbng/issues/100)) ([9fa851e](https://github.com/veesix-networks/osvbng/commit/9fa851e12039289c872fa27146c0ee979efc74f6))
* **southbound:** hardcode LCP dataplane namespace and handle existing pairs on restart ([#127](https://github.com/veesix-networks/osvbng/issues/127)) ([fb884a6](https://github.com/veesix-networks/osvbng/commit/fb884a6c422f0b1323632041829db307b67c4a18))
* **southbound:** set af-packet MAC at creation instead of post-hoc sync ([#123](https://github.com/veesix-networks/osvbng/issues/123)) ([43b0b69](https://github.com/veesix-networks/osvbng/commit/43b0b69003bd19dc4b9c419fe82dff88dd2fc1f6))

## [0.1.2](https://github.com/veesix-networks/osvbng/compare/v0.1.1...v0.1.2) (2026-02-13)


### Bug Fixes

* **dataplane:** default QinQ outer TPID to 802.1ad ([008e63c](https://github.com/veesix-networks/osvbng/commit/008e63c7fa57bb9128f4c26dd0c70048ad77559b))
* **dataplane:** default QinQ outer TPID to 802.1ad ([#86](https://github.com/veesix-networks/osvbng/issues/86)) ([008e63c](https://github.com/veesix-networks/osvbng/commit/008e63c7fa57bb9128f4c26dd0c70048ad77559b))
* **dataplane:** default QinQ outer TPID to 802.1ad with per-group override ([691090a](https://github.com/veesix-networks/osvbng/commit/691090ab60b888ff87f633be019d02692d506658))
* **routing:** use loaded config for FRR config generation ([680f559](https://github.com/veesix-networks/osvbng/commit/680f559bb16794569c0751bc9e773385a0ce22f9))
* **routing:** use loaded config for FRR config generation ([5fe6dac](https://github.com/veesix-networks/osvbng/commit/5fe6daca8a38cf6016a711246cf99bfedfa654c5))
* **routing:** use loaded config for FRR config generation ([#88](https://github.com/veesix-networks/osvbng/issues/88)) ([680f559](https://github.com/veesix-networks/osvbng/commit/680f559bb16794569c0751bc9e773385a0ce22f9))

## [0.1.1](https://github.com/veesix-networks/osvbng/compare/v0.1.0...v0.1.1) (2026-02-10)


### Bug Fixes

* **build:** copy template subdirectories in qemu image build ([f003ae7](https://github.com/veesix-networks/osvbng/commit/f003ae7c003543f56663df2d8c22129d8ea795a0))
* **build:** copy template subdirectories in qemu image build ([#81](https://github.com/veesix-networks/osvbng/issues/81)) ([e4e9410](https://github.com/veesix-networks/osvbng/commit/e4e9410fc5960085453fba74d396a52a4f9c3020))
* **ipoe:** reset stale AAA-approved sessions ([#82](https://github.com/veesix-networks/osvbng/issues/82)) ([3d26c2e](https://github.com/veesix-networks/osvbng/commit/3d26c2e720fec2c939b88fadc2f4c539b747ca16))

## [0.1.0](https://github.com/veesix-networks/osvbng/compare/v0.0.4...v0.1.0) (2026-02-10)


### Features

* **ipoe:** punt IPv6 RS to control plane for per-subscriber RA handling ([#73](https://github.com/veesix-networks/osvbng/issues/73)) ([8fe8956](https://github.com/veesix-networks/osvbng/commit/8fe89567952c847b1ca789c837b5630c844ee2fe))
* **models:** add username to subscriber session model ([#76](https://github.com/veesix-networks/osvbng/issues/76)) ([718c3b0](https://github.com/veesix-networks/osvbng/commit/718c3b02ae05b3bbdf48a204dcef451e3b8b4eb9))
* **monitoring:** add subscriber session prometheus metrics and grafana dashboard ([#78](https://github.com/veesix-networks/osvbng/issues/78)) ([cb5f1b6](https://github.com/veesix-networks/osvbng/commit/cb5f1b6e7d32b4227ee191a9e6bd87a281a9cae6))
* **subscriber:** subscriber clear session functionality ([#77](https://github.com/veesix-networks/osvbng/issues/77)) ([854beff](https://github.com/veesix-networks/osvbng/commit/854beff21b4ebff4068686208ed489652304cda8))


### Bug Fixes

* **pppoe:** resolve PPPoE session egress and unicast packet handling ([#75](https://github.com/veesix-networks/osvbng/issues/75)) ([4533c25](https://github.com/veesix-networks/osvbng/commit/4533c25b86a394718eb7e1ca50f6a3f53e479917))
* **subscriber:** count dual-stack sessions by address presence and fix in-memory cache scan ([#79](https://github.com/veesix-networks/osvbng/issues/79)) ([9a806f0](https://github.com/veesix-networks/osvbng/commit/9a806f019b0141785472050a48b01c6a58330951))
