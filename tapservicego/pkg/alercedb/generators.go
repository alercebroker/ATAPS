package alercedb

import (
	"fmt"
	"math/rand/v2"
	"sort"
)

// GenerateObjects creates a slice of Object structs
// with the provided number of objects
// and each object is created with random values
func generateObjects(number int) []Object {
	objects := make([]Object, number)
	for i := 1; i <= number; i++ {
		oid := fmt.Sprintf("oid%d", i)
		meanRA := rand.Float64() * 360
		meanDec := rand.Float64() * 180
		sigmaRA := rand.Float64() * 10
		sigmaDec := rand.Float64() * 10
		firstMJD := rand.Float64() * 10000
		lastMJD := firstMJD + rand.Float64()*100
		nDet := rand.IntN(1000)
		stellar := rand.IntN(2) == 1
		corrected := rand.IntN(2) == 1
		objects[i-1] = Object{
			Oid:       oid,
			MeanRA:    meanRA,
			MeanDec:   meanDec,
			SigmaRA:   sigmaRA,
			SigmaDec:  sigmaDec,
			FirstMJD:  firstMJD,
			LastMJD:   lastMJD,
			NDet:      nDet,
			Stellar:   stellar,
			Corrected: corrected,
		}
	}
	return objects
}

// generateDetections creates a slice of Detection structs
func generateDetections(number int, oidPool []string) []Detection {
	detections := make([]Detection, number)
	for i := 0; i < number; i++ {
		candid := i + 100001
		oid := oidPool[rand.IntN(len(oidPool))]
		mjd := rand.Float64() * 10000
		fid := rand.IntN(5)
		pid := rand.Float64() * 10000
		diffmaglim := rand.Float64() * 100
		isdiffpos := rand.IntN(2)
		ra := rand.Float64() * 360
		dec := rand.Float64() * 180
		magpsf := rand.Float64() * 100
		sigmapsf := rand.Float64() * 100
		magpsfCorr := rand.Float64() * 100
		sigmapsfCorr := rand.Float64() * 100
		sigmapsfCorrExt := rand.Float64() * 100
		distnr := rand.Float64() * 100
		corrected := rand.IntN(2) == 1
		dubious := rand.IntN(2) == 1
		parentCandid := rand.IntN(100000)
		hasStamp := rand.IntN(2) == 1
		detections[i] = Detection{
			Candid:          candid,
			Oid:             oid,
			MJD:             mjd,
			FID:             fid,
			PID:             pid,
			Diffmaglim:      diffmaglim,
			Isdiffpos:       isdiffpos,
			RA:              ra,
			Dec:             dec,
			Magpsf:          magpsf,
			Sigmapsf:        sigmapsf,
			MagpsfCorr:      magpsfCorr,
			SigmapsfCorr:    sigmapsfCorr,
			SigmapsfCorrExt: sigmapsfCorrExt,
			Distnr:          distnr,
			Corrected:       corrected,
			Dubious:         dubious,
			ParentCandid:    parentCandid,
			HasStamp:        hasStamp,
		}
	}
	return detections
}

// generateNonDetections creates a slice of NonDetection structs
func generateNonDetections(number int, oidPool []string) []NonDetection {
	nonDetections := make([]NonDetection, number)
	for i := 0; i < number; i++ {
		oid := oidPool[rand.IntN(len(oidPool))]
		fid := rand.IntN(5)
		mjd := rand.Float64() * 10000
		diffmaglim := rand.Float64() * 100
		nonDetections[i] = NonDetection{
			Oid:        oid,
			Fid:        fid,
			Mjd:        mjd,
			Diffmaglim: diffmaglim,
		}
	}
	return nonDetections
}

// generateForcedPhotometry creates a slice of ForcedPhotometry structs
func generateForcedPhotometry(number int, oidPool []string) []ForcedPhotometry {
	forcedPhotometry := make([]ForcedPhotometry, number)
	for i := 0; i < number; i++ {
		candid := i + 100001
		oid := oidPool[rand.IntN(len(oidPool))]
		mjd := rand.Float64() * 10000
		fid := rand.IntN(5)
		pid := rand.Float64() * 10000
		diffmaglim := rand.Float64() * 100
		isdiffpos := rand.IntN(2)
		ra := rand.Float64() * 360
		dec := rand.Float64() * 180
		magpsf := rand.Float64() * 100
		sigmapsf := rand.Float64() * 100
		magpsfCorr := rand.Float64() * 100
		sigmapsfCorr := rand.Float64() * 100
		sigmapsfCorrExt := rand.Float64() * 100
		distnr := rand.Float64() * 100
		corrected := rand.IntN(2) == 1
		dubious := rand.IntN(2) == 1
		parentCandid := rand.IntN(100000)
		hasStamp := rand.IntN(2) == 1
		forcedPhotometry[i] = ForcedPhotometry{
			Candid:          candid,
			Oid:             oid,
			MJD:             mjd,
			FID:             fid,
			PID:             pid,
			Diffmaglim:      diffmaglim,
			Isdiffpos:       isdiffpos,
			RA:              ra,
			Dec:             dec,
			Magpsf:          magpsf,
			Sigmapsf:        sigmapsf,
			MagpsfCorr:      magpsfCorr,
			SigmapsfCorr:    sigmapsfCorr,
			SigmapsfCorrExt: sigmapsfCorrExt,
			Distnr:          distnr,
			Corrected:       corrected,
			Dubious:         dubious,
			ParentCandid:    parentCandid,
			HasStamp:        hasStamp,
		}
	}
	return forcedPhotometry
}

// generateFeatures creates a slice of Feature structs
func generateFeatures(number int, oidPool []string) []Feature {
	features := make([]Feature, number)
	for i := 0; i < number; i++ {
		oid := oidPool[rand.IntN(len(oidPool))]
		name := fmt.Sprintf("feature%d", i)
		value := rand.Float64() * 100
		fid := rand.IntN(3) + 1
		version := "test"
		features[i] = Feature{
			Oid:     oid,
			Name:    name,
			Value:   value,
			Fid:     fid,
			Version: version,
		}
	}
	return features
}

// generateProbabilities creates a slice of Probability structs
func generateProbabilities(oidPool []string, classesPool []string, classifierName string) []Probability {
	probabilities := make([]Probability, len(oidPool)*len(classesPool))
	probValues := make([]float64, len(classesPool))
	for i := 0; i < len(classesPool); i++ {
		probValues[i] = rand.Float64()
	}
	sort.Float64s(probValues)
	for i := 0; i < len(oidPool); i++ {
		for class := 0; class < len(classesPool); class++ {
			probabilities[len(classesPool)*i+class] = Probability{
				Oid:               oidPool[i],
				ClassName:         classesPool[class],
				ClassifierName:    classifierName,
				ClassifierVersion: "1.0.0",
				Probability:       probValues[class],
				Ranking:           class + 1,
			}
		}
	}
	return probabilities
}
