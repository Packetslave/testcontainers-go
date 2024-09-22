package envoy_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

    "github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/envoy"
)

func TestEnvoy(t *testing.T) {
	ctx := context.Background()

    ctr, err := envoy.Run(ctx, "envoyproxy/envoy:v1.31-latest")
	testcontainers.CleanupContainer(t, ctr)
	require.NoError(t, err)

	// perform assertions
}
