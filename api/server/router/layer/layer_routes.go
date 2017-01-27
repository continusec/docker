package layer

import (
	"net/http"

	"golang.org/x/net/context"

	"github.com/docker/docker/api/server/httputils"
)

func (s *layerRouter) getLayersByName(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	li, err := s.backend.InspectLayer(vars["name"])
	if err != nil {
		return err
	}
	return httputils.WriteJSON(w, http.StatusOK, li)
}
