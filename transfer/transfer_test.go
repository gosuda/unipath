package transfer

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHTTP(t *testing.T) {
	err := Transfer(context.Background(), "https://raw.githubusercontent.com/gosuda/unipath/main/go.mod", "test/go.mod")
	require.NoError(t, err)
}
