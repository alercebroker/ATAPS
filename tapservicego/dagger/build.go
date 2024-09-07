package main

import (
	"context"
	"dagger/tapservicego/internal/dagger"
	"fmt"
	"math"
	"math/rand/v2"
	"strings"
)

// Builds the go package
func (m *Tapservicego) BuildEnv(ctx context.Context, source *dagger.Directory) *dagger.Container {
	return dag.Container().
		From("golang:1.23.0-bookworm").
		WithExec([]string{"apt-get", "update"}).
		WithExec([]string{"apt-get", "install", "libcfitsio-dev", "--yes"}).
		WithWorkdir("/usr/src/app").
		WithFile("go.mod", source.File("go.mod")).
		WithFile("go.sum", source.File("go.sum")).
		WithExec([]string{"go", "mod", "download"}).
		WithExec([]string{"go", "mod", "verify"}).
		WithDirectory("/usr/src/app", source).
		WithExec([]string{"go", "build", "-o", "/usr/local/bin/ataps", "cmd/tapservice/main.go"})
}

// Build the tap service
func (m *Tapservicego) Build(ctx context.Context, source *dagger.Directory, port int) *dagger.Container {
	return dag.Container().
		From("golang:1.23.0-bookworm").
		WithExec([]string{"apt-get", "update"}).
		WithExec([]string{"apt-get", "install", "libcfitsio-dev", "--yes"}).
		WithFile("/bin/ataps", m.BuildEnv(ctx, source).File("/usr/local/bin/ataps")).
		WithExposedPort(port).
		WithEntrypoint([]string{"ataps"})
}

// Publish the tap service container image
func (m *Tapservicego) PublishContainer(
	ctx context.Context,
	source *dagger.Directory,
	username *string,
	password *dagger.Secret,
	tags *string,
) ([]string, error) {
	container := m.Build(ctx, source, 8080)
	if username == nil && password == nil {
		address, err := container.Publish(ctx, fmt.Sprintf("ttl.sh/tapservice-%.0f:2h", math.Floor(rand.Float64()*10000000)))
		if err != nil {
			return nil, err
		}
		return []string{address}, nil
	} else {
		tagList := make([]string, 0)
		if tags == nil {
			tagList = append(tagList, "latest")
		}
		splitTags := strings.Split(*tags, ",")
		for _, tag := range splitTags {
			tagList = append(tagList, tag)
		}
		ctr := container.
			WithRegistryAuth("ghcr.io", *username, password)
		addresses := make([]string, len(tagList))
		for i, tag := range tagList {
			address, err := ctr.Publish(ctx, fmt.Sprintf("ghcr.io/%s/tapservice:%s", *username, tag))
			if err != nil {
				return nil, err
			}
			addresses[i] = address
		}
		return addresses, nil
	}
}
