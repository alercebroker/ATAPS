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

type PsqlRepositoryTestSuite struct {
	suite.Suite
	dbcontainer *postgres.PostgresContainer
	context    context.Context
}

func populateDb (db *sql.DB) error {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS test (id SERIAL PRIMARY KEY, name TEXT, number INT)")
	if err != nil {
		return err
	}
	_, err = db.Exec("INSERT INTO test (name, number) VALUES ('test', 1)")
	if err != nil {
		return err
	}
	return nil
}

func (suite *PsqlRepositoryTestSuite) SetupSuite() {
	ctx := context.Background()
	suite.context = ctx
	dbcontainer, err := testhelpers.CreatePostgresContainer(ctx)
	if err != nil {
		suite.T().Fatal(err)
	}
	suite.dbcontainer = dbcontainer
	connStr, err := dbcontainer.ConnectionString(ctx)
	if err != nil {
		suite.T().Fatal(err)
	}
	os.Setenv("DATABASE_URL", connStr)
	db, err := GetDB(connStr)
	if err != nil {
		suite.T().Fatal(err)
	}
	err = populateDb(db)
	if err != nil {
		suite.T().Fatal(err)
	}
}

func (suite *PsqlRepositoryTestSuite) TearDownSuite() {
	testhelpers.CleanUpContainer(suite.context, suite.dbcontainer)	
}

func (suite *PsqlRepositoryTestSuite) TestHandleSQLQuery() {
	db, err := GetDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		suite.T().Fatal(err)
	}
	query := "SELECT * FROM test"
	result, err := HandleSQLQuery(query, db)
	if err != nil {
		suite.T().Fatal(err)
	}
	suite.Equal(1, len(result))
	suite.Equal("test", result[0]["name"])
	suite.Equal(int64(1), result[0]["number"])
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestPsqlRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(PsqlRepositoryTestSuite))
}
