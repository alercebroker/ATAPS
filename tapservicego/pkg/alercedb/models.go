package alercedb

// Object represents an object in the database
type Object struct {
	Oid       string  `db:"oid"`
	MeanRA    float64 `db:"meanra"`
	MeanDec   float64 `db:"meandec"`
	SigmaRA   float64 `db:"sigmara"`
	SigmaDec  float64 `db:"sigmadec"`
	FirstMJD  float64 `db:"firstmjd"`
	LastMJD   float64 `db:"lastmjd"`
	NDet      int     `db:"ndet"`
	Stellar   bool    `db:"stellar"`
	Corrected bool    `db:"corrected"`
}

// Detection represents a detection in the database
type Detection struct {
	Candid          int     `db:"candid"`
	Oid             string  `db:"oid"`
	MJD             float64 `db:"mjd"`
	FID             int     `db:"fid"`
	PID             float64 `db:"pid"`
	Diffmaglim      float64 `db:"diffmaglim"`
	Isdiffpos       int     `db:"isdiffpos"`
	RA              float64 `db:"ra"`
	Dec             float64 `db:"dec"`
	Magpsf          float64 `db:"magpsf"`
	Sigmapsf        float64 `db:"sigmapsf"`
	MagpsfCorr      float64 `db:"magpsf_corr"`
	SigmapsfCorr    float64 `db:"sigmapsf_corr"`
	SigmapsfCorrExt float64 `db:"sigmapsf_corr_ext"`
	Distnr          float64 `db:"distnr"`
	Corrected       bool    `db:"corrected"`
	Dubious         bool    `db:"dubious"`
	ParentCandid    int     `db:"parent_candid"`
	HasStamp        bool    `db:"has_stamp"`
}

// NonDetection represents a non-detection in the database
type NonDetection struct {
	Oid        string  `db:"oid"`
	Fid        int     `db:"fid"`
	Mjd        float64 `db:"mjd"`
	Diffmaglim float64 `db:"diffmaglim"`
}

// ForcedPhotometry represents a forced photometry in the database
type ForcedPhotometry struct {
	Candid          int     `db:"candid"`
	Oid             string  `db:"oid"`
	MJD             float64 `db:"mjd"`
	FID             int     `db:"fid"`
	PID             float64 `db:"pid"`
	Diffmaglim      float64 `db:"diffmaglim"`
	Isdiffpos       int     `db:"isdiffpos"`
	RA              float64 `db:"ra"`
	Dec             float64 `db:"dec"`
	Magpsf          float64 `db:"magpsf"`
	Sigmapsf        float64 `db:"sigmapsf"`
	MagpsfCorr      float64 `db:"magpsf_corr"`
	SigmapsfCorr    float64 `db:"sigmapsf_corr"`
	SigmapsfCorrExt float64 `db:"sigmapsf_corr_ext"`
	Distnr          float64 `db:"distnr"`
	Corrected       bool    `db:"corrected"`
	Dubious         bool    `db:"dubious"`
	ParentCandid    int     `db:"parent_candid"`
	HasStamp        bool    `db:"has_stamp"`
}

// Feature represents a feature in the database
type Feature struct {
	Oid     string  `db:"oid"`
	Name    string  `db:"name"`
	Value   float64 `db:"value"`
	Fid     int     `db:"fid"`
	Version string  `db:"version"`
}

// Probability represents a probability in the database
type Probability struct {
	Oid               string  `db:"oid"`
	ClassName         string  `db:"class_name"`
	ClassifierName    string  `db:"classifier_name"`
	ClassifierVersion string  `db:"classifier_version"`
	Probability       float64 `db:"probability"`
	Ranking           int     `db:"ranking"`
}
