package utils

import (
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/ehazlett/interlock/ext"
)

func WebsocketEndpoints(config types.Container) []string {
	websocketEndpoints := []string{}

	for l, v := range config.Labels {
		// this is for labels like interlock.websocket_endpoint.1=foo
		if strings.Index(l, ext.InterlockWebsocketEndpointLabel) > -1 {
			websocketEndpoints = append(websocketEndpoints, v)
		}
	}

	return websocketEndpoints
}
