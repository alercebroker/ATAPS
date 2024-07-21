package testhelpers

import (
	"ataps/pkg/alercedb"
	"database/sql"
	"os"
)

func GetDB(url string) (*sql.DB, error) {
	db, err := sql.Open("pgx", url)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func PopulateDb(db *sql.DB) error {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS test (id SERIAL PRIMARY KEY, name TEXT, number INT)")
	if err != nil {
		return err
	}
	_, err = db.Exec("INSERT INTO test (name, number) VALUES ('test', 1)")
	if err != nil {
		return err
	}
	return nil
}

func ClearDataFromTable(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM test")
	if err != nil {
		return err
	}
	return nil
}

func PopulateALeRCEDB(db *sql.DB) error {
	err := alercedb.CreateTables(db)
	if err != nil {
		return err
	}
	oidPool, err := alercedb.InsertSampleObjects(db, 100)
	if err != nil {
		return err
	}
	err = alercedb.InsertSampleDetections(db, 1000, oidPool)
	if err != nil {
		return err
	}
	err = alercedb.InsertSampleNonDetections(db, 1000, oidPool)
	if err != nil {
		return err
	}
	err = alercedb.InsertSampleForcedPhotometry(db, 1000, oidPool)
	if err != nil {
		return err
	}
	err = alercedb.InsertSampleFeatures(db, 100, oidPool)
	if err != nil {
		return err
	}
	err = alercedb.InsertSampleProbabilities(db, oidPool, []string{"SN", "AGN", "VS", "Asteroid", "Bogus"}, "stamp_classifier")
	if err != nil {
		return err
	}
	return nil
}

func ClearALeRCEDB() error {
	db, err := GetDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DISCARD ALL")
	if err != nil {
		return err
	}
	err = alercedb.DropTables(db)
	if err != nil {
		return err
	}
	return nil
}
