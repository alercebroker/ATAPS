package tapsync

import (
	"ataps/pkg/alercedb"
	"net/http"
)

func (suite *AlerceTestSuite) TestText_Object() {
	w := SendTestQuery("LANG=PSQL&&FORMAT=text&&QUERY=SELECT * FROM object LIMIT 3", suite.Service)
	suite.Require().Equal(http.StatusOK, w.Code)
	suite.Require().Equal("text/plain", w.Header().Get("Content-Type"))
	var data []string
	var headers []string
	err := ParseTextTable(w.Body.String(), &data, &headers)
	suite.Require().NoError(err)
	var rows []map[string]string
	for i := 0; i < len(data); i += len(headers) {
		row := make(map[string]string, len(headers))
		for j, header := range headers {
			row[header] = data[i+j]
		}
		rows = append(rows, row)
	}
	suite.Require().Len(rows, 3)
	columnNames := GetColumnNames(alercedb.Object{})
	suite.Require().ElementsMatch(headers, columnNames)
}

func (suite *AlerceTestSuite) TestText_Detection() {
	w := SendTestQuery("LANG=PSQL&&FORMAT=text&&QUERY=SELECT * FROM detection LIMIT 3", suite.Service)
	suite.Require().Equal(http.StatusOK, w.Code)
	suite.Require().Equal("text/plain", w.Header().Get("Content-Type"))
	var data []string
	var headers []string
	err := ParseTextTable(w.Body.String(), &data, &headers)
	suite.Require().NoError(err)
	var rows []map[string]string
	for i := 0; i < len(data); i += len(headers) {
		row := make(map[string]string, len(headers))
		for j, header := range headers {
			row[header] = data[i+j]
		}
		rows = append(rows, row)
	}
	suite.Require().Len(rows, 3)
	columnNames := GetColumnNames(alercedb.Detection{})
	suite.Require().ElementsMatch(headers, columnNames)
}

func (suite *AlerceTestSuite) TestText_NonDetection() {
	w := SendTestQuery("LANG=PSQL&&FORMAT=text&&QUERY=SELECT * FROM non_detection LIMIT 3", suite.Service)
	suite.Require().Equal(http.StatusOK, w.Code)
	suite.Require().Equal("text/plain", w.Header().Get("Content-Type"))
	var data []string
	var headers []string
	err := ParseTextTable(w.Body.String(), &data, &headers)
	suite.Require().NoError(err)
	var rows []map[string]string
	for i := 0; i < len(data); i += len(headers) {
		row := make(map[string]string, len(headers))
		for j, header := range headers {
			row[header] = data[i+j]
		}
		rows = append(rows, row)
	}
	suite.Require().Len(rows, 3)
	columnNames := GetColumnNames(alercedb.NonDetection{})
	suite.Require().ElementsMatch(headers, columnNames)
}

func (suite *AlerceTestSuite) TestText_ForcedPhotometry() {
	w := SendTestQuery("LANG=PSQL&&FORMAT=text&&QUERY=SELECT * FROM forced_photometry LIMIT 3", suite.Service)
	suite.Require().Equal(http.StatusOK, w.Code)
	suite.Require().Equal("text/plain", w.Header().Get("Content-Type"))
	var data []string
	var headers []string
	err := ParseTextTable(w.Body.String(), &data, &headers)
	suite.Require().NoError(err)
	var rows []map[string]string
	for i := 0; i < len(data); i += len(headers) {
		row := make(map[string]string, len(headers))
		for j, header := range headers {
			row[header] = data[i+j]
		}
		rows = append(rows, row)
	}
	suite.Require().Len(rows, 3)
	columnNames := GetColumnNames(alercedb.ForcedPhotometry{})
	suite.Require().ElementsMatch(headers, columnNames)
}

func (suite *AlerceTestSuite) TestText_Features() {
	w := SendTestQuery("LANG=PSQL&&FORMAT=text&&QUERY=SELECT * FROM feature LIMIT 3", suite.Service)
	suite.Require().Equal(http.StatusOK, w.Code)
	suite.Require().Equal("text/plain", w.Header().Get("Content-Type"))
	var data []string
	var headers []string
	err := ParseTextTable(w.Body.String(), &data, &headers)
	suite.Require().NoError(err)
	var rows []map[string]string
	for i := 0; i < len(data); i += len(headers) {
		row := make(map[string]string, len(headers))
		for j, header := range headers {
			row[header] = data[i+j]
		}
		rows = append(rows, row)
	}
	suite.Require().Len(rows, 3)
	columnNames := GetColumnNames(alercedb.Feature{})
	suite.Require().ElementsMatch(headers, columnNames)
}

func (suite *AlerceTestSuite) TestText_Probabilities() {
	w := SendTestQuery("LANG=PSQL&&FORMAT=text&&QUERY=SELECT * FROM probability LIMIT 3", suite.Service)
	suite.Require().Equal(http.StatusOK, w.Code)
	suite.Require().Equal("text/plain", w.Header().Get("Content-Type"))
	var data []string
	var headers []string
	err := ParseTextTable(w.Body.String(), &data, &headers)
	suite.Require().NoError(err)
	var rows []map[string]string
	for i := 0; i < len(data); i += len(headers) {
		row := make(map[string]string, len(headers))
		for j, header := range headers {
			row[header] = data[i+j]
		}
		rows = append(rows, row)
	}
	suite.Require().Len(rows, 3)
	columnNames := GetColumnNames(alercedb.Probability{})
	suite.Require().ElementsMatch(headers, columnNames)
}
