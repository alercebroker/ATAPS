package tapsync

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLangSuccess(t *testing.T) {
	service := TapSyncService()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/sync", strings.NewReader("LANG=PSQL&&QUERY=SELECT * FROM table"))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	service.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestLangFailure(t *testing.T) {
	service := TapSyncService()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/sync", strings.NewReader(""))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	service.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestFormatSuccess(t *testing.T) {
	service := TapSyncService()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/sync", strings.NewReader("LANG=PSQL&&FORMAT=VOTable&&QUERY=SELECT * FROM table"))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	service.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestResponseFormatSuccess(t *testing.T) {
	service := TapSyncService()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/sync", strings.NewReader("LANG=PSQL&&RESPONSEFORMAT=VOTable&&QUERY=SELECT * FROM table"))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	service.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestFormatSuccessWithoutSpecifying(t *testing.T) {
	service := TapSyncService()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/sync", strings.NewReader("LANG=PSQL&&QUERY=SELECT * FROM table"))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	service.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestFormatFailureWhenProvidingBoth(t *testing.T) {
	service := TapSyncService()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/sync", strings.NewReader("LANG=PSQL&&FORMAT=VOTable&&RESPONSEFORMAT=VOTable"))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	service.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestFormatFailureWhenProvidingUnknownFormat(t *testing.T) {
	service := TapSyncService()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/sync", strings.NewReader("LANG=PSQL&&FORMAT=Unknown"))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	service.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
func TestFormatFailureWhenProvidingUnknownResponseFormat(t *testing.T) {
	service := TapSyncService()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/sync", strings.NewReader("LANG=PSQL&&RESPONSEFORMAT=Unknown"))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	service.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
func TestBadRequestIfQueryIsEmpty(t *testing.T) {
	service := TapSyncService()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/sync", strings.NewReader("LANG=PSQL"))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	service.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
