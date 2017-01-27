package layer

import "github.com/docker/docker/api/server/router"

// layerRouter is a router to talk with the layer controller
type layerRouter struct {
	backend Backend
	routes  []router.Route
}

// NewRouter initializes a new image router
func NewRouter(backend Backend) router.Router {
	r := &layerRouter{
		backend: backend,
	}
	r.initRoutes()
	return r
}

// Routes returns the available routes to the layer controller
func (r *layerRouter) Routes() []router.Route {
	return r.routes
}

// initRoutes initializes the routes in the image router
func (r *layerRouter) initRoutes() {
	r.routes = []router.Route{
		// GET
		router.NewGetRoute("/layers/{name:.*}/json", r.getLayersByName),
	}
}
