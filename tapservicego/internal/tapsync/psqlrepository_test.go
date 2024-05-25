package tapsync

import (
	"ataps/internal/testhelpers"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestSimpleSQLQuery(t *testing.T) {
	db, err := GetDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		t.Fatal(err)
	}
	query := "SELECT 'test'"
	result, err := HandleSQLQuery(query, db)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 1, len(result))
	assert.Equal(t, "test", result[0]["?column?"])
}

func TestHandleSQLQuery(t *testing.T) {
	db, err := GetDB(os.Getenv("DATABASE_URL"))
	testhelpers.PopulateDb(db)
	defer testhelpers.ClearDataFromTable(db)
	if err != nil {
		t.Fatal(err)
	}
	query := "SELECT * FROM test"
	result, err := HandleSQLQuery(query, db)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 1, len(result))
	assert.Equal(t, "test", result[0]["name"])
	assert.Equal(t, int64(1), result[0]["number"])
}
