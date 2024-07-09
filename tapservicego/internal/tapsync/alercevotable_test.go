package tapsync

import (
	"ataps/internal/testhelpers"
	"ataps/pkg/alercedb"
	"ataps/pkg/votable"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVotable_Object(t *testing.T) {
	testhelpers.ClearALeRCEDB()
	db := populateAlerceDB()
	defer db.Close()
	service := NewTapSyncService()
	w := sendTestQuery("LANG=PSQL&&FORMAT=votable&&QUERY=SELECT * FROM object LIMIT 3", service)
	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, "application/x-votable+xml", w.Header().Get("Content-Type"))
	voTable, err := votable.NewVOTableFromString(w.Body.String())
	require.NoError(t, err)
	columnNames := getColumnNames(alercedb.Object{})
	require.Len(t, voTable.Resource.Tables[0].Fields, len(columnNames))
	for _, field := range voTable.Resource.Tables[0].Fields {
		require.Contains(t, columnNames, field.Name)
	}
	require.Len(t, voTable.Resource.Tables[0].Data.TableData.Rows, 3)
}

func TestVotable_Detection(t *testing.T) {
	testhelpers.ClearALeRCEDB()
	db := populateAlerceDB()
	defer db.Close()
	service := NewTapSyncService()
	w := sendTestQuery("LANG=PSQL&&FORMAT=votable&&QUERY=SELECT * FROM detection LIMIT 3", service)
	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, "application/x-votable+xml", w.Header().Get("Content-Type"))
	voTable, err := votable.NewVOTableFromString(w.Body.String())
	require.NoError(t, err)
	columnNames := getColumnNames(alercedb.Detection{})
	require.Len(t, voTable.Resource.Tables[0].Fields, len(columnNames))
	for _, field := range voTable.Resource.Tables[0].Fields {
		require.Contains(t, columnNames, field.Name)
	}
	require.Len(t, voTable.Resource.Tables[0].Data.TableData.Rows, 3)
}

func TestVotable_NonDetection(t *testing.T) {
	testhelpers.ClearALeRCEDB()
	db := populateAlerceDB()
	defer db.Close()
	service := NewTapSyncService()
	w := sendTestQuery("LANG=PSQL&&FORMAT=votable&&QUERY=SELECT * FROM non_detection LIMIT 3", service)
	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, "application/x-votable+xml", w.Header().Get("Content-Type"))
	voTable, err := votable.NewVOTableFromString(w.Body.String())
	require.NoError(t, err)
	columnNames := getColumnNames(alercedb.NonDetection{})
	require.Len(t, voTable.Resource.Tables[0].Fields, len(columnNames))
	for _, field := range voTable.Resource.Tables[0].Fields {
		require.Contains(t, columnNames, field.Name)
	}
	require.Len(t, voTable.Resource.Tables[0].Data.TableData.Rows, 3)
}

func TestVotable_ForcedPhotometry(t *testing.T) {
	testhelpers.ClearALeRCEDB()
	db := populateAlerceDB()
	defer db.Close()
	service := NewTapSyncService()
	w := sendTestQuery("LANG=PSQL&&FORMAT=votable&&QUERY=SELECT * FROM forced_photometry LIMIT 3", service)
	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, "application/x-votable+xml", w.Header().Get("Content-Type"))
	voTable, err := votable.NewVOTableFromString(w.Body.String())
	require.NoError(t, err)
	columnNames := getColumnNames(alercedb.ForcedPhotometry{})
	require.Len(t, voTable.Resource.Tables[0].Fields, len(columnNames))
	for _, field := range voTable.Resource.Tables[0].Fields {
		require.Contains(t, columnNames, field.Name)
	}
	require.Len(t, voTable.Resource.Tables[0].Data.TableData.Rows, 3)
}

func TestVotable_Features(t *testing.T) {
	testhelpers.ClearALeRCEDB()
	db := populateAlerceDB()
	defer db.Close()
	service := NewTapSyncService()
	w := sendTestQuery("LANG=PSQL&&FORMAT=votable&&QUERY=SELECT * FROM feature LIMIT 3", service)
	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, "application/x-votable+xml", w.Header().Get("Content-Type"))
	voTable, err := votable.NewVOTableFromString(w.Body.String())
	require.NoError(t, err)
	columnNames := getColumnNames(alercedb.Feature{})
	require.Len(t, voTable.Resource.Tables[0].Fields, len(columnNames))
	for _, field := range voTable.Resource.Tables[0].Fields {
		require.Contains(t, columnNames, field.Name)
	}
	require.Len(t, voTable.Resource.Tables[0].Data.TableData.Rows, 3)
}

func TestVotable_Probabilities(t *testing.T) {
	testhelpers.ClearALeRCEDB()
	db := populateAlerceDB()
	defer db.Close()
	service := NewTapSyncService()
	w := sendTestQuery("LANG=PSQL&&FORMAT=votable&&QUERY=SELECT * FROM probability LIMIT 3", service)
	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, "application/x-votable+xml", w.Header().Get("Content-Type"))
	voTable, err := votable.NewVOTableFromString(w.Body.String())
	require.NoError(t, err)
	columnNames := getColumnNames(alercedb.Probability{})
	require.Len(t, voTable.Resource.Tables[0].Fields, len(columnNames))
	for _, field := range voTable.Resource.Tables[0].Fields {
		require.Contains(t, columnNames, field.Name)
	}
	require.Len(t, voTable.Resource.Tables[0].Data.TableData.Rows, 3)
}
