package client

import (
	"testing"

	"github.com/devhossamali/ari-proxy/internal/integration"
)

func TestEndpointList(t *testing.T) {
	integration.TestEndpointList(t, &srv{})
}

func TestEndpointListByTech(t *testing.T) {
	integration.TestEndpointListByTech(t, &srv{})
}

func TestEndpointData(t *testing.T) {
	integration.TestEndpointData(t, &srv{})
}
