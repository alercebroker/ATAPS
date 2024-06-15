package parsers

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"sort"
)

// ParseCSV converts a slice of maps to a CSV string
// and returns the string.
// The maps should be keyed by column names
// and the values should be the column values.
//
// The CSV string should be formatted as follows:
//   - Each row should be on a new line
//   - Each column should be separated by a comma
//   - The first row should contain the column names
//
// Example input:
//
//	 [
//		    {"name": "Alice", "age": 30},
//		    {"name": "Bob", "age": 25}
//	 ]
//
// Example output:
//
//	name,age
//	Alice,30
//	Bob,25
func ParseCSV(data []map[string]interface{}) (string, error) {
	var csvResult bytes.Buffer
	w := csv.NewWriter(&csvResult)
	err := parseCsvData(data, w)
	if err != nil {
		return "", err
	}
	return csvResult.String(), nil
}

// ParseTSV converts a slice of maps to a TSV string
// and returns the string.
// The maps should be keyed by column names
// and the values should be the column values.
//
// The TSV string should be formatted as follows:
//   - Each row should be on a new line
//   - Each column should be separated by a tab character
//   - The first row should contain the column names
//
// Example input:
//
//	 [
//		    {"name": "Alice", "age": 30},
//		    {"name": "Bob", "age": 25}
//	 ]
//
// Example output:
//
//	name\tage
//	Alice\t30
//	Bob\t25
func ParseTSV(data []map[string]interface{}) (string, error) {
	var tsvResult bytes.Buffer
	w := csv.NewWriter(&tsvResult)
	w.Comma = '\t'
	err := parseCsvData(data, w)
	if err != nil {
		return "", err
	}
	return tsvResult.String(), nil
}

func parseCsvData(data []map[string]interface{}, w *csv.Writer) error {
	headers := getHeaders(data[0])
	err := w.Write(headers)
	if err != nil {
		log.Printf("Error writing headers: %v", err)
		return err
	}
	for _, row := range data {
		converted := convertRowValuesToString(row, headers)
		err := w.Write(converted)
		if err != nil {
			log.Printf("Error writing row: %v", err)
			return err
		}
	}
	w.Flush()
	if err := w.Error(); err != nil {
		log.Printf("Error flushing writer: %v", err)
		return err
	}
	return nil
}

func convertRowValuesToString(row map[string]interface{}, headers []string) []string {
	var converted []string
	for _, header := range headers {
		converted = append(converted, fmt.Sprintf("%v", row[header]))
	}
	return converted
}

func getHeaders(row map[string]interface{}) []string {
	var headers []string
	for header := range row {
		headers = append(headers, header)
	}
	// go map iteration order is random, so we need sort the headers
	sort.Strings(headers)
	return headers
}
