package client

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/docker/docker/api/types"
	"golang.org/x/net/context"
)

type layerNotFoundError struct {}

func (e *layerNotFoundError) NotFound() bool {
	return true
}

func (e *layerNotFoundError) Error() string {
	return "layer not found"
}

var (
	// ErrLayerNotFound returned when layer not found
	ErrLayerNotFound = &layerNotFoundError{}
)

// InspectLayer returns the layer information and its raw representation.
func (cli *Client) InspectLayer(ctx context.Context, name string) (*types.LayerInfo, []byte, error) {
	serverResp, err := cli.get(ctx, "/layers/"+name+"/json", nil, nil)
	if err != nil {
		if serverResp.statusCode == http.StatusNotFound {
			return nil, nil, ErrLayerNotFound
		}
		return nil, nil, err
	}
	defer ensureReaderClosed(serverResp)

	body, err := ioutil.ReadAll(serverResp.body)
	if err != nil {
		return nil, nil, err
	}

	var response types.LayerInfo
	rdr := bytes.NewReader(body)
	err = json.NewDecoder(rdr).Decode(&response)
	if err != nil {
		return nil, nil, err
	}
	return &response, body, nil
}
