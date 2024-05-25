package sqlparser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertRowToString(t *testing.T) {
        row := map[string]interface{}{"name": "Alice", "age": 30}
        expected := []string{"Alice", "30"}
        actual := convertRowValuesToString(row, []string{"name", "age"})
        assert.Equal(t, expected, actual)
}

func TestGetHeaders(t *testing.T) {
        row := map[string]interface{}{"name": "Alice", "age": 30}
        expected := []string{"age", "name"}
        actual := getHeaders(row)
        assert.Equal(t, expected, actual)
}

func TestParseCSV(t *testing.T) {
        data := []map[string]interface{}{
                {"name": "Alice", "age": 30},
                {"name": "Bob", "age": 25},
        }
        expected := "age,name\n30,Alice\n25,Bob\n"
        actual, err := ParseCSV(data)
        if err != nil {
                t.Errorf("Error parsing CSV: %v", err)
        }
        assert.Equal(t, expected, actual)
}

func TestParseTSV(t *testing.T) {
        data := []map[string]interface{}{
                {"name": "Alice", "age": 30},
                {"name": "Bob", "age": 25},
        }
        expected := "age\tname\n30\tAlice\n25\tBob\n"
        actual, err := ParseTSV(data)
        if err != nil {
                t.Errorf("Error parsing TSV: %v", err)
        }
        assert.Equal(t, expected, actual)
}
