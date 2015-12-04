package lib

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/cliconfig"
)

// RegistryLogin authenticates the docker server with a given docker registry.
// It returns UnauthorizerError when the authentication fails.
func (cli *Client) RegistryLogin(auth cliconfig.AuthConfig) (types.AuthResponse, error) {
	resp, err := cli.POST("/auth", url.Values{}, auth, nil)

	if resp != nil && resp.statusCode == http.StatusUnauthorized {
		return types.AuthResponse{}, unauthorizedError{err}
	}
	if err != nil {
		return types.AuthResponse{}, err
	}
	defer ensureReaderClosed(resp)

	var response types.AuthResponse
	err = json.NewDecoder(resp.body).Decode(&response)
	return response, err
}
