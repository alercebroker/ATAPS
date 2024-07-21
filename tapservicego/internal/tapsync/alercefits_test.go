package tapsync

import (
	"ataps/internal/testhelpers"
	"ataps/pkg/alercedb"
	"net/http"
	"testing"

	"github.com/dirodriguezm/fitsio"
	"github.com/stretchr/testify/require"
)

func TestFits_Object(t *testing.T) {
	testhelpers.ClearALeRCEDB()
	db := populateAlerceDB()
	defer db.Close()
	service := NewTapSyncService()
	w := sendTestQuery("LANG=PSQL&&FORMAT=fits&&QUERY=SELECT * FROM object LIMIT 3", service)
	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, "application/fits", w.Header().Get("Content-Type"))
	f, err := fitsio.Open(w.Body)
	require.Nil(t, err)
	defer f.Close()
	hdu := f.HDU(1)
	table := hdu.(*fitsio.Table)
	require.True(t, table.Type() == fitsio.BINARY_TBL)
	require.Equal(t, "results", table.Name())
	require.Equal(t, 10, table.NumCols())
	require.Equal(t, int64(3), table.NumRows())
	rows, err := table.Read(0, table.NumRows())
	require.Nil(t, err)
	defer rows.Close()
	count, parsedData := parseResponseData(t, rows)
	columnNames := getColumnNames(alercedb.Object{})
	assertColumnsExist(t, columnNames, parsedData)
	require.Equal(t, 3, count)
}

func TestFits_Detection(t *testing.T) {
	testhelpers.ClearALeRCEDB()
	db := populateAlerceDB()
	defer db.Close()
	service := NewTapSyncService()
	w := sendTestQuery("LANG=PSQL&&FORMAT=fits&&QUERY=SELECT * FROM detection LIMIT 3", service)
	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, "application/fits", w.Header().Get("Content-Type"))
	f, err := fitsio.Open(w.Body)
	require.Nil(t, err)
	defer f.Close()
	hdu := f.HDU(1)
	table := hdu.(*fitsio.Table)
	require.True(t, table.Type() == fitsio.BINARY_TBL)
	require.Equal(t, "results", table.Name())
	require.Equal(t, 19, table.NumCols())
	require.Equal(t, int64(3), table.NumRows())
	rows, err := table.Read(0, table.NumRows())
	require.Nil(t, err)
	defer rows.Close()
	count, parsedData := parseResponseData(t, rows)
	columnNames := getColumnNames(alercedb.Detection{})
	assertColumnsExist(t, columnNames, parsedData)
	require.Equal(t, 3, count)
}

func TestFits_NonDetection(t *testing.T) {
	testhelpers.ClearALeRCEDB()
	db := populateAlerceDB()
	defer db.Close()
	service := NewTapSyncService()
	w := sendTestQuery("LANG=PSQL&&FORMAT=fits&&QUERY=SELECT * FROM non_detection LIMIT 3", service)
	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, "application/fits", w.Header().Get("Content-Type"))
	f, err := fitsio.Open(w.Body)
	require.Nil(t, err)
	defer f.Close()
	hdu := f.HDU(1)
	table := hdu.(*fitsio.Table)
	require.True(t, table.Type() == fitsio.BINARY_TBL)
	require.Equal(t, "results", table.Name())
	require.Equal(t, 4, table.NumCols())
	require.Equal(t, int64(3), table.NumRows())
	rows, err := table.Read(0, table.NumRows())
	require.Nil(t, err)
	defer rows.Close()
	count, parsedData := parseResponseData(t, rows)
	columnNames := getColumnNames(alercedb.NonDetection{})
	assertColumnsExist(t, columnNames, parsedData)
	require.Equal(t, 3, count)
}

