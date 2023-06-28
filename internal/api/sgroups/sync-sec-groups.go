package sgroups

import (
	"context"

	model "github.com/H-BF/sgroups/internal/models/sgroups"
	registry "github.com/H-BF/sgroups/internal/registry/sgroups"

	sg "github.com/H-BF/protos/pkg/api/sgroups"
)

type syncGroups struct {
	wr     registry.Writer
	groups []*sg.SecGroup
	ops    sg.SyncReq_SyncOp
}

func (snc syncGroups) process(ctx context.Context) error {
	names := make([]string, 0, len(snc.groups))
	groups := make([]model.SecurityGroup, 0, len(snc.groups))
	var sc registry.Scope = registry.NoScope
	for _, g := range snc.groups {
		if snc.ops != sg.SyncReq_Delete {
			var x securityGroup
			x.from(g)
			groups = append(groups, x.SecurityGroup)
		}
		if snc.ops != sg.SyncReq_FullSync {
			names = append(names, g.GetName())
		}
	}
	if len(names) != 0 {
		sc = registry.SG(names[0], names[1:]...)
	}
	var opts []registry.Option
	if err := syncOptionsFromProto(snc.ops, &opts); err != nil {
		return err
	}
	return snc.wr.SyncSecurityGroups(ctx, groups, sc, opts...)
}

type securityGroup struct {
	model.SecurityGroup
}

func (n *securityGroup) from(g *sg.SecGroup) {
	n.Name = g.GetName()
	n.Networks = g.GetNetworks()
}
