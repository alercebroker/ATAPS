package sqlparser

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertRowToString(t *testing.T) {
        row := map[string]interface{}{"name": "Alice", "age": 30}
        expected := []string{"Alice", "30"}
        actual := convertRowValuesToString(row)
        if !reflect.DeepEqual(actual, expected) {
                t.Errorf("Expected %v, got %v", expected, actual)
        }
}

func TestGetHeaders(t *testing.T) {
        row := map[string]interface{}{"name": "Alice", "age": 30}
        expected := []string{"name", "age"}
        actual := getHeaders(row)
        if !reflect.DeepEqual(actual, expected) {
                t.Errorf("Expected %v, got %v", expected, actual)
        }
}

func TestParseCSV(t *testing.T) {
        data := []map[string]interface{}{
                {"name": "Alice", "age": 30},
                {"name": "Bob", "age": 25},
        }
        expected := "name,age\nAlice,30\nBob,25\n"
        actual, err := ParseCSV(data)
        if err != nil {
                t.Errorf("Error parsing CSV: %v", err)
        }
        if actual != expected {
                assert.Equal(t, expected, actual)
        }
}

func TestParseTSV(t *testing.T) {
        data := []map[string]interface{}{
                {"name": "Alice", "age": 30},
                {"name": "Bob", "age": 25},
        }
        expected := "name\tage\nAlice\t30\nBob\t25\n"
        actual, err := ParseTSV(data)
        if err != nil {
                t.Errorf("Error parsing TSV: %v", err)
        }
        if actual != expected {
                assert.Equal(t, expected, actual)
        }
}
