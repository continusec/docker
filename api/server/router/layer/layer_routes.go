package layer

import (
	"net/http"

	"golang.org/x/net/context"

	"github.com/Sirupsen/logrus"
	"github.com/docker/docker/api/server/httputils"
)

func (s *layerRouter) getLayersByName(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	li, err := s.backend.InspectLayer(vars["name"])
	if err != nil {
		// Swallow errors from client for now - TODO: differentiate between not found and other
		logrus.Errorf("Error inspecting layer: %s %s", vars["name"], err)
		return httputils.WriteJSON(w, http.StatusNotFound, nil)
	}
	return httputils.WriteJSON(w, http.StatusOK, li)
}
