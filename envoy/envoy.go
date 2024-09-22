package envoy

import (
	"context"
	"fmt"

	"github.com/testcontainers/testcontainers-go"
)

// EnvoyContainer represents the Envoy container type used in the module
type EnvoyContainer struct {
	testcontainers.Container
}

// Run creates an instance of the Envoy container type
func Run(ctx context.Context, img string, opts ...testcontainers.ContainerCustomizer) (*EnvoyContainer, error) {
	req := testcontainers.ContainerRequest{
		Image: img,
	}

	genericContainerReq := testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	}

	for _, opt := range opts {
		if err := opt.Customize(&genericContainerReq); err != nil {
			return nil, fmt.Errorf("customize: %w", err)
		}
	}

	container, err := testcontainers.GenericContainer(ctx, genericContainerReq)
	var c *EnvoyContainer
	if container != nil {
		c = &EnvoyContainer{Container: container}
	}

	if err != nil {
		return c, fmt.Errorf("generic container: %w", err)
	}

	return c, nil
}
