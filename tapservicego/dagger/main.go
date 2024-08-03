// A generated module for Tapservicego functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
	"log"
	"strings"
)

type Tapservicego struct{}

// Builds the go package
func (m *Tapservicego) BuildEnv(ctx context.Context, source *Directory) *Container {
	return dag.Container().
		From("golang:1.22.3-bookworm").
		WithExec([]string{"apt-get", "update"}).
		WithExec([]string{"apt-get", "install", "libcfitsio-dev", "--yes"}).
		WithWorkdir("/usr/src/app").
		WithFile("go.mod", source.File("go.mod")).
		WithFile("go.sum", source.File("go.sum")).
		WithExec([]string{"go", "mod", "download"}).
		WithExec([]string{"go", "mod", "verify"}).
		WithDirectory("/usr/src/app", source).
		WithExec([]string{"go", "build", "-o", "/usr/local/bin/ataps", "."})
}

// Tests the go package
func (m *Tapservicego) Test(ctx context.Context, source *Directory) (string, error) {
	db := dag.
		Container().
		From("docker.io/postgres:16-alpine").
		WithEnvVariable("POSTGRES_USER", "testuser").
		WithEnvVariable("POSTGRES_PASSWORD", "testpassword").
		AsService()
	return m.BuildEnv(ctx, source).
		WithEnvVariable("ENV", "CI").
		WithServiceBinding("db", db).
		WithExec([]string{"go", "test", "-failfast", "./..."}).
		Stdout(ctx)
}

// Build the tap service
func (m *Tapservicego) Build(ctx context.Context, source *Directory) *Container {
	return dag.Container().
		From("golang:1.22").
		WithFile("/bin/ataps", m.BuildEnv(ctx, source).File("/usr/local/bin/ataps")).
		WithExposedPort(8080).
		WithEntrypoint([]string{"ataps"})
}

// Run the tap service
func (m *Tapservicego) Run(ctx context.Context, source *Directory, db *Service) *Container {
	container := m.Build(ctx, source)
	if db == nil {
		log.Println("No database URL provided, launching a new postgres container.")
		dbCtr := dag.
			Container().
			From("index.docker.io/postgres").
			WithEnvVariable("POSTGRES_USER", "postgres").
			WithEnvVariable("POSTGRES_PASSWORD", "postgres").
			AsService()
		container = container.WithServiceBinding("db", dbCtr).
			WithEnvVariable("DATABASE_URL", "postgres://postgres:postgres@db:5432/postgres")
	} else {
		endpoint, err := db.Endpoint(ctx)
		if err != nil {
			log.Fatalf("Error getting database endpoint: %v", err)
		}
		host := strings.Split(endpoint, ":")[0]
		port := strings.Split(endpoint, ":")[1]
		container = container.
			WithServiceBinding("db", db).
			WithEnvVariable("DATABASE_URL", "postgres://postgres:postgres@"+host+":"+port+"/postgres")
	}
	return container.WithExec([]string{"ataps"})
}
