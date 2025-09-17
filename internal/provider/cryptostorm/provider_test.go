package cryptostorm

import (
	"testing"

	"github.com/qdm12/gluetun/internal/constants/providers"
)

func TestProvider_Name(t *testing.T) {
	t.Parallel()

	provider := Provider{}
	if provider.Name() != providers.Cryptostorm {
		t.Errorf("Expected provider name %s, got %s", providers.Cryptostorm, provider.Name())
	}
}