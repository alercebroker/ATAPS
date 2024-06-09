package tapsync

import (
	"ataps/internal/testhelpers"
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/dirodriguezm/fitsio"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

var container *postgres.PostgresContainer
var ctx context.Context

func globalSetup() {
	var err error
	ctx = context.Background()
	container, err = testhelpers.CreatePostgresContainer(ctx)
	if err != nil {
		log.Fatal(err)
	}
	var connStr string
	connStr, err = container.ConnectionString(ctx)
	if err != nil {
		log.Fatal(err)
	}
	os.Setenv("DATABASE_URL", connStr)
}

func globalTeardown() {
	testhelpers.CleanUpContainer(ctx, container)
}

func TestMain(m *testing.M) {
	globalSetup()
	code := m.Run()
	globalTeardown()
	os.Exit(code)
}

func TestQueryParams(t *testing.T) {
	t.Run("TestLangSuccess", func(t *testing.T) {
		service := NewTapSyncService()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/sync", strings.NewReader("LANG=PSQL&&QUERY=SELECT 'test'"))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		service.Router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("TestLangFailure", func(t *testing.T) {
		service := NewTapSyncService()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/sync", strings.NewReader(""))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		service.Router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "Invalid LANG ")
	})
	t.Run("TestFormatSuccess", func(t *testing.T) {
		service := NewTapSyncService()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/sync", strings.NewReader("LANG=PSQL&&FORMAT=votable&&QUERY=SELECT 'test'"))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		service.Router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("TestResponseFormatSuccess", func(t *testing.T) {
		service := NewTapSyncService()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/sync", strings.NewReader("LANG=PSQL&&RESPONSEFORMAT=votable&&QUERY=SELECT 'test'"))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		service.Router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("TestFormatSuccessWithoutSpecifying", func(t *testing.T) {
		service := NewTapSyncService()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/sync", strings.NewReader("LANG=PSQL&&QUERY=SELECT 'test'"))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		service.Router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("TestFormatFailureWhenProvidingBoth", func(t *testing.T) {
		service := NewTapSyncService()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/sync", strings.NewReader("LANG=PSQL&&FORMAT=votable&&RESPONSEFORMAT=votable"))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		service.Router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
	t.Run("TestFormatFailureWhenProvidingUnknownFormat", func(t *testing.T) {
		service := NewTapSyncService()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/sync", strings.NewReader("LANG=PSQL&&FORMAT=Unknown"))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		service.Router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
	t.Run("TestFormatFailureWhenProvidingUnknownResponseFormat", func(t *testing.T) {
		service := NewTapSyncService()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/sync", strings.NewReader("LANG=PSQL&&RESPONSEFORMAT=Unknown"))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		service.Router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
	t.Run("TestBadRequestIfQueryIsEmpty", func(t *testing.T) {
		service := NewTapSyncService()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/sync", strings.NewReader("LANG=PSQL"))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		service.Router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestCSVQueries(t *testing.T) {
	t.Run("TestCSVQuerySuccess", func(t *testing.T) {
		service := NewTapSyncService()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/sync", strings.NewReader("LANG=PSQL&&FORMAT=csv&&QUERY=SELECT 'test'"))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		service.Router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "?column?\ntest\n", w.Body.String())
		assert.Equal(t, "text/csv", w.Header().Get("Content-Type"))
	})
	t.Run("TestTSVQuerySuccess", func(t *testing.T) {
		service := NewTapSyncService()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/sync", strings.NewReader("LANG=PSQL&&FORMAT=tsv&&QUERY=SELECT 'test'"))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		service.Router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "?column?\ntest\n", w.Body.String())
		assert.Equal(t, "text/tab-separated-values", w.Header().Get("Content-Type"))
	})
	t.Run("TestCSVQueryFailure", func(t *testing.T) {
		service := NewTapSyncService()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/sync", strings.NewReader("LANG=PSQL&&FORMAT=csv&&QUERY=SELECT * from dontexist"))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		service.Router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Equal(t, "application/xml; charset=utf-8", w.Header().Get("Content-Type"))
		// the default gin xml render does not show quotes
		assert.Contains(t, w.Body.String(), "relation &#34;dontexist&#34; does not exist")
	})
}

func TestVOTableQueries(t *testing.T) {
	t.Run("TestVOTableQuerySuccess", func(t *testing.T) {
		service := NewTapSyncService()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/sync", strings.NewReader("LANG=PSQL&&FORMAT=votable&&QUERY=SELECT 'test'"))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		service.Router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, `<?xml version="1.0" encoding="UTF-8"?>
<VOTABLE version="1.4" xmlns="http://www.ivoa.net/xml/VOTable/v1.4">
	<RESOURCE type="results">
		<INFO name="QUERY_STATUS" value="OK"></INFO>
		<TABLE name="results">
			<DESCRIPTION>Results of the query</DESCRIPTION>
			<FIELD name="?column?" datatype=""></FIELD>
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

func TestFitsQueries(t *testing.T) {
	t.Run("TestFitsQuerySuccess", func(t *testing.T) {
		db, err := GetDB(os.Getenv("DATABASE_URL"))
		if err != nil {
			t.Fatal(err)
		}
		testhelpers.PopulateDb(db)
		defer testhelpers.ClearDataFromTable(db)
		service := NewTapSyncService()
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/sync", strings.NewReader("LANG=PSQL&&FORMAT=fits&&QUERY=SELECT * FROM test"))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		service.Router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "application/fits", w.Header().Get("Content-Type"))
		// read the fits file and parse it
		f, err := fitsio.Open(w.Body)
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()
		hdu := f.HDU(1)
		table := hdu.(*fitsio.Table)
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
			t.Logf("Row: %d, %v, %d", id, name, number)
			assert.Equal(t, "test", name)
			assert.Equal(t, 1, number)
			count++
		}
		assert.Equal(t, 1, count)
	})
}
