package tapsync

import (
	"ataps/pkg/alercedb"
	"encoding/csv"
	"fmt"
	"net/http"
	"sort"
)

func test_get_csv_data_from_object(suite *AlerceTestSuite, table string, ensureObjectExistIn *string, overrideOid *string) [][]string {
	var oid string
	if overrideOid == nil {
		if ensureObjectExistIn == nil {
			row := suite.DB.QueryRow("SELECT oid FROM object LIMIT 1")
			if row.Err() != nil {
				suite.T().Fatal(row.Err())
			}
			row.Scan(&oid)
		} else {
			row := suite.DB.QueryRow(fmt.Sprintf("SELECT oid FROM %s LIMIT 1", *ensureObjectExistIn))
			if row.Err() != nil {
				suite.T().Fatal(row.Err())
			}
			row.Scan(&oid)
		}
	} else {
		oid = *overrideOid
	}
	response := SendTestQuery(fmt.Sprintf("LANG=PSQL&&FORMAT=csv&&QUERY=SELECT * FROM %s WHERE oid='%s'", table, oid), suite.Service)
	suite.Require().Equal(http.StatusOK, response.Code)
	reader := csv.NewReader(response.Body)
	records, err := reader.ReadAll()
	if err != nil {
		suite.T().Fatal(err)
	}
	return records

}

func (suite *AlerceTestSuite) TestCsv_Detections() {
	w := SendTestQuery("LANG=PSQL&&FORMAT=csv&&QUERY=SELECT * FROM detection LIMIT 1", suite.Service)
	suite.Require().Equal(http.StatusOK, w.Code)
	suite.Require().Equal("text/csv", w.Header().Get("Content-Type"))
	records, err := csv.NewReader(w.Body).ReadAll()
	if err != nil {
		suite.T().Fatal(err)
	}
	columnNames := GetColumnNames(alercedb.Detection{})
	sort.Strings(columnNames)
	suite.Require().Equal(columnNames, records[0])
}

func (suite *AlerceTestSuite) TestCsv_DetectionsFromAnObject() {
	ensureObjectExistIn := "detection"
	records := test_get_csv_data_from_object(suite, "detection", &ensureObjectExistIn, nil)
	suite.Require().Greater(len(records), 1)
}

func (suite *AlerceTestSuite) TestCsv_DetectionsFromUnknownObject() {
	var oid = "unknown"
	records := test_get_csv_data_from_object(suite, "detection", nil, &oid)
	suite.Require().Equal(0, len(records))
}

func (suite *AlerceTestSuite) TestCsv_NonDetection() {
	w := SendTestQuery("LANG=PSQL&&FORMAT=csv&&QUERY=SELECT * FROM non_detection LIMIT 1", suite.Service)
	suite.Require().Equal(http.StatusOK, w.Code)
	suite.Require().Equal("text/csv", w.Header().Get("Content-Type"))
	columnNames := GetColumnNames(alercedb.NonDetection{})
	records, err := csv.NewReader(w.Body).ReadAll()
	if err != nil {
		suite.T().Fatal(err)
	}
	sort.Strings(columnNames)
	suite.Require().Equal(columnNames, records[0])
}

func (suite *AlerceTestSuite) TestCsv_NonDetectionFromAnObject() {
	oidFrom := "non_detection"
	records := test_get_csv_data_from_object(suite, "non_detection", &oidFrom, nil)
	suite.Require().Greater(len(records), 1)
}

func (suite *AlerceTestSuite) TestCsv_NonDetectionFromUnknownObject() {
	var oid = "unknown"
	records := test_get_csv_data_from_object(suite, "non_detection", nil, &oid)
	suite.Require().Equal(0, len(records))
}

func (suite *AlerceTestSuite) TestCsv_ForcedPhotometry() {
	w := SendTestQuery("LANG=PSQL&&FORMAT=csv&&QUERY=SELECT * FROM forced_photometry LIMIT 1", suite.Service)
	suite.Require().Equal(http.StatusOK, w.Code)
	suite.Require().Equal("text/csv", w.Header().Get("Content-Type"))
	columnNames := GetColumnNames(alercedb.ForcedPhotometry{})
	records, err := csv.NewReader(w.Body).ReadAll()
	if err != nil {
		suite.T().Fatal(err)
	}
	sort.Strings(columnNames)
	suite.Require().Equal(columnNames, records[0])
}

func (suite *AlerceTestSuite) TestCsv_ForcedPhotometryFromAnObject() {
	oidFrom := "forced_photometry"
	records := test_get_csv_data_from_object(suite, "forced_photometry", &oidFrom, nil)
	suite.Require().Greater(len(records), 1)
}

func (suite *AlerceTestSuite) TestCsv_ForcedPhotometryFromUnknownObject() {
	var oid = "unknown"
	records := test_get_csv_data_from_object(suite, "forced_photometry", nil, &oid)
	suite.Require().Equal(0, len(records))
}

func (suite *AlerceTestSuite) TestCsv_Features() {
	w := SendTestQuery("LANG=PSQL&&FORMAT=csv&&QUERY=SELECT * FROM feature LIMIT 1", suite.Service)
	suite.Require().Equal(http.StatusOK, w.Code)
	suite.Require().Equal("text/csv", w.Header().Get("Content-Type"))
	columnNames := GetColumnNames(alercedb.Feature{})
	records, err := csv.NewReader(w.Body).ReadAll()
	if err != nil {
		suite.T().Fatal(err)
	}
	sort.Strings(columnNames)
	suite.Require().Equal(columnNames, records[0])
}

func (suite *AlerceTestSuite) TestCsv_FeaturesFromAnObject() {
	oidFrom := "feature"
	records := test_get_csv_data_from_object(suite, "feature", &oidFrom, nil)
	suite.Require().Greater(len(records), 1)
}

func (suite *AlerceTestSuite) TestCsv_FeaturesFromUnknownObject() {
	var oid = "unknown"
	records := test_get_csv_data_from_object(suite, "feature", nil, &oid)
	suite.Require().Equal(0, len(records))
}

func (suite *AlerceTestSuite) TestCsv_Objects() {
	w := SendTestQuery("LANG=PSQL&&FORMAT=csv&&QUERY=SELECT * FROM object LIMIT 2", suite.Service)
	suite.Require().Equal(http.StatusOK, w.Code)
	suite.Require().Equal("text/csv", w.Header().Get("Content-Type"))
	columnNames := GetColumnNames(alercedb.Object{})
	records, err := csv.NewReader(w.Body).ReadAll()
	if err != nil {
		suite.T().Fatal(err)
	}
	sort.Strings(columnNames)
	suite.Require().Equal(columnNames, records[0])
	suite.Require().Equal(3, len(records)) // header + 2 objects
}

func (suite *AlerceTestSuite) TestCsv_Probabilities() {
	w := SendTestQuery("LANG=PSQL&&FORMAT=csv&&QUERY=SELECT * FROM probability LIMIT 1", suite.Service)
	suite.Require().Equal(http.StatusOK, w.Code)
	suite.Require().Equal("text/csv", w.Header().Get("Content-Type"))
	columnNames := GetColumnNames(alercedb.Probability{})
	records, err := csv.NewReader(w.Body).ReadAll()
	if err != nil {
		suite.T().Fatal(err)
	}
	sort.Strings(columnNames)
	suite.Require().Equal(columnNames, records[0])
}

func (suite *AlerceTestSuite) TestCsv_ProbabilitiesFromAnObject() {
	oidFrom := "probability"
	records := test_get_csv_data_from_object(suite, "probability", &oidFrom, nil)
	suite.Require().Greater(len(records), 1)
}
