package sqlparser

import (
	"bytes"
	"fmt"
	"log"
	"reflect"
	"sort"

	"github.com/dirodriguezm/fitsio"
)

func ParseFits(data []map[string]interface{}) (*bytes.Buffer, error) {
	// create buffer
	var fitsResult bytes.Buffer
	f, err := fitsio.Create(&fitsResult)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	// write primary hdu
	phdu, err := fitsio.NewPrimaryHDU(nil)
	if err != nil {
		return nil, err
	}
	err = f.Write(phdu)
	if err != nil {
		return nil, err
	}
	// create table
	fitsTable, err := CreateFits(data)
	if err != nil {
		return nil, err
	}
	defer fitsTable.Close()
	// write table
	err = f.Write(fitsTable)
	if err != nil {
		return nil, err
	}
	err = f.Close()
	if err != nil {
		return nil, err
	}
	return &fitsResult, nil
}

func CreateFits(data []map[string]interface{}) (*fitsio.Table, error) {
	// create columns using the keys of the first map
	// the type of the column is determined by the type of the value
	columns, err := createColumns(data)
	if err != nil {
		log.Printf("Error creating columns: %v", err)
		return nil, err
	}
	// create table from columns
	table, err := fitsio.NewTable("results", columns, fitsio.BINARY_TBL)
	if err != nil {
		return nil, err
	}
	defer table.Close()
	// populate the table
	rslice := reflect.ValueOf(data)
	for i := 0; i < rslice.Len(); i++ {
		row := rslice.Index(i).Addr()
		err := table.Write(row.Interface())
		if err != nil {
			log.Printf("Error writing record: %v", err)
			return nil, err
		}
	}
	nrows := table.NumRows()
	if nrows != int64(len(data)) {
		return nil, fmt.Errorf("Error creating table: number of rows written (%d) does not match number of rows in data (%d)", nrows, len(data))
	}
	return table, nil
}

func createColumns(data []map[string]interface{}) ([]fitsio.Column, error) {
	columns := make([]fitsio.Column, 0, len(data[0]))
	maxStringLength := getMaxStringLength(data)
	keys := make([]string, 0, len(data[0]))
	for key := range data[0] {
		keys = append(keys, key)
	}
	// sort the keys to ensure consistent column order
	sort.Strings(keys)
	for _, key := range keys {
		var format string
		value := data[0][key]
		switch value.(type) {
		case int32:
			format = "J" // 32-bit integer
		case int:
			format = "K" // integer
		case int64:
			format = "K" // 64-bit integer
		case float64:
			format = "D" // 64-bit floating point
		case string:
			format = fmt.Sprintf("%dA", maxStringLength+1) // Character string with length
		default:
			return nil, fmt.Errorf("unsupported data type for column %s", key)
		}
		columns = append(columns, fitsio.Column{
			Name:   key,
			Format: format,
		})
	}
	return columns, nil
}

// getMaxStringLength determines the maximum length of string values in the data
func getMaxStringLength(data []map[string]interface{}) int {
	maxLen := 0
	for _, row := range data {
		for _, value := range row {
			if str, ok := value.(string); ok {
				if len(str) > maxLen {
					maxLen = len(str)
				}
			}
		}
	}
	return maxLen
}
