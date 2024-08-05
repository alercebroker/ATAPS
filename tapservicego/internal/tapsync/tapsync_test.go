package tapsync

import (
	"ataps/internal/testhelpers"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/astrogo/cfitsio"
	"github.com/stretchr/testify/assert"
)

func (suite *TapSyncTestSuite) TestQueryParams() {
	t := suite.T()
	t.Run("TestLangSuccess", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/sync", strings.NewReader("LANG=PSQL&&QUERY=SELECT 'test'"))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		suite.Service.Router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("TestLangFailure", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/sync", strings.NewReader(""))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		suite.Service.Router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "Invalid LANG ")
	})
	t.Run("TestFormatSuccess", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/sync", strings.NewReader("LANG=PSQL&&FORMAT=votable&&QUERY=SELECT 'test'"))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		suite.Service.Router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("TestResponseFormatSuccess", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/sync", strings.NewReader("LANG=PSQL&&RESPONSEFORMAT=votable&&QUERY=SELECT 'test'"))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		suite.Service.Router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("TestFormatSuccessWithoutSpecifying", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/sync", strings.NewReader("LANG=PSQL&&QUERY=SELECT 'test'"))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		suite.Service.Router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("TestFormatFailureWhenProvidingBoth", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/sync", strings.NewReader("LANG=PSQL&&FORMAT=votable&&RESPONSEFORMAT=votable"))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		suite.Service.Router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
	t.Run("TestFormatFailureWhenProvidingUnknownFormat", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/sync", strings.NewReader("LANG=PSQL&&FORMAT=Unknown"))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		suite.Service.Router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
	t.Run("TestFormatFailureWhenProvidingUnknownResponseFormat", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/sync", strings.NewReader("LANG=PSQL&&RESPONSEFORMAT=Unknown"))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		suite.Service.Router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
	t.Run("TestBadRequestIfQueryIsEmpty", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/sync", strings.NewReader("LANG=PSQL"))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		suite.Service.Router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func (suite *TapSyncTestSuite) TestCSVQueries() {
	t := suite.T()
	t.Run("TestCSVQuerySuccess", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/sync", strings.NewReader("LANG=PSQL&&FORMAT=csv&&QUERY=SELECT 'test'"))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		suite.Service.Router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "?column?\ntest\n", w.Body.String())
		assert.Equal(t, "text/csv", w.Header().Get("Content-Type"))
	})
	t.Run("TestTSVQuerySuccess", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/sync", strings.NewReader("LANG=PSQL&&FORMAT=tsv&&QUERY=SELECT 'test'"))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		suite.Service.Router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "?column?\ntest\n", w.Body.String())
		assert.Equal(t, "text/tab-separated-values", w.Header().Get("Content-Type"))
	})
	t.Run("TestCSVQueryFailure", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/sync", strings.NewReader("LANG=PSQL&&FORMAT=csv&&QUERY=SELECT * from dontexist"))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		suite.Service.Router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, "application/xml; charset=utf-8", w.Header().Get("Content-Type"))
		// the default gin xml render does not show quotes
		assert.Contains(t, w.Body.String(), "relation &#34;dontexist&#34; does not exist")
	})
}

func (suite *TapSyncTestSuite) TestVOTableQueries() {
	t := suite.T()
	t.Run("TestVOTableQuerySuccess", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/sync", strings.NewReader("LANG=PSQL&&FORMAT=votable&&QUERY=SELECT 'test'"))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		suite.Service.Router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, `<?xml version="1.0" encoding="UTF-8"?>
<VOTABLE version="1.4" xmlns="http://www.ivoa.net/xml/VOTable/v1.4">
	<RESOURCE type="results">
		<INFO name="QUERY_STATUS" value="OK"></INFO>
		<TABLE name="results">
			<DESCRIPTION>Results of the query</DESCRIPTION>
			<FIELD name="?column?" datatype="char" arraysize="100*"></FIELD>
			<DATA>
				<TABLEDATA>
					<TR>
						<TD>test</TD>
					</TR>
				</TABLEDATA>
			</DATA>
		</TABLE>
	</RESOURCE>
</VOTABLE>`, w.Body.String())
	})
}

func (suite *TapSyncTestSuite) TestFitsQueries() {
	t := suite.T()
	t.Run("TestFitsQuerySuccess", func(t *testing.T) {
		err := testhelpers.PopulateDb(suite.DB)
		if err != nil {
			t.Fatal(err)
		}
		defer testhelpers.ClearDataFromTable(suite.DB)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/sync", strings.NewReader("LANG=PSQL&&FORMAT=fits&&QUERY=SELECT * FROM test"))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		suite.Service.Router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "application/fits", w.Header().Get("Content-Type"))
		// read the fits file and parse it
		fname := writeFitsFile(t, w)
		defer os.Remove(fname)
		f, err := cfitsio.Open(fname, cfitsio.ReadOnly)
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()
		hdu := f.HDU(1)
		table := hdu.(*cfitsio.Table)
		assert.Equal(t, "results", table.Name())
		assert.Equal(t, 3, table.NumCols()) // id, name, number
		assert.Equal(t, int64(1), table.NumRows())
		assert.Equal(t, "id", table.Col(0).Name)
		assert.Equal(t, "name", table.Col(1).Name)
		assert.Equal(t, "number", table.Col(2).Name)
		rows, err := table.Read(0, 2)
		if err != nil {
			t.Fatal(err)
		}
		count := 0
		for rows.Next() {
			var id int
			var name string
			var number int
			err = rows.Scan(&id, &name, &number)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, "test", name)
			assert.Equal(t, 1, number)
			count++
		}
		assert.Equal(t, 1, count)
	})
}

func (suite *TapSyncTestSuite) TestTextQueries() {
	t := suite.T()
	t.Run("TestTextQuerySuccess", func(t *testing.T) {
		testhelpers.PopulateDb(suite.DB)
		defer testhelpers.ClearDataFromTable(suite.DB)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/sync", strings.NewReader("LANG=PSQL&&FORMAT=text&&QUERY=SELECT * FROM test"))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		suite.Service.Router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "text/plain", w.Header().Get("Content-Type"))
		assert.Contains(t, w.Body.String(), "id | name | number")
		assert.Contains(t, w.Body.String(), "test | 1")
	})
}

func (suite *TapSyncTestSuite) TestHTMLQueries() {
	t := suite.T()
	t.Run("TestHTMLQuerySuccess", func(t *testing.T) {
		testhelpers.PopulateDb(suite.DB)
		defer testhelpers.ClearDataFromTable(suite.DB)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/sync", strings.NewReader("LANG=PSQL&&FORMAT=html&&QUERY=SELECT * FROM test"))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		suite.Service.Router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "text/html", w.Header().Get("Content-Type"))
		assert.Contains(t, w.Body.String(), "<th>id</th>")
		assert.Contains(t, w.Body.String(), "<th>name</th>")
		assert.Contains(t, w.Body.String(), "<th>number</th>")
		assert.Contains(t, w.Body.String(), "<td>test</td>")
		assert.Contains(t, w.Body.String(), "<td>1</td>")
	})
}
