package tapsync

import (
	"ataps/internal/testhelpers"
	"ataps/pkg/alercedb"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/net/html"
)

func TestHtml_Object(t *testing.T) {
	db, err := GetDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		t.Fatal(err)
	}
	testhelpers.ClearALeRCEDB(db)
	err = testhelpers.PopulateALeRCEDB(db)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	service := NewTapSyncService()
	w := sendTestQuery("LANG=PSQL&&FORMAT=html&&QUERY=SELECT * FROM object LIMIT 3", service)
	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, "text/html", w.Header().Get("Content-Type"))
	var data []string
	var headers []string
	doc, err := html.Parse(w.Body)
	require.NoError(t, err)
	parseHTMLTable(doc, &headers, "th")
	parseHTMLTable(doc, &data, "td")
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

func TestHtml_NonExistentTable(t *testing.T) {
	service := NewTapSyncService()
	w := sendTestQuery("LANG=PSQL&&FORMAT=html&&QUERY=SELECT * FROM non_existent_table LIMIT 3", service)
	require.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestHtml_Detection(t *testing.T) {
	db, err := GetDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		t.Fatal(err)
	}
	testhelpers.ClearALeRCEDB(db)
	err = testhelpers.PopulateALeRCEDB(db)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	service := NewTapSyncService()
	w := sendTestQuery("LANG=PSQL&&FORMAT=html&&QUERY=SELECT * FROM detection LIMIT 3", service)
	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, "text/html", w.Header().Get("Content-Type"))
	var data []string
	var headers []string
	doc, err := html.Parse(w.Body)
	require.NoError(t, err)
	parseHTMLTable(doc, &headers, "th")
	parseHTMLTable(doc, &data, "td")
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

func TestHtml_NonDetection(t *testing.T) {
	db, err := GetDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		t.Fatal(err)
	}
	testhelpers.ClearALeRCEDB(db)
	err = testhelpers.PopulateALeRCEDB(db)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	service := NewTapSyncService()
	w := sendTestQuery("LANG=PSQL&&FORMAT=html&&QUERY=SELECT * FROM non_detection LIMIT 3", service)
	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, "text/html", w.Header().Get("Content-Type"))
	var data []string
	var headers []string
	doc, err := html.Parse(w.Body)
	require.NoError(t, err)
	parseHTMLTable(doc, &headers, "th")
	parseHTMLTable(doc, &data, "td")
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

func TestHtml_ForcedPhotometry(t *testing.T) {
	db, err := GetDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		t.Fatal(err)
	}
	testhelpers.ClearALeRCEDB(db)
	err = testhelpers.PopulateALeRCEDB(db)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	service := NewTapSyncService()
	w := sendTestQuery("LANG=PSQL&&FORMAT=html&&QUERY=SELECT * FROM forced_photometry LIMIT 3", service)
	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, "text/html", w.Header().Get("Content-Type"))
	var data []string
	var headers []string
	doc, err := html.Parse(w.Body)
	require.NoError(t, err)
	parseHTMLTable(doc, &headers, "th")
	parseHTMLTable(doc, &data, "td")
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

func TestHtml_Features(t *testing.T) {
	db, err := GetDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		t.Fatal(err)
	}
	testhelpers.ClearALeRCEDB(db)
	err = testhelpers.PopulateALeRCEDB(db)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	service := NewTapSyncService()
	w := sendTestQuery("LANG=PSQL&&FORMAT=html&&QUERY=SELECT * FROM feature LIMIT 3", service)
	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, "text/html", w.Header().Get("Content-Type"))
	var data []string
	var headers []string
	doc, err := html.Parse(w.Body)
	require.NoError(t, err)
	parseHTMLTable(doc, &headers, "th")
	parseHTMLTable(doc, &data, "td")
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

func TestHtml_Probabilities(t *testing.T) {
	db, err := GetDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		t.Fatal(err)
	}
	testhelpers.ClearALeRCEDB(db)
	err = testhelpers.PopulateALeRCEDB(db)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	service := NewTapSyncService()
	w := sendTestQuery("LANG=PSQL&&FORMAT=html&&QUERY=SELECT * FROM probability LIMIT 3", service)
	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, "text/html", w.Header().Get("Content-Type"))
	var data []string
	var headers []string
	doc, err := html.Parse(w.Body)
	require.NoError(t, err)
	parseHTMLTable(doc, &headers, "th")
	parseHTMLTable(doc, &data, "td")
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
