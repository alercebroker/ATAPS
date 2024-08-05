package testhelpers

import (
	"ataps/pkg/alercedb"
	"database/sql"
	"log"
)

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
		log.Println("Error creating tables")
		return err
	}
	oidPool, err := alercedb.InsertSampleObjects(db, 100)
	if err != nil {
		log.Println("Error inserting sample objects")
		return err
	}
	err = alercedb.InsertSampleDetections(db, 1000, oidPool)
	if err != nil {
		log.Println("Error inserting sample detections")
		return err
	}
	err = alercedb.InsertSampleNonDetections(db, 1000, oidPool)
	if err != nil {
		log.Println("Error inserting sample non detections")
		return err
	}
	err = alercedb.InsertSampleForcedPhotometry(db, 1000, oidPool)
	if err != nil {
		log.Println("Error inserting sample forced photometry")
		return err
	}
	err = alercedb.InsertSampleFeatures(db, 100, oidPool)
	if err != nil {
		log.Println("Error inserting sample features")
		return err
	}
	err = alercedb.InsertSampleProbabilities(db, oidPool, []string{"SN", "AGN", "VS", "Asteroid", "Bogus"}, "stamp_classifier")
	if err != nil {
		log.Println("Error inserting sample probabilities")
		return err
	}
	return nil
}

func ClearALeRCEDB(db *sql.DB) error {
	_, err := db.Exec("DISCARD ALL")
	if err != nil {
		return err
	}
	err = alercedb.DropTables(db)
	if err != nil {
		return err
	}
	return nil
}
