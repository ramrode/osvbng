// Copyright 2026 The osvbng Authors
// Licensed under the GNU General Public License v3.0 or later.
// SPDX-License-Identifier: GPL-3.0-or-later

package l2gw

import (
	"context"

	l2gwcomp "github.com/veesix-networks/osvbng/internal/l2gw"
	"github.com/veesix-networks/osvbng/pkg/deps"
	"github.com/veesix-networks/osvbng/pkg/handlers/show"
	"github.com/veesix-networks/osvbng/pkg/handlers/show/paths"
	"github.com/veesix-networks/osvbng/pkg/telemetry"
)

func init() {
	show.RegisterFactory(NewCircuitsHandler)
	telemetry.RegisterMetric[l2gwcomp.CircuitSummary](paths.L2GWCircuits)
}

type CircuitsHandler struct {
	deps *deps.ShowDeps
}

func NewCircuitsHandler(d *deps.ShowDeps) show.ShowHandler {
	return &CircuitsHandler{deps: d}
}

func (h *CircuitsHandler) Collect(_ context.Context, _ *show.Request) (interface{}, error) {
	if h.deps.L2GW == nil {
		return []l2gwcomp.CircuitSummary{}, nil
	}
	return h.deps.L2GW.SnapshotCircuits(), nil
}

func (h *CircuitsHandler) PathPattern() paths.Path {
	return paths.L2GWCircuits
}

func (h *CircuitsHandler) Dependencies() []paths.Path {
	return nil
}

func (h *CircuitsHandler) Summary() string {
	return "Show layer 2 wholesale gateway circuits"
}

func (h *CircuitsHandler) Description() string {
	return "List l2gw circuits: access tuple, handoff group and egress VLANs, static vs dynamic origin, state, and dataplane circuit id. Per-circuit counters are exported at /osvbng/l2gw in the stats segment."
}
