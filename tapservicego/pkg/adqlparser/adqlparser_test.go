package adqlparser

import (
	"testing"

	"github.com/alecthomas/repr"
)

func TestParse(t *testing.T) {
	adql, err := Parse("SELECT a")
	if err != nil {
		t.Fatal(err)
	}
	repr.Println(adql, repr.Indent("  "), repr.OmitEmpty(true))
}
