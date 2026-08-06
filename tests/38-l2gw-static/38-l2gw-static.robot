# Copyright 2026 The osvbng Authors
# Licensed under the GNU General Public License v3.0 or later.
# SPDX-License-Identifier: GPL-3.0-or-later

*** Comments ***
L2GW static map suite. Whole access S-VLANs (10-12) are cross-connected
transparently to the ISP handoff port by configuration: no DHCP trigger,
no RADIUS, no per-subscriber state. The BNG Blaster terminates both ends:
IPoE clients on the QinQ access side and the a10nsp interface as the
retail ISP's lightweight BNG on the handoff side. Sessions establishing
proves DHCP is answered by the ISP side through the dataplane
cross-connect with the client MAC preserved (a10nsp matches sessions by
source MAC).

*** Settings ***
Library             OperatingSystem
Library             String
Library             Process
Resource            ../common.robot
Resource            ../bngblaster.robot
Resource            ../sessions.robot
Resource            ../l2gw.robot

Suite Setup         Deploy Topology    ${lab-file}
Suite Teardown      Teardown L2GW Static Test

*** Variables ***
${lab-name}         osvbng-l2gw-static
${lab-file}         ${CURDIR}/38-l2gw-static.clab.yml
${bng1}             clab-${lab-name}-bng1
${subscribers}      clab-${lab-name}-subscribers
${session-count}    3
${circuit-count}    3

*** Test Cases ***
Verify BNG Is Healthy
    [Documentation]    Wait for osvbng to fully start.
    Wait For osvbng Healthy    bng1    ${lab-name}

Verify L2GW Plugin Loaded
    [Documentation]    The osvbng_l2gw VPP plugin must be loaded and its CLI reachable.
    ${output} =    Execute VPP Command    ${bng1}    show osvbng l2gw circuits
    Should Not Contain    ${output}    unknown input

Verify Static Circuits Installed
    [Documentation]    One wildcard circuit per S-VLAN in the static map range,
    ...    installed at component start with no signalling.
    Wait For L2GW Circuit Count    ${bng1}    ${circuit-count}
    Verify L2GW Circuit Field    ${bng1}    c.get('static') and c.get('access_cvlan_any') and c.get('transparent')
    ...    static circuits must be wildcard (any C-VLAN) and transparent
    Verify L2GW Circuit Field    ${bng1}    c.get('handoff_group')=='isp-green'
    ...    all circuits must resolve to handoff group isp-green

Establish Subscriber Sessions Through Cross-Connect
    [Documentation]    IPoE DHCP must be answered by the a10nsp (ISP) side, not
    ...    the BNG: osvbng never terminates DHCP on l2gw circuits.
    Start BNG Blaster In Background    ${subscribers}
    Wait For Blaster Sessions Established    ${subscribers}    ${session-count}

Verify No Local Termination
    [Documentation]    The BNG subscriber table must stay empty; leases came
    ...    from the ISP side of the wholesale handoff.
    ${output} =    Get osvbng API Response    ${bng1}    /api/show/subscriber/sessions
    ${result} =    Run Process    python3    -c
    ...    import json,os; print(len(json.loads(os.environ['JSON']).get('data') or []))
    ...    env:JSON=${output}    stderr=STDOUT
    Should Be Equal As Strings    ${result.stdout}    0    l2gw subscribers were terminated locally

Verify Session Traffic Flows
    [Documentation]    Bidirectional session traffic between access and a10nsp
    ...    sides through the dataplane cross-connect.
    Wait Until Keyword Succeeds    60s    5s
    ...    Verify Traffic Flowing    ${subscribers}    ${session-count}

Verify Dataplane Circuit Counters
    [Documentation]    Per-circuit upstream/downstream counters in the l2gw
    ...    plugin must be counting the session traffic.
    Verify VPP L2GW Circuits Counters Non-Zero    ${bng1}

Verify L2GW Input Node Active
    [Documentation]    Frames on armed ports traverse the l2gw-input feature node.
    Verify VPP Node Calls Non-Zero    ${bng1}    l2gw-input

Verify BNG Blaster Report
    [Documentation]    Stop BNG Blaster and verify the final report.
    Stop BNG Blaster    ${subscribers}
    ${established} =    Get BNG Blaster Report Field    ${subscribers}    sessions-established
    Should Be Equal As Strings    ${established}    ${session-count}

*** Keywords ***
Teardown L2GW Static Test
    Run Keyword And Ignore Error    Stop BNG Blaster    ${subscribers}
    Destroy Topology    ${lab-file}
