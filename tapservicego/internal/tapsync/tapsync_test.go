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
		assert.Equal(t, "text/csv", w.Header().Get("Content-Type"))
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
