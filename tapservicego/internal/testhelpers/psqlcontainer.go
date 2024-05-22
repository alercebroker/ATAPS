package testhelpers

import (
	"context"
	"log"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func CreatePostgresContainer(ctx context.Context) (*postgres.PostgresContainer, error) {
	postgresContainer, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("docker.io/postgres:16-alpine"),
		// postgres.WithInitScripts(filepath.Join("testdata", "init-user-db.sh")),
		// postgres.WithConfigFile(filepath.Join("testdata", "my-postgres.conf")),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		log.Printf("failed to start container: %s", err)
		return nil, err
	}
	return postgresContainer, nil
}

func CleanUpContainer(ctx context.Context, postgresContainer *postgres.PostgresContainer) {
	if err := postgresContainer.Terminate(ctx); err != nil {
		log.Fatalf("failed to terminate container: %s", err)
	}
}
