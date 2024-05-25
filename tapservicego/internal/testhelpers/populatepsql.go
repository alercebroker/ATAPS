package testhelpers

import "database/sql"

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
