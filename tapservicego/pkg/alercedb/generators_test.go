package alercedb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateObjects(t *testing.T) {
	objects := generateObjects(10)
	assert.Equal(t, 10, len(objects))
}

func TestGenerateDetections(t *testing.T) {
	oidPool := []string{"oid1", "oid2", "oid3", "oid4", "oid5"}
	detections := generateDetections(10, oidPool)
	assert.Equal(t, 10, len(detections))
}

func TestGenerateNonDetections(t *testing.T) {
	oidPool := []string{"oid1", "oid2", "oid3", "oid4", "oid5"}
	nonDetections := generateNonDetections(10, oidPool)
	assert.Equal(t, 10, len(nonDetections))
}

func TestGenerateForcedPhotometry(t *testing.T) {
	oidPool := []string{"oid1", "oid2", "oid3", "oid4", "oid5"}
	forcedPhotometry := generateForcedPhotometry(10, oidPool)
	assert.Equal(t, 10, len(forcedPhotometry))
}

func TestGenerateFeatures(t *testing.T) {
	oidPool := []string{"oid1", "oid2", "oid3", "oid4", "oid5"}
	features := generateFeatures(10, oidPool)
	assert.Equal(t, 10, len(features))
}

func TestGenerateProbabilities(t *testing.T) {
	oidPool := []string{"oid1", "oid2", "oid3", "oid4", "oid5"}
	classes := []string{"class1", "class2", "class3", "class4", "class5"}
	probabilities := generateProbabilities(oidPool, classes, "classifier")
	assert.Equal(t, len(oidPool)*len(classes), len(probabilities))
	for i, p := range probabilities {
		assert.Equal(t, "classifier", p.ClassifierName)
		assert.Contains(t, oidPool, p.Oid)
		classIndex := i % len(classes)
		assert.Equal(t, classes[classIndex], p.ClassName)
		assert.Equal(t, classIndex, p.Ranking-1)
	}
}
