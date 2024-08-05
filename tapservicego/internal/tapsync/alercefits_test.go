package tapsync

import (
	"ataps/pkg/alercedb"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/astrogo/cfitsio"
)

func (suite *AlerceTestSuite) TestFits_Object() {
	w := SendTestQuery("LANG=PSQL&&FORMAT=fits&&QUERY=SELECT * FROM object LIMIT 3", suite.Service)
	suite.Require().Equal(http.StatusOK, w.Code)
	suite.Require().Equal("application/fits", w.Header().Get("Content-Type"))
	fname := writeFitsFile(suite.T(), w)
	defer os.Remove(fname)
	f, err := cfitsio.Open(fname, cfitsio.ReadOnly)
	suite.Require().Nil(err)
	defer f.Close()
	hdu := f.HDU(1)
	table := hdu.(*cfitsio.Table)
	suite.Require().True(table.Type() == cfitsio.BINARY_TBL)
	suite.Require().Equal("results", table.Name())
	suite.Require().Equal(10, table.NumCols())
	suite.Require().Equal(int64(3), table.NumRows())
	rows, err := table.Read(0, table.NumRows())
	suite.Require().Nil(err)
	defer rows.Close()
	count, parsedData := parseResponseData(suite, rows)
	columnNames := GetColumnNames(alercedb.Object{})
	assertColumnsExist(suite, columnNames, parsedData)
	suite.Require().Equal(3, count)
}

func (suite *AlerceTestSuite) TestFits_Detection() {
	w := SendTestQuery("LANG=PSQL&&FORMAT=fits&&QUERY=SELECT * FROM detection LIMIT 3", suite.Service)
	suite.Require().Equal(http.StatusOK, w.Code)
	suite.Require().Equal("application/fits", w.Header().Get("Content-Type"))
	fname := writeFitsFile(suite.T(), w)
	defer os.Remove(fname)
	f, err := cfitsio.Open(fname, cfitsio.ReadOnly)
	suite.Require().Nil(err)
	defer f.Close()
	hdu := f.HDU(1)
	table := hdu.(*cfitsio.Table)
	suite.Require().True(table.Type() == cfitsio.BINARY_TBL)
	suite.Require().Equal("results", table.Name())
	suite.Require().Equal(19, table.NumCols())
	suite.Require().Equal(int64(3), table.NumRows())
	rows, err := table.Read(0, table.NumRows())
	suite.Require().Nil(err)
	defer rows.Close()
	count, parsedData := parseResponseData(suite, rows)
	columnNames := GetColumnNames(alercedb.Detection{})
	assertColumnsExist(suite, columnNames, parsedData)
	suite.Require().Equal(3, count)
}

func (suite *AlerceTestSuite) TestFits_NonDetection() {
	w := SendTestQuery("LANG=PSQL&&FORMAT=fits&&QUERY=SELECT * FROM non_detection LIMIT 3", suite.Service)
	suite.Require().Equal(http.StatusOK, w.Code)
	suite.Require().Equal("application/fits", w.Header().Get("Content-Type"))
	fname := writeFitsFile(suite.T(), w)
	defer os.Remove(fname)
	f, err := cfitsio.Open(fname, cfitsio.ReadOnly)
	suite.Require().Nil(err)
	defer f.Close()
	hdu := f.HDU(1)
	table := hdu.(*cfitsio.Table)
	suite.Require().True(table.Type() == cfitsio.BINARY_TBL)
	suite.Require().Equal("results", table.Name())
	suite.Require().Equal(4, table.NumCols())
	suite.Require().Equal(int64(3), table.NumRows())
	rows, err := table.Read(0, table.NumRows())
	suite.Require().Nil(err)
	defer rows.Close()
	count, parsedData := parseResponseData(suite, rows)
	columnNames := GetColumnNames(alercedb.NonDetection{})
	assertColumnsExist(suite, columnNames, parsedData)
	suite.Require().Equal(3, count)
}

