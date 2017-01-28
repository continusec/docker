package daemon

import (
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/layer"
	digest "github.com/opencontainers/go-digest"
)

// InspectLayer returns list of files modified by layer
func (daemon *Daemon) InspectLayer(name string) (*types.LayerInfo, error) {
	id, err := digest.Parse(name)
	if err != nil {
		// Try with commonly ommitted prefix
		id, err = digest.Parse("sha256:" + name)
	}
	if err != nil {
		return nil, err
	}

	lref, err := daemon.layerStore.Get(layer.ChainID(id))
	if err != nil {
		return nil, err
	}
	
	md, err := lref.Metadata()
	if err != nil {
		return nil, err
	}

	fmt.Printf("%#v\n", md)

	return &types.LayerInfo{
		Name: id.String(),
		Diffs: []*types.DiffInfo{
			&types.DiffInfo{
				Path: "foo",
			},
		},
	}, nil
}
