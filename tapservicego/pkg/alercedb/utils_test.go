package alercedb

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

type AlerceSuite struct {
	suite.Suite

	DB                *sql.DB
	postgresContainer *postgres.PostgresContainer
	ctx               context.Context
}

func TestAlerceSuite(t *testing.T) {
	suite.Run(t, new(AlerceSuite))
}

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

func (suite *AlerceSuite) initializeLocalDB() {
	suite.T().Log("Using local database")
	var err error
	suite.ctx = context.Background()
	suite.postgresContainer, err = CreatePostgresContainer(suite.ctx)
	if err != nil {
		suite.T().Fatal(err)
	}
	var connStr string
	connStr, err = suite.postgresContainer.ConnectionString(suite.ctx)
	if err != nil {
		suite.T().Fatal(err)
	}
	os.Setenv("DATABASE_URL", connStr)
	suite.DB, err = GetDB(connStr)
	if err != nil {
		log.Fatal(err)
	}
}

func (suite *AlerceSuite) initializeDaggerDB() {
	suite.T().Log("Using Dagger database")
	connUrl := "host=db user=testuser password=testpassword port=5432"
	os.Setenv("DATABASE_URL", connUrl)
	tmpDb, err := GetDB(connUrl)
	if err != nil {
		log.Fatal(err)
	}
	_, err = tmpDb.Exec("CREATE DATABASE alercedb")
	if err != nil {
		log.Fatal(err)
	}
	tmpDb.Close()
	connUrl = connUrl + " dbname=alercedb"
	os.Setenv("DATABASE_URL", connUrl)
	suite.DB, err = GetDB(connUrl)
	if err != nil {
		log.Fatal(err)
	}
}

func (suite *AlerceSuite) SetupSuite() {
	if os.Getenv("ENV") == "DEV" || os.Getenv("ENV") == "" {
		suite.initializeLocalDB()
	} else if os.Getenv("ENV") == "CI" {
		suite.initializeDaggerDB()
	} else {
		log.Fatal("Unknown environment")
	}
}

func (suite *AlerceSuite) TearDownSuite() {
	if os.Getenv("ENV") == "DEV" || os.Getenv("ENV") == "" {
		CleanUpContainer(suite.ctx, suite.postgresContainer)
	}
	suite.DB.Close()
}

func RestoreDatabase(db *sql.DB, suite *AlerceSuite) {
	err := DropTables(db)
	if err != nil {
		suite.T().Fatal(err)
	}
}