func (suite *AlerceTestSuite) TestFits_ForcedPhotometry() {
	w := SendTestQuery("LANG=PSQL&&FORMAT=fits&&QUERY=SELECT * FROM forced_photometry LIMIT 3", suite.Service)
	suite.Require().Equal(http.StatusOK, w.Code)
	suite.Require().Equal("application/fits", w.Header().Get("Content-Type"))
	fname := writeFitsFile(suite.T(), w)
	defer os.Remove(fname)
	f, err := cfitsio.Open(fname, cfitsio.ReadOnly)
	suite.Require().Nil(err)
	defer f.Close()
	hdu := f.HDU(1)
	table := hdu.(*cfitsio.Table)
	suite.Require().True(table.Type() == cfitsio.BINARY_TBL)
	suite.Require().Equal("results", table.Name())
	suite.Require().Equal(19, table.NumCols())
	suite.Require().Equal(int64(3), table.NumRows())
	rows, err := table.Read(0, table.NumRows())
	suite.Require().Nil(err)
	defer rows.Close()
	count, parsedData := parseResponseData(suite, rows)
	columnNames := GetColumnNames(alercedb.ForcedPhotometry{})
	assertColumnsExist(suite, columnNames, parsedData)
	suite.Require().Equal(3, count)
}

func (suite *AlerceTestSuite) TestFits_Features() {
	w := SendTestQuery("LANG=PSQL&&FORMAT=fits&&QUERY=SELECT * FROM feature LIMIT 3", suite.Service)
	suite.Require().Equal(http.StatusOK, w.Code)
	suite.Require().Equal("application/fits", w.Header().Get("Content-Type"))
	fname := writeFitsFile(suite.T(), w)
	defer os.Remove(fname)
	f, err := cfitsio.Open(fname, cfitsio.ReadOnly)
	suite.Require().Nil(err)
	defer f.Close()
	hdu := f.HDU(1)
	table := hdu.(*cfitsio.Table)
	suite.Require().True(table.Type() == cfitsio.BINARY_TBL)
	suite.Require().Equal("results", table.Name())
	suite.Require().Equal(5, table.NumCols())
	suite.Require().Equal(int64(3), table.NumRows())
	rows, err := table.Read(0, table.NumRows())
	suite.Require().Nil(err)
	defer rows.Close()
	count, parsedData := parseResponseData(suite, rows)
	columnNames := GetColumnNames(alercedb.Feature{})
	assertColumnsExist(suite, columnNames, parsedData)
	suite.Require().Equal(3, count)
}

func (suite *AlerceTestSuite) TestFits_Probabilities() {
	w := SendTestQuery("LANG=PSQL&&FORMAT=fits&&QUERY=SELECT * FROM probability LIMIT 3", suite.Service)
	suite.Require().Equal(http.StatusOK, w.Code)
	suite.Require().Equal("application/fits", w.Header().Get("Content-Type"))
	fname := writeFitsFile(suite.T(), w)
	defer os.Remove(fname)
	f, err := cfitsio.Open(fname, cfitsio.ReadOnly)
	suite.Require().Nil(err)
	defer f.Close()
	hdu := f.HDU(1)
	table := hdu.(*cfitsio.Table)
	suite.Require().True(table.Type() == cfitsio.BINARY_TBL)
	suite.Require().Equal("results", table.Name())
	suite.Require().Equal(6, table.NumCols())
	suite.Require().Equal(int64(3), table.NumRows())
	rows, err := table.Read(0, table.NumRows())
	suite.Require().Nil(err)
	defer rows.Close()
	count, parsedData := parseResponseData(suite, rows)
	columnNames := GetColumnNames(alercedb.Probability{})
	assertColumnsExist(suite, columnNames, parsedData)
	suite.Require().Equal(3, count)
}

func parseResponseData(suite *AlerceTestSuite, rows *cfitsio.Rows) (int, []map[string]interface{}) {
	parsedData := []map[string]interface{}{}
	count := 0
	for rows.Next() {
		row := map[string]interface{}{}
		err := rows.Scan(&row)
		suite.Require().Nil(err)
		parsedData = append(parsedData, row)
		count = count + 1
	}
	return count, parsedData
}

func assertColumnsExist(suite *AlerceTestSuite, columnNames []string, parsedData []map[string]interface{}) {
	for _, data := range parsedData {
		for _, columnName := range columnNames {
			_, ok := data[columnName]
			suite.Require().True(ok)
		}
	}
}

func writeFitsFile(t *testing.T, w *httptest.ResponseRecorder) string {
	f, err := os.CreateTemp("", "result.fits")
	if err != nil {
		t.Fatal(err)
	}
	f.Write(w.Body.Bytes())
	defer f.Close()
	return f.Name()
}
