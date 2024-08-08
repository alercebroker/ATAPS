package parsers

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"sort"
	"strings"

	"github.com/astrogo/cfitsio"
)

func ParseFits(data []map[string]interface{}) (string, error) {
	// create fits file
	f, err := os.CreateTemp("", "*.fits")
	if err != nil {
		return "", err
	}
	fname := f.Name()
	f.Close() // close immediately since we only need the filename
	os.Remove(fname)
	fitsFile, err := cfitsio.Create(fname)
	if err != nil {
		return "", err
	}
	// write primary hdu
	phdu, err := cfitsio.NewPrimaryHDU(&fitsFile, cfitsio.NewDefaultHeader())
	if err != nil {
		return "", err
	}
	defer phdu.Close()
	// create table
	fitsTable, err := CreateFits(data, &fitsFile)
	if err != nil {
		return "", err
	}
	defer fitsTable.Close()
	err = fitsFile.Close()
	if err != nil {
		return "", err
	}
	return fname, err
}

func CreateFits(data []map[string]interface{}, file *cfitsio.File) (*cfitsio.Table, error) {
	// create columns using the keys of the first map
	// the type of the column is determined by the type of the value
	columns, err := createColumns(data)
	if err != nil {
		log.Printf("Error creating columns: %v", err)
		return nil, err
	}
	data = cleanData(data, columns)
	// create table from columns
	table, err := cfitsio.NewTable(file, "results", columns, cfitsio.BINARY_TBL)
	if err != nil {
		return nil, err
	}
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

func createColumns(data []map[string]interface{}) ([]cfitsio.Column, error) {
	columns := make([]cfitsio.Column, 0, len(data[0]))
	keys := make([]string, 0, len(data[0]))
	for key := range data[0] {
		keys = append(keys, key)
	}
	// sort the keys to ensure consistent column order
	sort.Strings(keys)
	var format string
	for _, key := range keys {
		for i := 0; i < len(data); i++ {
			if data[i][key] == nil {
				continue
			}
			value := data[i][key]
			format = getFormat(value)
			if format == "" {
				return nil, fmt.Errorf("Error creating columns: unsupported type %T for key %s", value, key)
			}
			columns = append(columns, cfitsio.Column{
				Name:   key,
				Format: format,
			})
			break
		}
	}
	return columns, nil
}

func getFormat(value interface{}) string {
	switch value.(type) {
	case int16:
		return "I" // 16-bit integer
	case int32:
		return "J" // 32-bit integer
	case int:
		return "K" // 64-bit integer
	case int64:
		return "K" // 64-bit integer
	case float32:
		return "E" // 32-bit floating point
	case float64:
		return "D" // 64-bit floating point
	case string:
		stringLength := len(value.(string))
		return fmt.Sprintf("%dA", stringLength) // Character string with length
	case bool:
		return "L" // logical
	default:
		return ""
	}
}

func cleanData(data []map[string]interface{}, columns []cfitsio.Column) []map[string]interface{} {
	cleanedData := make([]map[string]interface{}, len(data))
	for i := 0; i < len(data); i++ {
		cleanedData[i] = make(map[string]interface{}, len(columns))
	}
	keys := make([]string, 0, len(data[0]))
	for key := range data[0] {
		keys = append(keys, key)
	}
	// sort the keys to ensure consistent column order
	sort.Strings(keys)
	for _, key := range keys {
		for i := 0; i < len(data); i++ {
			if data[i][key] != nil {
				cleanedData[i][key] = data[i][key]
				continue
			}
			for j := 0; j < len(columns); j++ {
				if columns[j].Name != key {
					continue
				}
				format := columns[j].Format
				cleanedData[i][key] = getZeroValueForFormat(format)
			}
		}
	}
	return cleanedData
}

func getZeroValueForFormat(format string) interface{} {
	if strings.Contains(format, "A") {
		format = "A"
	}
	switch format {
	case "I":
		return int16(0)
	case "J":
		return int32(0)
	case "K":
		return int64(0)
	case "E":
		return float32(0)
	case "D":
		return float64(0)
	case "A":
		return ""
	case "L":
		return false
	default:
		return nil
	}
}
