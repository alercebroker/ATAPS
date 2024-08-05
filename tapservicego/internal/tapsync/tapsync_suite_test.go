package tapsync

import (
	"ataps/internal/testhelpers"
	"context"
	"database/sql"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

type TapSyncTestSuite struct {
	suite.Suite
	DB            *sql.DB
	connUrl       string
	PsqlContainer *postgres.PostgresContainer
	ctx           context.Context
	Service       *TapSyncService
}

func (suite *TapSyncTestSuite) SetupSuite() {
	if os.Getenv("ENV") == "DEV" || os.Getenv("ENV") == "" {
		suite.InitializeLocalDB()
	} else if os.Getenv("ENV") == "CI" {
		suite.InitializeDaggerDB()
	} else {
		suite.T().Fatal("Unknown environment")
	}
	suite.Service = NewTapSyncService()
}

func (suite *TapSyncTestSuite) TearDownSuite() {
	if os.Getenv("ENV") == "DEV" || os.Getenv("ENV") == "" {
		testhelpers.CleanUpContainer(suite.ctx, suite.PsqlContainer)
	}
	suite.DB.Close()
}

func (suite *TapSyncTestSuite) SetupTest() {
}

func (suite *TapSyncTestSuite) TearDownTest() {
}

func TestTapSyncTestSuite(t *testing.T) {
	suite.Run(t, new(TapSyncTestSuite))
}

func (suite *TapSyncTestSuite) InitializeLocalDB() {
	suite.T().Log("Using local database")
	var err error
	suite.ctx = context.Background()
	suite.PsqlContainer, err = testhelpers.CreatePostgresContainer(suite.ctx, "tapsync")
	if err != nil {
		suite.T().Fatal(err)
	}
	var connStr string
	connStr, err = suite.PsqlContainer.ConnectionString(suite.ctx)
	if err != nil {
		suite.T().Fatal(err)
	}
	suite.connUrl = connStr
	os.Setenv("DATABASE_URL", connStr)
	db, err := GetDB(connStr)
	if err != nil {
		suite.T().Log("Could not connect")
		suite.T().Fatal(err)
	}
	suite.DB = db
}

func (suite *TapSyncTestSuite) InitializeDaggerDB() {
	suite.T().Log("Using Dagger database")
	connUrl := "host=db user=testuser password=testpassword port=5432"
	// create tapsync database
	db, err := GetDB(connUrl)
	if err != nil {
		suite.T().Log("Could not connect")
		suite.T().Fatal(err)
	}
	_, err = db.Exec("CREATE DATABASE tapsync")
	if err != nil {
		suite.T().Log("Could not create tapsync database")
		suite.T().Fatal(err)
	}
	db.Close()
	// connect to tapsync database
	connUrl = connUrl + " dbname=tapsync"
	suite.connUrl = connUrl
	os.Setenv("DATABASE_URL", connUrl)
	db, err = GetDB(connUrl)
	if err != nil {
		suite.T().Log("Could not connect")
		suite.T().Fatal(err)
	}
	suite.DB = db
}
