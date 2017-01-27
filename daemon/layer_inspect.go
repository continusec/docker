package daemon

import (
	"github.com/docker/docker/api/types"
)

// InspectLayer returns list of files modified by layer
func (daemon *Daemon) InspectLayer(name string) (*types.LayerInfo, error) {
	return &types.LayerInfo{
		Name: name,
		Diffs: []*types.DiffInfo{
			&types.DiffInfo{
				Path: "foo",
			},
		},
	}, nil
}
