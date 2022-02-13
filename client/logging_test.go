package client

import (
	"testing"

	"github.com/devhossamali/ari-proxy/v5/internal/integration"
)

func TestLoggingList(t *testing.T) {
	integration.TestLoggingList(t, &srv{})
}

func TestLoggingCreate(t *testing.T) {
	integration.TestLoggingCreate(t, &srv{})
}

func TestLoggingRotate(t *testing.T) {
	integration.TestLoggingRotate(t, &srv{})
}

func TestLoggingDelete(t *testing.T) {
	integration.TestLoggingDelete(t, &srv{})
}
