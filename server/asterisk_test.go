package server

import (
	"testing"

	"github.com/devhossamali/ari-proxy/v5/internal/integration"
)

func TestAsteriskInfo(t *testing.T) {
	integration.TestAsteriskInfo(t, &srv{})
}

func TestAsteriskVariablesGet(t *testing.T) {
	integration.TestAsteriskVariablesGet(t, &srv{})
}

func TestAsteriskVariablesSet(t *testing.T) {
	integration.TestAsteriskVariablesSet(t, &srv{})
}
