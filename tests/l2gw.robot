# Copyright 2026 The osvbng Authors
# Licensed under the GNU General Public License v3.0 or later.
# SPDX-License-Identifier: GPL-3.0-or-later

*** Settings ***
Library             OperatingSystem
Library             String
Library             Process
Resource            common.robot
Resource            bngblaster.robot

*** Keywords ***
Wait For Blaster Sessions Established
    [Documentation]    L2GW subscribers are never terminated on the BNG, so the
    ...    subscriber-sessions API stays empty; the BNG Blaster established
    ...    counter is the end-to-end proof (DHCP answered by the a10nsp side
    ...    through the cross-connect).
    [Arguments]    ${bngblaster}    ${expected_count}    ${timeout}=120s    ${interval}=2s
    Wait Until Keyword Succeeds    ${timeout}    ${interval}
    ...    Blaster Sessions Established    ${bngblaster}    ${expected_count}

Blaster Sessions Established
    [Arguments]    ${bngblaster}    ${expected_count}
    ${rc}    ${cli_output} =    BNG Blaster CLI Command    ${bngblaster}    session-counters
    Should Be Equal As Integers    ${rc}    0
    ${rc}    ${established} =    Run And Return Rc And Output
    ...    echo '${cli_output}' | python3 -c "import sys,json; d=json.load(sys.stdin); print(d.get('session-counters',{}).get('sessions-established',0))"
    Should Be Equal As Strings    ${established}    ${expected_count}    BNG Blaster established ${established}/${expected_count} sessions

Get L2GW Circuits
    [Arguments]    ${container}
    ${output} =    Get osvbng API Response    ${container}    /api/show/l2gw/circuits
    RETURN    ${output}

Verify L2GW Circuit Count
    [Arguments]    ${container}    ${expected}    ${state}=installed
    ${output} =    Get L2GW Circuits    ${container}
    ${script} =    Catenate    SEPARATOR=${SPACE}
    ...    import json,os;
    ...    d=json.loads(os.environ['JSON']);
    ...    c=[x for x in (d.get('data') or []) if x.get('state')==os.environ['STATE']];
    ...    print(len(c))
    ${result} =    Run Process    python3    -c    ${script}
    ...    env:JSON=${output}    env:STATE=${state}    stderr=STDOUT
    Should Be Equal As Integers    ${result.rc}    0
    Should Be Equal As Strings    ${result.stdout}    ${expected}    Expected ${expected} ${state} l2gw circuits, got ${result.stdout}

Wait For L2GW Circuit Count
    [Arguments]    ${container}    ${expected}    ${state}=installed    ${timeout}=60s    ${interval}=2s
    Wait Until Keyword Succeeds    ${timeout}    ${interval}
    ...    Verify L2GW Circuit Count    ${container}    ${expected}    ${state}

Verify L2GW Circuit Field
    [Documentation]    Assert every circuit in the API snapshot satisfies a
    ...    python expression over the circuit dict bound as `c`.
    [Arguments]    ${container}    ${expression}    ${message}
    ${output} =    Get L2GW Circuits    ${container}
    ${script} =    Catenate    SEPARATOR=${SPACE}
    ...    import json,os,sys;
    ...    d=json.loads(os.environ['JSON']);
    ...    cs=d.get('data') or [];
    ...    bad=[c for c in cs if not (${expression})];
    ...    sys.exit(0 if cs and not bad else (print('violating circuits: %s' % bad if bad else 'no circuits') or 1))
    ${result} =    Run Process    python3    -c    ${script}
    ...    env:JSON=${output}    stderr=STDOUT
    Log    ${result.stdout}
    Should Be Equal As Integers    ${result.rc}    0    ${message}

Verify VPP L2GW Circuits Counters Non-Zero
    [Documentation]    Every circuit entry in the dataplane must show
    ...    non-zero packets in both directions once traffic has flowed.
    [Arguments]    ${container}
    ${output} =    Execute VPP Command    ${container}    show osvbng l2gw circuits
    ${script} =    Catenate    SEPARATOR=${SPACE}
    ...    import os,re,sys;
    ...    out=os.environ['OUT'];
    ...    pkts=[int(m) for m in re.findall(r'(\\d+) pkts', out)];
    ...    sys.exit(0 if pkts and all(p>0 for p in pkts) else (print(out) or 1))
    ${result} =    Run Process    python3    -c    ${script}
    ...    env:OUT=${output}    stderr=STDOUT
    Log    ${result.stdout}
    Should Be Equal As Integers    ${result.rc}    0    l2gw dataplane counters missing or zero

Snapshot L2GW Circuit IDs
    [Documentation]    Return the sorted (session id, access tuple, handoff
    ...    tuple) set for restart-survival comparison. Session ids are
    ...    included deliberately: a restored circuit keeps its session id,
    ...    so any re-authentication shows up as a changed id even when the
    ...    VLAN tuples survive. Dataplane circuit ids may be reassigned on
    ...    re-install; they are excluded.
    [Arguments]    ${container}
    ${output} =    Get L2GW Circuits    ${container}
    ${script} =    Catenate    SEPARATOR=${SPACE}
    ...    import json,os;
    ...    d=json.loads(os.environ['JSON']);
    ...    print(sorted((c.get('session_id'),c.get('mac'),c.get('access_interface'),c.get('access_svlan'),c.get('access_cvlan',0),c.get('handoff_group'),c.get('handoff_svlan',0),c.get('handoff_cvlan',0)) for c in (d.get('data') or [])))
    ${result} =    Run Process    python3    -c    ${script}
    ...    env:JSON=${output}    stderr=STDOUT
    Should Be Equal As Integers    ${result.rc}    0
    RETURN    ${result.stdout}
