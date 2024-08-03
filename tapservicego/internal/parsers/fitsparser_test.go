package parsers

import (
	// "reflect"

	"os"
	"testing"

	"github.com/astrogo/cfitsio"
	"github.com/stretchr/testify/assert"
)

func TestParseFits(t *testing.T) {
	data := []map[string]interface{}{
		{
			"name":         "Alice",
			"age":          int64(30),
			"boolvalue":    true,
			"nullvalue":    nil,
			"float32value": float32(1.0),
			"float64value": float64(1.0),
		},
		{
			"name":         "Bob",
			"age":          int64(25),
			"boolvalue":    false,
			"nullvalue":    nil,
			"float32value": float32(2.0),
			"float64value": float64(2.0),
		},
		{
			"name":         nil,
			"age":          nil,
			"boolvalue":    nil,
			"nullvalue":    nil,
			"float32value": nil,
			"float64value": nil,
		},
	}
	fname, err := ParseFits(data)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(fname)
	// should be able to read the fits file and parse it
	fitsFile, err := cfitsio.Open(fname, cfitsio.ReadOnly)
	if err != nil {
		t.Fatal(err)
	}
	defer fitsFile.Close()
	table := fitsFile.HDU(1).(*cfitsio.Table)
	assert.True(t, table.Type() == cfitsio.BINARY_TBL)
	assert.Equal(t, "results", table.Name())
	assert.Equal(t, len(data[0])-1, table.NumCols())
	assert.Equal(t, int64(len(data)), table.NumRows())
	rows, err := table.Read(0, table.NumRows())
	assert.Nil(t, err)
	defer rows.Close()
	count := 0
	parsedData := []map[string]interface{}{}
	for rows.Next() {
		row := map[string]interface{}{}
		err = rows.Scan(&row)
		assert.Nil(t, err)
		parsedData = append(parsedData, row)
		count = count + 1
	}
	for i, row := range parsedData {
		if i == len(data)-1 {
			assert.NotContains(t, row, "nullvalue")
			assert.Equal(t, row["name"], " ")
			assert.Equal(t, row["age"], int64(0))
			assert.Equal(t, row["boolvalue"], false)
			assert.Equal(t, row["float32value"], float32(0.0))
			assert.Equal(t, row["float64value"], float64(0.0))
			break
		}
		assert.NotContains(t, row, "nullvalue")
		assert.Equal(t, data[i]["name"], row["name"])
		assert.Equal(t, data[i]["age"], row["age"])
		assert.Equal(t, data[i]["boolvalue"], row["boolvalue"])
		assert.Equal(t, data[i]["float32value"], row["float32value"])
		assert.Equal(t, data[i]["float64value"], row["float64value"])
	}
	assert.Equal(t, len(data), count)
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
	assert.Equal(t, "5A", col[1].Format)
}

func TestCreateFits(t *testing.T) {
	data := []map[string]interface{}{
		{"name": "Alice", "age": 30},
		{"name": "Bob", "age": 25},
	}
	// create fits file
	f, err := os.CreateTemp("", "*.fits")
	if err != nil {
		t.Fatal(err)
	}
	fname := f.Name()
	os.Remove(fname)
	defer os.Remove(fname)
	f.Close() // close immediately since we only need the filename
	fitsFile, err := cfitsio.Create(fname)
	if err != nil {
		t.Fatal(err)
	}
	defer fitsFile.Close()
	// write primary hdu
	phdu, err := cfitsio.NewPrimaryHDU(&fitsFile, cfitsio.NewDefaultHeader())
	if err != nil {
		t.Fatal(err)
	}
	defer phdu.Close()
	result, err := CreateFits(data, &fitsFile)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, result)
	assert.Equal(t, "results", result.Name())
	assert.Equal(t, 2, result.NumCols())
	assert.Equal(t, "age", result.Col(0).Name)
	assert.Equal(t, "name", result.Col(1).Name)
	assert.Equal(t, "K", result.Col(0).Format)
	assert.Equal(t, "5A", result.Col(1).Format)
}
