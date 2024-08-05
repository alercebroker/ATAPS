package alercedb

import (
	"database/sql"
	"fmt"
	"log"
)

func CreateTables(db *sql.DB) error {
	err := createObjectTable(db)
	if err != nil {
		log.Println("Error creating object table")
		return err
	}
	err = createDetectionsTable(db)
	if err != nil {
		log.Println("Error creating detections table")
		return err
	}
	err = createNonDetectionsTable(db)
	if err != nil {
		log.Println("Error creating non detections table")
		return err
	}
	err = createForcedPhotometryTable(db)
	if err != nil {
		log.Println("Error creating forced photometry table")
		return err
	}
	err = createFeaturesTable(db)
	if err != nil {
		log.Println("Error creating features table")
		return err
	}
	err = createProbabilitiesTable(db)
	if err != nil {
		log.Println("Error creating probabilities table")
		return err
	}
	return nil
}

func DropTables(db *sql.DB) error {
	_, err := db.Exec(`
		DROP SCHEMA public CASCADE;
		CREATE SCHEMA public;
		GRANT ALL ON SCHEMA public TO testuser;
		GRANT ALL ON SCHEMA public TO public;
		`)
	return err
}

func createObjectTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS object (
		oid VARCHAR(12) PRIMARY KEY,
		meanra DOUBLE PRECISION NOT NULL,
		meandec DOUBLE PRECISION NOT NULL,
		sigmara DOUBLE PRECISION NOT NULL,
		sigmadec DOUBLE PRECISION NOT NULL,
		firstmjd DOUBLE PRECISION NOT NULL,
		lastmjd DOUBLE PRECISION NOT NULL,
		ndet INTEGER NOT NULL,
		stellar BOOLEAN NOT NULL,
		corrected BOOLEAN NOT NULL
	);`
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	err = createIndex(db, "ix_object_ndet", "object", "oid")
	if err != nil {
		return err
	}
	err = createIndex(db, "ix_object_firstmjd", "object", "firstmjd")
	if err != nil {
		return err
	}
	err = createIndex(db, "ix_object_meanra", "object", "meanra")
	if err != nil {
		return err
	}
	err = createIndex(db, "ix_object_meandec", "object", "meandec")
	if err != nil {
		return err
	}
	return nil
}

// Create detections table
func createDetectionsTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS detection (
		candid BIGINT,
		oid VARCHAR(12),
		mjd DOUBLE PRECISION NOT NULL,
		fid INTEGER NOT NULL,
		pid DOUBLE PRECISION NOT NULL,
		diffmaglim DOUBLE PRECISION,
		isdiffpos SMALLINT NOT NULL,
		ra DOUBLE PRECISION NOT NULL,
		dec DOUBLE PRECISION NOT NULL,
		magpsf DOUBLE PRECISION NOT NULL,
		sigmapsf DOUBLE PRECISION NOT NULL,
		magpsf_corr DOUBLE PRECISION, 
		sigmapsf_corr DOUBLE PRECISION,
		sigmapsf_corr_ext DOUBLE PRECISION,
		distnr DOUBLE PRECISION,
		corrected BOOLEAN NOT NULL,
		dubious BOOLEAN NOT NULL,
		parent_candid BIGINT,
		has_stamp BOOLEAN NOT NULL,
		PRIMARY KEY (candid, oid),
		FOREIGN KEY (oid) REFERENCES object(oid)
	);`
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	err = createIndexUsing(db, "ix_detection_oid", "detection", "oid", "hash")
	if err != nil {
		return err
	}
	return nil
}

func createNonDetectionsTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS non_detection (
		oid VARCHAR(12),
		fid INTEGER NOT NULL,
		mjd DOUBLE PRECISION NOT NULL,
		diffmaglim DOUBLE PRECISION,
		PRIMARY KEY (oid, fid, mjd),
		FOREIGN KEY (oid) REFERENCES object(oid)
	);`
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	err = createIndexUsing(db, "ix_nondetection_oid", "non_detection", "oid", "hash")
	if err != nil {
		return err
	}
	return nil
}

func createForcedPhotometryTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS forced_photometry (
		candid BIGINT,
		oid VARCHAR(12),
		mjd DOUBLE PRECISION NOT NULL,
		fid INTEGER NOT NULL,
		pid DOUBLE PRECISION NOT NULL,
		diffmaglim DOUBLE PRECISION,
		isdiffpos SMALLINT NOT NULL,
		ra DOUBLE PRECISION NOT NULL,
		dec DOUBLE PRECISION NOT NULL,
		magpsf DOUBLE PRECISION NOT NULL,
		sigmapsf DOUBLE PRECISION NOT NULL,
		magpsf_corr DOUBLE PRECISION, 
		sigmapsf_corr DOUBLE PRECISION,
		sigmapsf_corr_ext DOUBLE PRECISION,
		distnr DOUBLE PRECISION,
		corrected BOOLEAN NOT NULL,
		dubious BOOLEAN NOT NULL,
		parent_candid BIGINT,
		has_stamp BOOLEAN NOT NULL,
		PRIMARY KEY (candid, oid),
		FOREIGN KEY (oid) REFERENCES object(oid)
	);`
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	err = createIndexUsing(db, "ix_forced_photometry_oid", "forced_photometry", "oid", "hash")
	if err != nil {
		return err
	}
	return nil
}

func createFeaturesTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS feature (
		oid VARCHAR(12),
		name VARCHAR(255) NOT NULL,
		value DOUBLE PRECISION,
		fid INTEGER NOT NULL,
		version VARCHAR(16) NOT NULL,
		PRIMARY KEY (oid, name, fid, version),
		FOREIGN KEY (oid) REFERENCES object(oid)
	);`
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	err = createIndexUsing(db, "ix_feature_oid_2", "feature", "oid", "hash")
	if err != nil {
		return err
	}
	return nil
}

func createProbabilitiesTable(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS probability (
		oid VARCHAR(12),
		class_name VARCHAR(255) NOT NULL,
		classifier_name VARCHAR(255) NOT NULL,
		classifier_version VARCHAR(16) NOT NULL,
		probability DOUBLE PRECISION NOT NULL,
		ranking INTEGER NOT NULL,
		PRIMARY KEY (oid, class_name, classifier_name, classifier_version),
		FOREIGN KEY (oid) REFERENCES object(oid)
	);`
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	err = createIndexUsing(db, "ix_probability_oid", "probability", "oid", "hash")
	if err != nil {
		return err
	}
	err = createIndex(db, "ix_probability_probability", "probability", "probability")
	if err != nil {
		return err
	}
	err = createIndex(db, "ix_probability_ranking", "probability", "ranking")
	if err != nil {
		return err
	}
	err = createIndexWhere(db, "ix_probability_ranking_1", "probability", "ranking", "ranking = 1")
	if err != nil {
		return err
	}
	return nil
}

func createIndex(db *sql.DB, index_name string, table string, column string) error {
	query := fmt.Sprintf("CREATE INDEX IF NOT EXISTS %s ON %s (%s);", index_name, table, column)
	_, err := db.Exec(query)
	return err
}

func createIndexUsing(db *sql.DB, index_name string, table string, column string, using string) error {
	query := fmt.Sprintf("CREATE INDEX IF NOT EXISTS %s ON %s USING %s (%s);", index_name, table, using, column)
	_, err := db.Exec(query)
	return err
}

func createIndexWhere(db *sql.DB, index_name string, table string, column string, where string) error {
	query := fmt.Sprintf("CREATE INDEX IF NOT EXISTS %s ON %s (%s) WHERE %s;", index_name, table, column, where)
	_, err := db.Exec(query)
	return err
}
