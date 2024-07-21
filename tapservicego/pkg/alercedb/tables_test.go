package alercedb

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/stretchr/testify/assert"
)

func TestCreateTables(t *testing.T) {
	restoreDatabase()
	db, err := GetDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = CreateTables(db)
	assert.Nil(t, err)
	assertTableExists("object", db, t)
	assertTableExists("detection", db, t)
	assertTableExists("non_detection", db, t)
}

func TestCreateObjectsTable(t *testing.T) {
	restoreDatabase()
	db, err := GetDB(os.Getenv("DATABASE_URL"))
	assert.Nil(t, err)
	defer db.Close()
	err = createObjectTable(db)
	assert.Nil(t, err)
	assertTableExists("object", db, t)
}

func TestCreateObjectsTableIndex(t *testing.T) {
	restoreDatabase()
	db, err := GetDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = createObjectTable(db)
	assert.Nil(t, err)
	query := `SELECT COUNT(*) FROM pg_indexes WHERE tablename = 'object'`
	var count int
	err = db.QueryRow(query).Scan(&count)
	assert.Nil(t, err)
	assert.Equal(t, 5, count)
}

func TestCreateDetectionsTable(t *testing.T) {
	restoreDatabase()
	db, err := GetDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = createObjectTable(db)
	assert.Nil(t, err)
	err = createDetectionsTable(db)
	assert.Nil(t, err)
	assertTableExists("detection", db, t)
}

func TestCreateDetectionsTableIndex(t *testing.T) {
	restoreDatabase()
	db, err := GetDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = createObjectTable(db)
	assert.Nil(t, err)
	err = createDetectionsTable(db)
	assert.Nil(t, err)
	query := `SELECT COUNT(*) FROM pg_indexes WHERE tablename = 'detection'`
	var count int
	err = db.QueryRow(query).Scan(&count)
	assert.Nil(t, err)
	assert.Equal(t, 2, count)
}

func TestCreateNonDetectionsTable(t *testing.T) {
	restoreDatabase()
	db, err := GetDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = createObjectTable(db)
	assert.Nil(t, err)
	err = createNonDetectionsTable(db)
	assert.Nil(t, err)
	assertTableExists("non_detection", db, t)
}

func TestCreateNonDetectionsTableIndex(t *testing.T) {
	restoreDatabase()
	db, err := GetDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = createObjectTable(db)
	assert.Nil(t, err)
	err = createNonDetectionsTable(db)
	assert.Nil(t, err)
	query := `SELECT COUNT(*) FROM pg_indexes WHERE tablename = 'non_detection'`
	var count int
	err = db.QueryRow(query).Scan(&count)
	assert.Nil(t, err)
	assert.Equal(t, 2, count)
}

func TestCreateForcedPhotometryTable(t *testing.T) {
	restoreDatabase()
	db, err := GetDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = createObjectTable(db)
	assert.Nil(t, err)
	err = createForcedPhotometryTable(db)
	assert.Nil(t, err)
	assertTableExists("forced_photometry", db, t)
}
func TestCreateForcedPhotometryTableIndex(t *testing.T) {
	restoreDatabase()
	db, err := GetDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = createObjectTable(db)
	assert.Nil(t, err)
	err = createForcedPhotometryTable(db)
	assert.Nil(t, err)
	query := `SELECT COUNT(*) FROM pg_indexes WHERE tablename = 'forced_photometry'`
	var count int
	err = db.QueryRow(query).Scan(&count)
	assert.Nil(t, err)
	assert.Equal(t, 2, count)
}

func TestCreateFeatureTable(t *testing.T) {
	restoreDatabase()
	db, err := GetDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = createObjectTable(db)
	assert.Nil(t, err)
	err = createFeaturesTable(db)
	assert.Nil(t, err)
	assertTableExists("feature", db, t)
}

func TestCreateFeatureTableIndex(t *testing.T) {
	restoreDatabase()
	db, err := GetDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = createObjectTable(db)
	assert.Nil(t, err)
	err = createFeaturesTable(db)
	assert.Nil(t, err)
	query := `SELECT COUNT(*) FROM pg_indexes WHERE tablename = 'feature'`
	var count int
	err = db.QueryRow(query).Scan(&count)
	assert.Nil(t, err)
	assert.Equal(t, 2, count)
}

func TestCreateProbabilityTable(t *testing.T) {
	restoreDatabase()
	db, err := GetDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = createObjectTable(db)
	assert.Nil(t, err)
	err = createProbabilitiesTable(db)
	assert.Nil(t, err)
	assertTableExists("probability", db, t)
}

func TestCreateProbabilityTableIndex(t *testing.T) {
	restoreDatabase()
	db, err := GetDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = createObjectTable(db)
	assert.Nil(t, err)
	err = createProbabilitiesTable(db)
	assert.Nil(t, err)
	query := `SELECT COUNT(*) FROM pg_indexes WHERE tablename = 'probability'`
	var count int
	err = db.QueryRow(query).Scan(&count)
	assert.Nil(t, err)
	assert.Equal(t, 5, count)
}

func assertTableExists(tableName string, db *sql.DB, t *testing.T) {
	query := `SELECT COUNT(*) FROM information_schema.tables WHERE table_name = $1;`
	var count int
	err := db.QueryRow(query, tableName).Scan(&count)
	assert.Nil(t, err)
	assert.Equal(t, 1, count)
}
