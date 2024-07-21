package tapsync

import (
	"ataps/internal/testhelpers"
	"ataps/pkg/alercedb"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestText_Object(t *testing.T) {
	testhelpers.ClearALeRCEDB()
	db := populateAlerceDB()
	defer db.Close()
	service := NewTapSyncService()
	w := sendTestQuery("LANG=PSQL&&FORMAT=text&&QUERY=SELECT * FROM object LIMIT 3", service)
	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, "text/plain", w.Header().Get("Content-Type"))
	var data []string
	var headers []string
	err := parseTextTable(w.Body.String(), &data, &headers)
	require.NoError(t, err)
	var rows []map[string]string
	for i := 0; i < len(data); i += len(headers) {
		row := make(map[string]string)
		for j, header := range headers {
			row[header] = data[i+j]
		}
		rows = append(rows, row)
	}
	require.Len(t, rows, 3)
	columnNames := getColumnNames(alercedb.Object{})
	require.ElementsMatch(t, headers, columnNames)
}

func TestText_Detection(t *testing.T) {
	testhelpers.ClearALeRCEDB()
	db := populateAlerceDB()
	defer db.Close()
	service := NewTapSyncService()
	w := sendTestQuery("LANG=PSQL&&FORMAT=text&&QUERY=SELECT * FROM detection LIMIT 3", service)
	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, "text/plain", w.Header().Get("Content-Type"))
	var data []string
	var headers []string
	err := parseTextTable(w.Body.String(), &data, &headers)
	require.NoError(t, err)
	var rows []map[string]string
	for i := 0; i < len(data); i += len(headers) {
		row := make(map[string]string)
		for j, header := range headers {
			row[header] = data[i+j]
		}
		rows = append(rows, row)
	}
	require.Len(t, rows, 3)
	columnNames := getColumnNames(alercedb.Detection{})
	require.ElementsMatch(t, headers, columnNames)
}

func TestText_NonDetection(t *testing.T) {
	testhelpers.ClearALeRCEDB()
	db := populateAlerceDB()
	defer db.Close()
	service := NewTapSyncService()
	w := sendTestQuery("LANG=PSQL&&FORMAT=text&&QUERY=SELECT * FROM non_detection LIMIT 3", service)
	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, "text/plain", w.Header().Get("Content-Type"))
	var data []string
	var headers []string
	err := parseTextTable(w.Body.String(), &data, &headers)
	require.NoError(t, err)
	var rows []map[string]string
	for i := 0; i < len(data); i += len(headers) {
		row := make(map[string]string)
		for j, header := range headers {
			row[header] = data[i+j]
		}
		rows = append(rows, row)
	}
	require.Len(t, rows, 3)
	columnNames := getColumnNames(alercedb.NonDetection{})
	require.ElementsMatch(t, headers, columnNames)
}

func TestText_ForcedPhotometry(t *testing.T) {
	testhelpers.ClearALeRCEDB()
	db := populateAlerceDB()
	defer db.Close()
	service := NewTapSyncService()
	w := sendTestQuery("LANG=PSQL&&FORMAT=text&&QUERY=SELECT * FROM forced_photometry LIMIT 3", service)
	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, "text/plain", w.Header().Get("Content-Type"))
	var data []string
	var headers []string
	err := parseTextTable(w.Body.String(), &data, &headers)
	require.NoError(t, err)
	var rows []map[string]string
	for i := 0; i < len(data); i += len(headers) {
		row := make(map[string]string)
		for j, header := range headers {
			row[header] = data[i+j]
		}
		rows = append(rows, row)
	}
	require.Len(t, rows, 3)
	columnNames := getColumnNames(alercedb.ForcedPhotometry{})
	require.ElementsMatch(t, headers, columnNames)
}

func TestText_Features(t *testing.T) {
	testhelpers.ClearALeRCEDB()
	db := populateAlerceDB()
	defer db.Close()
	service := NewTapSyncService()
	w := sendTestQuery("LANG=PSQL&&FORMAT=text&&QUERY=SELECT * FROM feature LIMIT 3", service)
	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, "text/plain", w.Header().Get("Content-Type"))
	var data []string
	var headers []string
	err := parseTextTable(w.Body.String(), &data, &headers)
	require.NoError(t, err)
	var rows []map[string]string
	for i := 0; i < len(data); i += len(headers) {
		row := make(map[string]string)
		for j, header := range headers {
			row[header] = data[i+j]
		}
		rows = append(rows, row)
	}
	require.Len(t, rows, 3)
	columnNames := getColumnNames(alercedb.Feature{})
	require.ElementsMatch(t, headers, columnNames)
}

func TestText_Probabilities(t *testing.T) {
	testhelpers.ClearALeRCEDB()
	db := populateAlerceDB()
	defer db.Close()
	service := NewTapSyncService()
	w := sendTestQuery("LANG=PSQL&&FORMAT=text&&QUERY=SELECT * FROM probability LIMIT 3", service)
	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, "text/plain", w.Header().Get("Content-Type"))
	var data []string
	var headers []string
	err := parseTextTable(w.Body.String(), &data, &headers)
	require.NoError(t, err)
	var rows []map[string]string
	for i := 0; i < len(data); i += len(headers) {
		row := make(map[string]string)
		for j, header := range headers {
			row[header] = data[i+j]
		}
		rows = append(rows, row)
	}
	require.Len(t, rows, 3)
	columnNames := getColumnNames(alercedb.Probability{})
	require.ElementsMatch(t, headers, columnNames)
}
