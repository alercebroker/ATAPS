package sqlparser

import (
	// "reflect"

	"testing"

	"github.com/dirodriguezm/fitsio"
	"github.com/stretchr/testify/assert"
)

func TestParseFits(t *testing.T) {
	data := []map[string]interface{}{
		{"name": "Alice", "age": int64(30)},
		{"name": "Bob", "age": int64(25)},
	}
	result, err := ParseFits(data)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	// should be able to read the fits file and parse it
	f, err := fitsio.Open(result)
	assert.Nil(t, err)
	defer f.Close()
	hdu := f.HDU(1)
	table := hdu.(*fitsio.Table)
	assert.True(t, table.Type() == fitsio.BINARY_TBL)
	assert.Equal(t, "results", table.Name())
	assert.Equal(t, 2, table.NumCols())
	assert.Equal(t, int64(2), table.NumRows())
	rows, err := table.Read(0, table.NumRows())
	assert.Nil(t, err)
	defer rows.Close()
	count := 0
	parsedData := []map[string]interface{}{}
	for rows.Next() {
		row := map[string]interface{}{}
		err = rows.Scan(&row)
		assert.Nil(t, err)
		assert.Equal(t, data[count]["name"], row["name"])
		assert.Equal(t, data[count]["age"], row["age"])
		parsedData = append(parsedData, row)
		count = count + 1
	}
	t.Logf("Parsed data: %v", parsedData)
	assert.Equal(t, 2, count)
}

func TestCreateColumns(t *testing.T) {
	data := []map[string]interface{}{
		{"name": "Alice", "age": 30},
		{"name": "Bob", "age": 25},
	}
	col, err := createColumns(data)
	assert.Nil(t, err)
	assert.NotNil(t, col)
	assert.Equal(t, 2, len(col))
	assert.Equal(t, "age", col[0].Name)
	assert.Equal(t, "name", col[1].Name)
	assert.Equal(t, "K", col[0].Format)
	assert.Equal(t, "6A", col[1].Format)
}

func TestCreateFits(t *testing.T) {
	data := []map[string]interface{}{
		{"name": "Alice", "age": 30},
		{"name": "Bob", "age": 25},
	}
	result, err := CreateFits(data)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "results", result.Name())
	assert.Equal(t, 2, result.NumCols())
	assert.Equal(t, "age", result.Col(0).Name)
	assert.Equal(t, "name", result.Col(1).Name)
	assert.Equal(t, "K", result.Col(0).Format)
	assert.Equal(t, "6A", result.Col(1).Format)
}
