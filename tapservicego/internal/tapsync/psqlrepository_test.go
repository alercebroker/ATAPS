package tapsync

import (
	"ataps/internal/testhelpers"
)

func (suite *TapSyncTestSuite) TestSimpleSQLQuery() {
	query := "SELECT 'test'"
	result, err := HandleSQLQuery(query, suite.DB)
	if err != nil {
		suite.T().Fatal(err)
	}
	suite.Equal(1, len(result))
	suite.Equal("test", result[0]["?column?"])
}

func (suite *TapSyncTestSuite) TestHandleSQLQuery() {
	testhelpers.PopulateDb(suite.DB)
	defer testhelpers.ClearDataFromTable(suite.DB)
	query := "SELECT * FROM test"
	result, err := HandleSQLQuery(query, suite.DB)
	if err != nil {
		suite.T().Fatal(err)
	}
	suite.Equal(1, len(result))
	suite.Equal("test", result[0]["name"])
	suite.Equal(int64(1), result[0]["number"])
}
