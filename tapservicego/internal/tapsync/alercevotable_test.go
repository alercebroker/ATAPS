package tapsync

import (
	"ataps/pkg/alercedb"
	"ataps/pkg/votable"
	"net/http"
)

func (suite *AlerceTestSuite) TestVotable_Object() {
	w := SendTestQuery("LANG=PSQL&&FORMAT=votable&&QUERY=SELECT * FROM object LIMIT 3", suite.Service)
	suite.Require().Equal(http.StatusOK, w.Code)
	suite.Require().Equal("application/x-votable+xml", w.Header().Get("Content-Type"))
	voTable, err := votable.NewVOTableFromString(w.Body.String())
	suite.Require().NoError(err)
	columnNames := GetColumnNames(alercedb.Object{})
	suite.Require().Len(voTable.Resource.Tables[0].Fields, len(columnNames))
	for _, field := range voTable.Resource.Tables[0].Fields {
		suite.Require().Contains(columnNames, field.Name)
	}
	suite.Require().Len(voTable.Resource.Tables[0].Data.TableData.Rows, 3)
}

func (suite *AlerceTestSuite) TestVotable_Detection() {
	w := SendTestQuery("LANG=PSQL&&FORMAT=votable&&QUERY=SELECT * FROM detection LIMIT 3", suite.Service)
	suite.Require().Equal(http.StatusOK, w.Code)
	suite.Require().Equal("application/x-votable+xml", w.Header().Get("Content-Type"))
	voTable, err := votable.NewVOTableFromString(w.Body.String())
	suite.Require().NoError(err)
	columnNames := GetColumnNames(alercedb.Detection{})
	suite.Require().Len(voTable.Resource.Tables[0].Fields, len(columnNames))
	for _, field := range voTable.Resource.Tables[0].Fields {
		suite.Require().Contains(columnNames, field.Name)
	}
	suite.Require().Len(voTable.Resource.Tables[0].Data.TableData.Rows, 3)
}

func (suite *AlerceTestSuite) TestVotable_NonDetection() {
	w := SendTestQuery("LANG=PSQL&&FORMAT=votable&&QUERY=SELECT * FROM non_detection LIMIT 3", suite.Service)
	suite.Require().Equal(http.StatusOK, w.Code)
	suite.Require().Equal("application/x-votable+xml", w.Header().Get("Content-Type"))
	voTable, err := votable.NewVOTableFromString(w.Body.String())
	suite.Require().NoError(err)
	columnNames := GetColumnNames(alercedb.NonDetection{})
	suite.Require().Len(voTable.Resource.Tables[0].Fields, len(columnNames))
	for _, field := range voTable.Resource.Tables[0].Fields {
		suite.Require().Contains(columnNames, field.Name)
	}
	suite.Require().Len(voTable.Resource.Tables[0].Data.TableData.Rows, 3)
}

func (suite *AlerceTestSuite) TestVotable_ForcedPhotometry() {
	w := SendTestQuery("LANG=PSQL&&FORMAT=votable&&QUERY=SELECT * FROM forced_photometry LIMIT 3", suite.Service)
	suite.Require().Equal(http.StatusOK, w.Code)
	suite.Require().Equal("application/x-votable+xml", w.Header().Get("Content-Type"))
	voTable, err := votable.NewVOTableFromString(w.Body.String())
	suite.Require().NoError(err)
	columnNames := GetColumnNames(alercedb.ForcedPhotometry{})
	suite.Require().Len(voTable.Resource.Tables[0].Fields, len(columnNames))
	for _, field := range voTable.Resource.Tables[0].Fields {
		suite.Require().Contains(columnNames, field.Name)
	}
	suite.Require().Len(voTable.Resource.Tables[0].Data.TableData.Rows, 3)
}

func (suite *AlerceTestSuite) TestVotable_Features() {
	w := SendTestQuery("LANG=PSQL&&FORMAT=votable&&QUERY=SELECT * FROM feature LIMIT 3", suite.Service)
	suite.Require().Equal(http.StatusOK, w.Code)
	suite.Require().Equal("application/x-votable+xml", w.Header().Get("Content-Type"))
	voTable, err := votable.NewVOTableFromString(w.Body.String())
	suite.Require().NoError(err)
	columnNames := GetColumnNames(alercedb.Feature{})
	suite.Require().Len(voTable.Resource.Tables[0].Fields, len(columnNames))
	for _, field := range voTable.Resource.Tables[0].Fields {
		suite.Require().Contains(columnNames, field.Name)
	}
	suite.Require().Len(voTable.Resource.Tables[0].Data.TableData.Rows, 3)
}

func (suite *AlerceTestSuite) TestVotable_Probabilities() {
	w := SendTestQuery("LANG=PSQL&&FORMAT=votable&&QUERY=SELECT * FROM probability LIMIT 3", suite.Service)
	suite.Require().Equal(http.StatusOK, w.Code)
	suite.Require().Equal("application/x-votable+xml", w.Header().Get("Content-Type"))
	voTable, err := votable.NewVOTableFromString(w.Body.String())
	suite.Require().NoError(err)
	columnNames := GetColumnNames(alercedb.Probability{})
	suite.Require().Len(voTable.Resource.Tables[0].Fields, len(columnNames))
	for _, field := range voTable.Resource.Tables[0].Fields {
		suite.Require().Contains(columnNames, field.Name)
	}
	suite.Require().Len(voTable.Resource.Tables[0].Data.TableData.Rows, 3)
}
