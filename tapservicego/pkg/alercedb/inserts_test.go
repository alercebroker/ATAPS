package alercedb

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInsertSampleObjects(t *testing.T) {
	restoreDatabase()
	db, err := GetDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = CreateTables(db)
	require.Nil(t, err)
	_, err = InsertSampleObjects(db, 10)
	query := `SELECT COUNT(*) FROM object;`
	var count int
	err = db.QueryRow(query).Scan(&count)
	require.Nil(t, err)
	require.Equal(t, 10, count)
}

func TestInsertSampleDetections(t *testing.T) {
	restoreDatabase()
	db, err := GetDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = CreateTables(db)
	require.Nil(t, err)
	oidPool, err := InsertSampleObjects(db, 10)
	require.Nil(t, err)
	err = InsertSampleDetections(db, 100, oidPool)
	require.Nil(t, err)
	query := `SELECT COUNT(*) FROM detection;`
	var count int
	err = db.QueryRow(query).Scan(&count)
	require.Nil(t, err)
	require.Equal(t, 100, count)
}

func TestInsertSampleNonDetections(t *testing.T) {
	restoreDatabase()
	db, err := GetDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = CreateTables(db)
	require.Nil(t, err)
	oidPool, err := InsertSampleObjects(db, 10)
	require.Nil(t, err)
	err = InsertSampleNonDetections(db, 100, oidPool)
	require.Nil(t, err)
	query := `SELECT COUNT(*) FROM non_detection;`
	var count int
	err = db.QueryRow(query).Scan(&count)
	require.Nil(t, err)
	require.Equal(t, 100, count)
}

func TestInsertSampleForcedPhotometry(t *testing.T) {
	restoreDatabase()
	db, err := GetDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = CreateTables(db)
	require.Nil(t, err)
	oidPool, err := InsertSampleObjects(db, 10)
	require.Nil(t, err)
	err = InsertSampleForcedPhotometry(db, 100, oidPool)
	require.Nil(t, err)
	query := `SELECT COUNT(*) FROM forced_photometry;`
	var count int
	err = db.QueryRow(query).Scan(&count)
	require.Nil(t, err)
	require.Equal(t, 100, count)
}

func TestInsertSampleFeatures(t *testing.T) {
	restoreDatabase()
	db, err := GetDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = CreateTables(db)
	require.Nil(t, err)
	oidPool, err := InsertSampleObjects(db, 10)
	require.Nil(t, err)
	err = InsertSampleFeatures(db, 100, oidPool)
	require.Nil(t, err)
	query := `SELECT COUNT(*) FROM feature;`
	var count int
	err = db.QueryRow(query).Scan(&count)
	require.Nil(t, err)
	require.Equal(t, 100, count)
}

func TestInsertSampleProbabilities(t *testing.T) {
	restoreDatabase()
	db, err := GetDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = CreateTables(db)
	require.Nil(t, err)
	oidPool, err := InsertSampleObjects(db, 10)
	require.Nil(t, err)
	classPool := []string{"class1", "class2", "class3"}
	err = InsertSampleProbabilities(db, oidPool, classPool, "classifier")
	require.Nil(t, err)
	query := `SELECT COUNT(*) FROM probability;`
	var count int
	err = db.QueryRow(query).Scan(&count)
	require.Nil(t, err)
	require.Equal(t, len(classPool)*len(oidPool), count)
}
