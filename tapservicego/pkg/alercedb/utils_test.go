package alercedb

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

var ctx context.Context
var container *postgres.PostgresContainer

func CreatePostgresContainer(ctx context.Context) (*postgres.PostgresContainer, error) {
	postgresContainer, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("docker.io/postgres:16-alpine"),
		postgres.WithDatabase("alercedb"),
		postgres.WithUsername("testuser"),
		postgres.WithPassword("testpassword"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		log.Printf("failed to start container: %s", err)
		return nil, err
	}
	err = postgresContainer.Snapshot(ctx, postgres.WithSnapshotName("initial"))
	if err != nil {
		log.Printf("failed to snapshot container: %s", err)
		return nil, err
	}
	return postgresContainer, nil
}

func CleanUpContainer(ctx context.Context, postgresContainer *postgres.PostgresContainer) {
	if err := postgresContainer.Terminate(ctx); err != nil {
		log.Fatalf("failed to terminate container: %s", err)
	}
}

// GetDB creates a new database connection
// using the provided URL
// and returns a pointer to the connection
// or an error if one occurs.
// The URL should be in the format:
// postgresql://user:password@host:port/database
func GetDB(url string) (*sql.DB, error) {
	db, err := sql.Open("pgx", url)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func initializeLocalDB() {
	log.Println("Using local database")
	var err error
	ctx = context.Background()
	container, err = CreatePostgresContainer(ctx)
	if err != nil {
		log.Fatal(err)
	}
	var connStr string
	connStr, err = container.ConnectionString(ctx)
	if err != nil {
		log.Fatal(err)
	}
	os.Setenv("DATABASE_URL", connStr)
}

func initializeDaggerDB() {
	log.Println("Using Dagger database")
	os.Setenv("DATABASE_URL", "host=db user=testuser password=testpassword port=5432")
}

func globalSetup() {
	if os.Getenv("ENV") == "DEV" || os.Getenv("ENV") == "" {
		initializeLocalDB()
	} else if os.Getenv("ENV") == "CI" {
		initializeDaggerDB()
	} else {
		log.Fatal("Unknown environment")
	}
	setUpTestDatabase()
}

func setUpTestDatabase() {
	if os.Getenv("ENV") == "DEV" || os.Getenv("ENV") == "" {
		return
	}
	db, err := GetDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	_, err = db.Exec("CREATE DATABASE alercedb")
	if err != nil {
		log.Fatal(err)
	}
	os.Setenv("DATABASE_URL", os.Getenv("DATABASE_URL")+" dbname=alercedb")
}

func globalTeardown() {
	if os.Getenv("ENV") == "DEV" || os.Getenv("ENV") == "" {
		CleanUpContainer(ctx, container)
	}
}

func TestMain(m *testing.M) {
	globalSetup()
	code := m.Run()
	globalTeardown()
	os.Exit(code)
}

func restoreDatabase() {
	db, err := GetDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = DropTables(db)
	if err != nil {
		log.Fatal(err)
	}
}
