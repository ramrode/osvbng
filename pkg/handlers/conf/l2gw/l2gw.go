// Copyright 2026 The osvbng Authors
// Licensed under the GNU General Public License v3.0 or later.
// SPDX-License-Identifier: GPL-3.0-or-later

package l2gw

import (
	"context"
	"fmt"

	l2gwcfg "github.com/veesix-networks/osvbng/pkg/config/l2gw"
	"github.com/veesix-networks/osvbng/pkg/deps"
	"github.com/veesix-networks/osvbng/pkg/handlers/conf"
	"github.com/veesix-networks/osvbng/pkg/handlers/conf/paths"
)

func init() {
	conf.RegisterFactory(NewL2GWHandler)
}

// L2GWHandler reconciles the whole l2gw block on commit. It holds the
// ConfDeps pointer rather than the component because the component is
// created after the bootstrap handler registration; during bootstrap
// Apply is a no-op and the component's own Start applies the config.
type L2GWHandler struct {
	deps *deps.ConfDeps
}

func NewL2GWHandler(d *deps.ConfDeps) conf.Handler {
	return &L2GWHandler{deps: d}
}

func (h *L2GWHandler) Validate(_ context.Context, hctx *conf.HandlerContext) error {
	if hctx.NewValue == nil {
		return nil
	}
	cfg, ok := hctx.NewValue.(*l2gwcfg.L2GWConfig)
	if !ok {
		return fmt.Errorf("expected *l2gw.L2GWConfig, got %T", hctx.NewValue)
	}
	if err := cfg.Validate(); err != nil {
		return err
	}
	if hctx.Config != nil {
		for name, hg := range cfg.HandoffGroups {
			if _, ok := hctx.Config.Interfaces[hg.Interface]; !ok {
				return fmt.Errorf("l2gw handoff-group %q: interface %q is not defined in interfaces", name, hg.Interface)
			}
		}
		for i, sm := range cfg.StaticMaps {
			if sm == nil {
				continue
			}
			if _, ok := hctx.Config.Interfaces[sm.AccessInterface]; !ok {
				return fmt.Errorf("l2gw static-map %d: access-interface %q is not defined in interfaces", i, sm.AccessInterface)
			}
		}
	}
	return nil
}

func (h *L2GWHandler) Apply(_ context.Context, hctx *conf.HandlerContext) error {
	if h.deps.L2GW == nil {
		return nil
	}
	return h.deps.L2GW.ReconcileConfig(hctx.Config)
}

func (h *L2GWHandler) Rollback(_ context.Context, hctx *conf.HandlerContext) error {
	if h.deps.L2GW == nil {
		return nil
	}
	return h.deps.L2GW.ReconcileConfig(hctx.Config)
}

func (h *L2GWHandler) PathPattern() paths.Path {
	return paths.L2GW
}

func (h *L2GWHandler) Dependencies() []paths.Path {
	return nil
}

func (h *L2GWHandler) Callbacks() *conf.Callbacks {
	return nil
}

func (h *L2GWHandler) Summary() string {
	return "Layer 2 wholesale gateway configuration"
}

func (h *L2GWHandler) Description() string {
	return "Handoff groups and static maps for the layer 2 wholesale gateway. Commits reconcile static circuits (add, remove, re-point) and rebuild the egress VLAN allocators; installed dynamic circuits are unaffected."
}
