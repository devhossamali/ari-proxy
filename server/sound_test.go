package server

import (
	"testing"

	"github.com/devhossamali/ari-proxy/v5/internal/integration"
)

func TestSoundData(t *testing.T) {
	integration.TestSoundData(t, &srv{})
}

func TestSoundList(t *testing.T) {
	integration.TestSoundList(t, &srv{})
}
