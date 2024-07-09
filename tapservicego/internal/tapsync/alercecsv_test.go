package tapsync

import (
	"ataps/internal/testhelpers"
	"ataps/pkg/alercedb"
	"encoding/csv"
	"fmt"
	"net/http"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCsv(t *testing.T) {
	test_get_csv_data_from_object := func(t *testing.T, table string, ensureObjectExistIn *string, overrideOid *string) [][]string {
		testhelpers.ClearALeRCEDB()
		db := populateAlerceDB()
		defer db.Close()
		var oid string
		if overrideOid == nil {
			if ensureObjectExistIn == nil {
				row := db.QueryRow("SELECT oid FROM object LIMIT 1")
				if row.Err() != nil {
					t.Fatal(row.Err())
				}
				row.Scan(&oid)
			} else {
				row := db.QueryRow(fmt.Sprintf("SELECT oid FROM %s LIMIT 1", *ensureObjectExistIn))
				if row.Err() != nil {
					t.Fatal(row.Err())
				}
				row.Scan(&oid)
			}
		} else {
			oid = *overrideOid
		}
		service := NewTapSyncService()
		response := sendTestQuery(fmt.Sprintf("LANG=PSQL&&FORMAT=csv&&QUERY=SELECT * FROM %s WHERE oid='%s'", table, oid), service)
		require.Equal(t, http.StatusOK, response.Code)
		reader := csv.NewReader(response.Body)
		records, err := reader.ReadAll()
		if err != nil {
			t.Fatal(err)
		}
		return records
	}

	t.Run("TestCsv_Detections", func(t *testing.T) {
		testhelpers.ClearALeRCEDB()
		db := populateAlerceDB()
		defer db.Close()
		service := NewTapSyncService()
		w := sendTestQuery("LANG=PSQL&&FORMAT=csv&&QUERY=SELECT * FROM detection LIMIT 1", service)
		require.Equal(t, http.StatusOK, w.Code)
		require.Equal(t, "text/csv", w.Header().Get("Content-Type"))
		records, err := csv.NewReader(w.Body).ReadAll()
		if err != nil {
			t.Fatal(err)
		}
		columnNames := getColumnNames(alercedb.Detection{})
		sort.Strings(columnNames)
		require.Equal(t, columnNames, records[0])

	})

	t.Run("TestCsv_DetectionsFromAnObject", func(t *testing.T) {
		ensureObjectExistIn := "detection"
		records := test_get_csv_data_from_object(t, "detection", &ensureObjectExistIn, nil)
		require.Greater(t, len(records), 1)
	})

	t.Run("TestCsv_DetectionsFromUnknownObject", func(t *testing.T) {
		var oid = "unknown"
		records := test_get_csv_data_from_object(t, "detection", nil, &oid)
		require.Equal(t, 0, len(records))
	})

	t.Run("TestCsv_NonDetection", func(t *testing.T) {
		testhelpers.ClearALeRCEDB()
		db := populateAlerceDB()
		defer db.Close()
		service := NewTapSyncService()
		w := sendTestQuery("LANG=PSQL&&FORMAT=csv&&QUERY=SELECT * FROM non_detection LIMIT 1", service)
		require.Equal(t, http.StatusOK, w.Code)
		require.Equal(t, "text/csv", w.Header().Get("Content-Type"))
		columnNames := getColumnNames(alercedb.NonDetection{})
		records, err := csv.NewReader(w.Body).ReadAll()
		if err != nil {
			t.Fatal(err)
		}
		sort.Strings(columnNames)
		require.Equal(t, columnNames, records[0])
	})

	t.Run("TestCsv_NonDetectionFromAnObject", func(t *testing.T) {
		oidFrom := "non_detection"
		records := test_get_csv_data_from_object(t, "non_detection", &oidFrom, nil)
		require.Greater(t, len(records), 1)
	})

	t.Run("TestCsv_NonDetectionFromUnknownObject", func(t *testing.T) {
		var oid = "unknown"
		records := test_get_csv_data_from_object(t, "non_detection", nil, &oid)
		require.Equal(t, 0, len(records))
	})

	t.Run("TestCsv_ForcedPhotometry", func(t *testing.T) {
		testhelpers.ClearALeRCEDB()
		db := populateAlerceDB()
		defer db.Close()
		service := NewTapSyncService()
		w := sendTestQuery("LANG=PSQL&&FORMAT=csv&&QUERY=SELECT * FROM forced_photometry LIMIT 1", service)
		require.Equal(t, http.StatusOK, w.Code)
		require.Equal(t, "text/csv", w.Header().Get("Content-Type"))
		columnNames := getColumnNames(alercedb.ForcedPhotometry{})
		records, err := csv.NewReader(w.Body).ReadAll()
		if err != nil {
			t.Fatal(err)
		}
		sort.Strings(columnNames)
		require.Equal(t, columnNames, records[0])
	})

	t.Run("TestCsv_ForcedPhotometryFromAnObject", func(t *testing.T) {
		oidFrom := "forced_photometry"
		records := test_get_csv_data_from_object(t, "forced_photometry", &oidFrom, nil)
		require.Greater(t, len(records), 1)
	})

	t.Run("TestCsv_ForcedPhotometryFromUnknownObject", func(t *testing.T) {
		var oid = "unknown"
		records := test_get_csv_data_from_object(t, "forced_photometry", nil, &oid)
		require.Equal(t, 0, len(records))
	})

	t.Run("TestCsv_Features", func(t *testing.T) {
		testhelpers.ClearALeRCEDB()
		db := populateAlerceDB()
		defer db.Close()
		service := NewTapSyncService()
		w := sendTestQuery("LANG=PSQL&&FORMAT=csv&&QUERY=SELECT * FROM feature LIMIT 1", service)
		require.Equal(t, http.StatusOK, w.Code)
		require.Equal(t, "text/csv", w.Header().Get("Content-Type"))
		columnNames := getColumnNames(alercedb.Feature{})
		records, err := csv.NewReader(w.Body).ReadAll()
		if err != nil {
			t.Fatal(err)
		}
		sort.Strings(columnNames)
		require.Equal(t, columnNames, records[0])
	})

	t.Run("TestCsv_FeaturesFromAnObject", func(t *testing.T) {
		oidFrom := "feature"
		records := test_get_csv_data_from_object(t, "feature", &oidFrom, nil)
		require.Greater(t, len(records), 1)
	})

	t.Run("TestCsv_FeaturesFromUnknownObject", func(t *testing.T) {
		var oid = "unknown"
		records := test_get_csv_data_from_object(t, "feature", nil, &oid)
		require.Equal(t, 0, len(records))
	})

	t.Run("TestCsv_Objects", func(t *testing.T) {
		testhelpers.ClearALeRCEDB()
		db := populateAlerceDB()
		defer db.Close()
		service := NewTapSyncService()
		w := sendTestQuery("LANG=PSQL&&FORMAT=csv&&QUERY=SELECT * FROM object LIMIT 2", service)
		require.Equal(t, http.StatusOK, w.Code)
		require.Equal(t, "text/csv", w.Header().Get("Content-Type"))
		columnNames := getColumnNames(alercedb.Object{})
		records, err := csv.NewReader(w.Body).ReadAll()
		if err != nil {
			t.Fatal(err)
		}
		sort.Strings(columnNames)
		require.Equal(t, columnNames, records[0])
		require.Equal(t, 3, len(records)) // header + 2 objects
	})

	t.Run("TestCsv_Probabilities", func(t *testing.T) {
		testhelpers.ClearALeRCEDB()
		db := populateAlerceDB()
		defer db.Close()
		service := NewTapSyncService()
		w := sendTestQuery("LANG=PSQL&&FORMAT=csv&&QUERY=SELECT * FROM probability LIMIT 1", service)
		require.Equal(t, http.StatusOK, w.Code)
		require.Equal(t, "text/csv", w.Header().Get("Content-Type"))
		columnNames := getColumnNames(alercedb.Probability{})
		records, err := csv.NewReader(w.Body).ReadAll()
		if err != nil {
			t.Fatal(err)
		}
		sort.Strings(columnNames)
		require.Equal(t, columnNames, records[0])
	})

	t.Run("TestCsv_ProbabilitiesFromAnObject", func(t *testing.T) {
		oidFrom := "probability"
		records := test_get_csv_data_from_object(t, "probability", &oidFrom, nil)
		require.Greater(t, len(records), 1)
	})
}
