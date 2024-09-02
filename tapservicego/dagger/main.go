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
	"dagger/tapservicego/internal/dagger"
	"log"
	"strings"
)

type Tapservicego struct {
	HelmValuesSource *string
	ChartUrl         string
	DryRun           bool
}

// Run the tap service locally
func (m *Tapservicego) Run(
	ctx context.Context,
	source *dagger.Directory,
	db *dagger.Service,
	username string,
	password string,
	dbname *string,
	port *int,
) *dagger.Container {
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
