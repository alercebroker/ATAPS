package tapsync

import (
	"ataps/internal/testhelpers"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

type TAPSyncTestSuite struct {
	suite.Suite
	dbcontainer *postgres.PostgresContainer
	context    context.Context
}

func (suite *TAPSyncTestSuite) SetupSuite() {
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
	fmt.Println(connStr)
	os.Setenv("DATABASE_URL", connStr)
}

func (suite *TAPSyncTestSuite) TearDownSuite() {
	testhelpers.CleanUpContainer(suite.context, suite.dbcontainer)	
}

func (suite *TAPSyncTestSuite) TestLangSuccess() {
	service := NewTapSyncService()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/sync", strings.NewReader("LANG=PSQL&&QUERY=SELECT 'test'"))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	service.Router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)
}

func (suite *TAPSyncTestSuite) TestLangFailure() {
	service := NewTapSyncService()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/sync", strings.NewReader(""))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	service.Router.ServeHTTP(w, req)

	suite.Equal(http.StatusBadRequest, w.Code)
}

func (suite *TAPSyncTestSuite) TestFormatSuccess() {
	service := NewTapSyncService()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/sync", strings.NewReader("LANG=PSQL&&FORMAT=votable&&QUERY=SELECT 'test'"))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	service.Router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)
}

func (suite *TAPSyncTestSuite) TestResponseFormatSuccess() {
	service := NewTapSyncService()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/sync", strings.NewReader("LANG=PSQL&&RESPONSEFORMAT=votable&&QUERY=SELECT 'test'"))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	service.Router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)
}

func (suite *TAPSyncTestSuite) TestFormatSuccessWithoutSpecifying() {
	service := NewTapSyncService()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/sync", strings.NewReader("LANG=PSQL&&QUERY=SELECT 'test'"))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	service.Router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)
}

func (suite *TAPSyncTestSuite) TestFormatFailureWhenProvidingBoth() {
	service := NewTapSyncService()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/sync", strings.NewReader("LANG=PSQL&&FORMAT=votable&&RESPONSEFORMAT=votable"))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	service.Router.ServeHTTP(w, req)

	suite.Equal(http.StatusBadRequest, w.Code)
}

func (suite *TAPSyncTestSuite) TestFormatFailureWhenProvidingUnknownFormat() {
	service := NewTapSyncService()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/sync", strings.NewReader("LANG=PSQL&&FORMAT=Unknown"))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	service.Router.ServeHTTP(w, req)

	suite.Equal(http.StatusBadRequest, w.Code)
}
func (suite *TAPSyncTestSuite) TestFormatFailureWhenProvidingUnknownResponseFormat() {
	service := NewTapSyncService()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/sync", strings.NewReader("LANG=PSQL&&RESPONSEFORMAT=Unknown"))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	service.Router.ServeHTTP(w, req)

	suite.Equal(http.StatusBadRequest, w.Code)
}
func (suite *TAPSyncTestSuite) TestBadRequestIfQueryIsEmpty() {
	service := NewTapSyncService()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/sync", strings.NewReader("LANG=PSQL"))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	service.Router.ServeHTTP(w, req)

	suite.Equal(http.StatusBadRequest, w.Code)
}

func (suite *TAPSyncTestSuite) TestPerformQuerySuccess() {
	service := NewTapSyncService()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/sync", strings.NewReader("LANG=PSQL&&QUERY=SELECT 'test'"))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	service.Router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestTAPSyncTestSuite(t *testing.T) {
	suite.Run(t, new(TAPSyncTestSuite))
}
