package tapsync

import (
	"ataps/pkg/alercedb"
	"net/http"

	"golang.org/x/net/html"
)

func (suite *AlerceTestSuite) TestHtml_Object() {
	w := SendTestQuery("LANG=PSQL&&FORMAT=html&&QUERY=SELECT * FROM object LIMIT 3", suite.Service)
	suite.Require().Equal(http.StatusOK, w.Code)
	suite.Require().Equal("text/html", w.Header().Get("Content-Type"))
	var data []string
	var headers []string
	doc, err := html.Parse(w.Body)
	suite.Require().NoError(err)
	ParseHTMLTable(doc, &headers, "th")
	ParseHTMLTable(doc, &data, "td")
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

func (suite *AlerceTestSuite) TestHtml_NonExistentTable() {
	w := SendTestQuery("LANG=PSQL&&FORMAT=html&&QUERY=SELECT * FROM non_existent_table LIMIT 3", suite.Service)
	suite.Require().Equal(http.StatusInternalServerError, w.Code)
}

func (suite *AlerceTestSuite) TestHtml_Detection() {
	w := SendTestQuery("LANG=PSQL&&FORMAT=html&&QUERY=SELECT * FROM detection LIMIT 3", suite.Service)
	suite.Require().Equal(http.StatusOK, w.Code)
	suite.Require().Equal("text/html", w.Header().Get("Content-Type"))
	var data []string
	var headers []string
	doc, err := html.Parse(w.Body)
	suite.Require().NoError(err)
	ParseHTMLTable(doc, &headers, "th")
	ParseHTMLTable(doc, &data, "td")
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

func (suite *AlerceTestSuite) TestHtml_NonDetection() {
	w := SendTestQuery("LANG=PSQL&&FORMAT=html&&QUERY=SELECT * FROM non_detection LIMIT 3", suite.Service)
	suite.Require().Equal(http.StatusOK, w.Code)
	suite.Require().Equal("text/html", w.Header().Get("Content-Type"))
	var data []string
	var headers []string
	doc, err := html.Parse(w.Body)
	suite.Require().NoError(err)
	ParseHTMLTable(doc, &headers, "th")
	ParseHTMLTable(doc, &data, "td")
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

func (suite *AlerceTestSuite) TestHtml_ForcedPhotometry() {
	w := SendTestQuery("LANG=PSQL&&FORMAT=html&&QUERY=SELECT * FROM forced_photometry LIMIT 3", suite.Service)
	suite.Require().Equal(http.StatusOK, w.Code)
	suite.Require().Equal("text/html", w.Header().Get("Content-Type"))
	var data []string
	var headers []string
	doc, err := html.Parse(w.Body)
	suite.Require().NoError(err)
	ParseHTMLTable(doc, &headers, "th")
	ParseHTMLTable(doc, &data, "td")
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

func (suite *AlerceTestSuite) TestHtml_Features() {
	w := SendTestQuery("LANG=PSQL&&FORMAT=html&&QUERY=SELECT * FROM feature LIMIT 3", suite.Service)
	suite.Require().Equal(http.StatusOK, w.Code)
	suite.Require().Equal("text/html", w.Header().Get("Content-Type"))
	var data []string
	var headers []string
	doc, err := html.Parse(w.Body)
	suite.Require().NoError(err)
	ParseHTMLTable(doc, &headers, "th")
	ParseHTMLTable(doc, &data, "td")
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

func (suite *AlerceTestSuite) TestHtml_Probabilities() {
	w := SendTestQuery("LANG=PSQL&&FORMAT=html&&QUERY=SELECT * FROM probability LIMIT 3", suite.Service)
	suite.Require().Equal(http.StatusOK, w.Code)
	suite.Require().Equal("text/html", w.Header().Get("Content-Type"))
	var data []string
	var headers []string
	doc, err := html.Parse(w.Body)
	suite.Require().NoError(err)
	ParseHTMLTable(doc, &headers, "th")
	ParseHTMLTable(doc, &data, "td")
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
