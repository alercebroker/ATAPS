package main

import (
	"context"
	"dagger/tapservicego/internal/dagger"
	"strings"
)

// Tests the go package
func (m *Tapservicego) Test(ctx context.Context, source *dagger.Directory, extraArgs *string) (string, error) {
	db := dag.
		Container().
		From("docker.io/postgres:16-alpine").
		WithEnvVariable("POSTGRES_USER", "testuser").
		WithEnvVariable("POSTGRES_PASSWORD", "testpassword").
		AsService()
	goTestCommand := []string{"go", "test", "./..."}
	if extraArgs != nil {
		goTestArgs := strings.Split(*extraArgs, ",")
		goTestCommand = append(goTestCommand, goTestArgs...)
	}
	return m.BuildEnv(ctx, source).
		WithEnvVariable("ENV", "CI").
		WithServiceBinding("db", db).
		WithExec(goTestCommand).
		Stdout(ctx)
}
