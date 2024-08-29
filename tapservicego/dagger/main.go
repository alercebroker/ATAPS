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
	"fmt"
	"log"
	"math"
	"math/rand/v2"
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
		WithExec([]string{"go", "build", "-o", "/usr/local/bin/ataps", "cmd/tapservice/main.go"})
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
func (m *Tapservicego) Build(ctx context.Context, source *Directory, port int) *Container {
	return dag.Container().
		From("golang:1.22.3-bookworm").
		WithExec([]string{"apt-get", "update"}).
		WithExec([]string{"apt-get", "install", "libcfitsio-dev", "--yes"}).
		WithFile("/bin/ataps", m.BuildEnv(ctx, source).File("/usr/local/bin/ataps")).
		WithExposedPort(port).
		WithEntrypoint([]string{"ataps"})
}

// Publish the tap service container image
func (m *Tapservicego) PublishContainer(
	ctx context.Context,
	source *Directory,
	username *string,
	password *Secret,
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

// Publish the tap service Helm chart to GHCR
func (m *Tapservicego) PublishHelmChart(
	ctx context.Context,
	chartDir *Directory,
	username string,
	password *Secret,
	ghOrg *string,
) (string, error) {
	container := dag.Container().
		From("alpine/k8s:1.31.0").
		WithDirectory("/usr/src/chart", chartDir).
		WithExec([]string{"helm", "package", "/usr/src/chart", "-d", "/usr/src"})
	version, err := container.
		WithExec([]string{"sh", "-c", "grep '^version:' /usr/src/chart/Chart.yaml | awk '{print $2}' | tr -d '\"'"}).
		Stdout(ctx)
	version = strings.Trim(version, "\n")
	if err != nil {
		log.Fatalf("Error getting chart version: %v", err)
	}
	pwd, err := password.Plaintext(ctx)
	if err != nil {
		log.Fatalf("Error getting password: %v", err)
	}
	if ghOrg == nil {
		ghOrg = &username
	}
	registry := fmt.Sprintf("oci://ghcr.io/%s/tapservice-chart", *ghOrg)
	return container.
		WithExec([]string{"helm", "registry", "login", "-u", username, "-p", pwd, "ghcr.io"}).
		WithExec([]string{"helm", "push", fmt.Sprintf("/usr/src/tapservice-%s.tgz", version), registry}).
		Stdout(ctx)
}

// Run the tap service
func (m *Tapservicego) Run(
	ctx context.Context,
	source *Directory,
	db *Service,
	username string,
	password string,
	dbname *string,
	port *int,
) *Container {
	portOverride := 8080
	if port == nil {
		port = &portOverride
	}
	container := m.Build(ctx, source, *port)
	if dbname == nil {
		dbname = &username
	}
	if db == nil {
		log.Println("No database URL provided, launching a new postgres container.")
		dbCtr := dag.
			Container().
			From("index.docker.io/postgres").
			WithEnvVariable("POSTGRES_USER", username).
			WithEnvVariable("POSTGRES_PASSWORD", password).
			AsService()
		container = container.WithServiceBinding("db", dbCtr).
			WithEnvVariable("DATABASE_URL", "postgres://"+username+":"+password+"@db:5432/"+*dbname)
	} else {
		endpoint, err := db.Endpoint(ctx)
		if err != nil {
			log.Fatalf("Error getting database endpoint: %v", err)
		}
		host := strings.Split(endpoint, ":")[0]
		port := strings.Split(endpoint, ":")[1]
		container = container.
			WithServiceBinding("db", db).
			WithEnvVariable("DATABASE_URL", "postgres://"+username+":"+password+"@"+host+":"+port+"/"+*dbname)
	}
	return container.WithExec([]string{"ataps"})
}
