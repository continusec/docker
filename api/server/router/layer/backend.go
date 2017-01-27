package layer

import "github.com/docker/docker/api/types"

// Backend is all the methods that need to be implemented
// to provide layer specific functionality.
type Backend interface {
	layerBackend
}

type layerBackend interface {
	InspectLayer(name string) (*types.LayerInfo, error)
}
