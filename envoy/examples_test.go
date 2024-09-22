package envoy_test

import (
	"context"
	"fmt"
	"log"

    "github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/envoy"
)

func ExampleRun() {
	// runEnvoyContainer {
	ctx := context.Background()

    envoyContainer, err := envoy.Run(ctx, "envoyproxy/envoy:v1.31-latest")
	defer func() {
		if err := testcontainers.TerminateContainer(envoyContainer); err != nil {
			log.Printf("failed to terminate container: %s", err)
		}
	}()
	if err != nil {
		log.Printf("failed to start container: %s", err)
		return
	}
	// }

	state, err := envoyContainer.State(ctx)
	if err != nil {
		log.Printf("failed to get container state: %s", err)
		return
	}

	fmt.Println(state.Running)

	// Output:
	// true
}
