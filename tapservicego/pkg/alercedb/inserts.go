package alercedb

import (
	"database/sql"
	"log"
)

func InsertSampleDetections(db *sql.DB, number int, oidPool []string) error {
	detections := generateDetections(number, oidPool)
	err := insertDetections(detections, db)
	if err != nil {
		return err
	}
	return nil
}

func InsertSampleObjects(db *sql.DB, number int) ([]string, error) {
	objects := generateObjects(number)
	oids := make([]string, len(objects))
	for i, object := range objects {
		oids[i] = object.Oid
	}
	err := insertObjects(objects, db)
	if err != nil {
		return nil, err
	}
	return oids, nil
}

func InsertSampleNonDetections(db *sql.DB, number int, oidPool []string) error {
	nonDetections := generateNonDetections(number, oidPool)
	err := insertNonDetections(nonDetections, db)
	if err != nil {
		return err
	}
	return nil
}

func InsertSampleForcedPhotometry(db *sql.DB, number int, oidPool []string) error {
	forcedPhotometry := generateForcedPhotometry(number, oidPool)
	err := insertForcedPhotometry(forcedPhotometry, db)
	if err != nil {
		return err
	}
	return nil
}

func InsertSampleFeatures(db *sql.DB, number int, oidPool []string) error {
	features := generateFeatures(number, oidPool)
	err := insertFeatures(features, db)
	if err != nil {
		return err
	}
	return nil
}

func InsertSampleProbabilities(db *sql.DB, oidPool []string, classPool []string, classifierName string) error {
	probabilities := generateProbabilities(oidPool, classPool, classifierName)
	err := insertProbabilities(probabilities, db)
	if err != nil {
		return err
	}
	return nil
}

// InsertObjects inserts the provided slice of objects
// into the database using the provided database connection
func insertObjects(objects []Object, db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`INSERT INTO object (
        oid, meanra, meandec, sigmara, sigmadec, firstmjd, lastmjd, ndet, stellar, corrected
    ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, object := range objects {
		_, err := stmt.Exec(
			object.Oid,
			object.MeanRA,
			object.MeanDec,
			object.SigmaRA,
			object.SigmaDec,
			object.FirstMJD,
			object.LastMJD,
			object.NDet,
			object.Stellar,
			object.Corrected,
		)
		if err != nil {
			log.Printf("Error with object %s: %v", object.Oid, err)
			return err
		}
	}

	return tx.Commit()
}

// Insert detections into the database
func insertDetections(detections []Detection, db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`INSERT INTO detection (
		candid,
		oid,
		mjd,
		fid,
		pid,
		diffmaglim,
		isdiffpos,
		ra,
		dec,
		magpsf,
		sigmapsf,
		magpsf_corr,
		sigmapsf_corr,
		sigmapsf_corr_ext,
		distnr,
		corrected,
		dubious,
		parent_candid,
		has_stamp
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19);`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, detection := range detections {
		_, err := stmt.Exec(
			detection.Candid,
			detection.Oid,
			detection.MJD,
			detection.FID,
			detection.PID,
			detection.Diffmaglim,
			detection.Isdiffpos,
			detection.RA,
			detection.Dec,
			detection.Magpsf,
			detection.Sigmapsf,
			detection.MagpsfCorr,
			detection.SigmapsfCorr,
			detection.SigmapsfCorrExt,
			detection.Distnr,
			detection.Corrected,
			detection.Dubious,
			detection.ParentCandid,
			detection.HasStamp,
		)
		if err != nil {
			log.Printf("Error with detection %d: %v", detection.Candid, err)
			return err
		}
	}

	return tx.Commit()
}

// insertNonDetections inserts the provided slice of non-detections
func insertNonDetections(nonDetections []NonDetection, db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`INSERT INTO non_detection (
		oid,
		fid,
		mjd,
		diffmaglim
	) VALUES ($1, $2, $3, $4);`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, nonDetection := range nonDetections {
		_, err := stmt.Exec(
			nonDetection.Oid,
			nonDetection.Fid,
			nonDetection.Mjd,
			nonDetection.Diffmaglim,
		)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

// insertForcedPhotometry inserts the provided slice of forced photometry
func insertForcedPhotometry(forcedPhotometry []ForcedPhotometry, db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`INSERT INTO forced_photometry (
		candid,
		oid,
		mjd,
		fid,
		pid,
		diffmaglim,
		isdiffpos,
		ra,
		dec,
		magpsf,
		sigmapsf,
		magpsf_corr,
		sigmapsf_corr,
		sigmapsf_corr_ext,
		distnr,
		corrected,
		dubious,
		parent_candid,
		has_stamp
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19);`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, forcedPhotometry := range forcedPhotometry {
		_, err := stmt.Exec(
			forcedPhotometry.Candid,
			forcedPhotometry.Oid,
			forcedPhotometry.MJD,
			forcedPhotometry.FID,
			forcedPhotometry.PID,
			forcedPhotometry.Diffmaglim,
			forcedPhotometry.Isdiffpos,
			forcedPhotometry.RA,
			forcedPhotometry.Dec,
			forcedPhotometry.Magpsf,
			forcedPhotometry.Sigmapsf,
			forcedPhotometry.MagpsfCorr,
			forcedPhotometry.SigmapsfCorr,
			forcedPhotometry.SigmapsfCorrExt,
			forcedPhotometry.Distnr,
			forcedPhotometry.Corrected,
			forcedPhotometry.Dubious,
			forcedPhotometry.ParentCandid,
			forcedPhotometry.HasStamp,
		)
		if err != nil {
			log.Printf("Error with forced photometry %d: %v", forcedPhotometry.Candid, err)
			return err
		}
	}

	return tx.Commit()
}

// insertFeatures inserts the provided slice of features
func insertFeatures(features []Feature, db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`INSERT INTO feature (
		oid,
		name,
		value,
		fid,
		version
	) VALUES ($1, $2, $3, $4, $5);`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, feature := range features {
		_, err := stmt.Exec(
			feature.Oid,
			feature.Name,
			feature.Value,
			feature.Fid,
			feature.Version,
		)
		if err != nil {
			log.Printf("Error with feature %s: %v", feature.Name, err)
			return err
		}
	}
	return tx.Commit()
}

// insertProbabilities inserts the provided slice of probabilities
func insertProbabilities(probabilities []Probability, db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`INSERT INTO probability (
		oid,
		class_name,
		classifier_name,
		classifier_version,
		probability,
		ranking
	) VALUES ($1, $2, $3, $4, $5, $6);`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, probability := range probabilities {
		_, err := stmt.Exec(
			probability.Oid,
			probability.ClassName,
			probability.ClassifierName,
			probability.ClassifierVersion,
			probability.Probability,
			probability.Ranking,
		)
		if err != nil {
			log.Printf("Error with probability %s: %v", probability.ClassName, err)
			return err
		}
	}
	return tx.Commit()
}