func TestFits_ForcedPhotometry(t *testing.T) {
	testhelpers.ClearALeRCEDB()
	db := populateAlerceDB()
	defer db.Close()
	service := NewTapSyncService()
	w := sendTestQuery("LANG=PSQL&&FORMAT=fits&&QUERY=SELECT * FROM forced_photometry LIMIT 3", service)
	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, "application/fits", w.Header().Get("Content-Type"))
	f, err := fitsio.Open(w.Body)
	require.Nil(t, err)
	defer f.Close()
	hdu := f.HDU(1)
	table := hdu.(*fitsio.Table)
	require.True(t, table.Type() == fitsio.BINARY_TBL)
	require.Equal(t, "results", table.Name())
	require.Equal(t, 19, table.NumCols())
	require.Equal(t, int64(3), table.NumRows())
	rows, err := table.Read(0, table.NumRows())
	require.Nil(t, err)
	defer rows.Close()
	count, parsedData := parseResponseData(t, rows)
	columnNames := getColumnNames(alercedb.ForcedPhotometry{})
	assertColumnsExist(t, columnNames, parsedData)
	require.Equal(t, 3, count)
}

func TestFits_Features(t *testing.T) {
	testhelpers.ClearALeRCEDB()
	db := populateAlerceDB()
	defer db.Close()
	service := NewTapSyncService()
	w := sendTestQuery("LANG=PSQL&&FORMAT=fits&&QUERY=SELECT * FROM feature LIMIT 3", service)
	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, "application/fits", w.Header().Get("Content-Type"))
	f, err := fitsio.Open(w.Body)
	require.Nil(t, err)
	defer f.Close()
	hdu := f.HDU(1)
	table := hdu.(*fitsio.Table)
	require.True(t, table.Type() == fitsio.BINARY_TBL)
	require.Equal(t, "results", table.Name())
	require.Equal(t, 5, table.NumCols())
	require.Equal(t, int64(3), table.NumRows())
	rows, err := table.Read(0, table.NumRows())
	require.Nil(t, err)
	defer rows.Close()
	count, parsedData := parseResponseData(t, rows)
	columnNames := getColumnNames(alercedb.Feature{})
	assertColumnsExist(t, columnNames, parsedData)
	require.Equal(t, 3, count)
}

func TestFits_Probabilities(t *testing.T) {
	testhelpers.ClearALeRCEDB()
	db := populateAlerceDB()
	defer db.Close()
	service := NewTapSyncService()
	w := sendTestQuery("LANG=PSQL&&FORMAT=fits&&QUERY=SELECT * FROM probability LIMIT 3", service)
	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, "application/fits", w.Header().Get("Content-Type"))
	f, err := fitsio.Open(w.Body)
	require.Nil(t, err)
	defer f.Close()
	hdu := f.HDU(1)
	table := hdu.(*fitsio.Table)
	require.True(t, table.Type() == fitsio.BINARY_TBL)
	require.Equal(t, "results", table.Name())
	require.Equal(t, 6, table.NumCols())
	require.Equal(t, int64(3), table.NumRows())
	rows, err := table.Read(0, table.NumRows())
	require.Nil(t, err)
	defer rows.Close()
	count, parsedData := parseResponseData(t, rows)
	columnNames := getColumnNames(alercedb.Probability{})
	assertColumnsExist(t, columnNames, parsedData)
	require.Equal(t, 3, count)
}

func parseResponseData(t *testing.T, rows *fitsio.Rows) (int, []map[string]interface{}) {
	parsedData := []map[string]interface{}{}
	count := 0
	for rows.Next() {
		row := map[string]interface{}{}
		err := rows.Scan(&row)
		require.Nil(t, err)
		parsedData = append(parsedData, row)
		count = count + 1
	}
	return count, parsedData
}

func assertColumnsExist(t *testing.T, columnNames []string, parsedData []map[string]interface{}) {
	for _, data := range parsedData {
		for _, columnName := range columnNames {
			_, ok := data[columnName]
			require.True(t, ok)
		}
	}
}
