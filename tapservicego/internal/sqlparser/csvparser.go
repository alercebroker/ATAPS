package sqlparser

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
)

// ParseCSV converts a slice of maps to a CSV string
// and returns the string.
// The maps should be keyed by column names
// and the values should be the column values.
//
// The CSV string should be formatted as follows:
//  - Each row should be on a new line
//  - Each column should be separated by a comma
//  - The first row should contain the column names
//
// Example input:
//  [
// 	    {"name": "Alice", "age": 30},
// 	    {"name": "Bob", "age": 25}
//  ]
//
// Example output:
//  name,age
//  Alice,30
//  Bob,25
func ParseCSV(data []map[string]interface{}) (string, error) {
        var csvResult bytes.Buffer
        w := csv.NewWriter(&csvResult)
        err := parsedata(data, w)
        if err != nil {
                return "", err
        }
        return csvResult.String(), nil
}

func ParseTSV(data []map[string]interface{}) (string, error) {
        var tsvResult bytes.Buffer
        w := csv.NewWriter(&tsvResult)
        w.Comma = '\t'
        err := parsedata(data, w)
        if err != nil {
                return "", err
        }
        return tsvResult.String(), nil
}

func parsedata(data []map[string]interface{}, w *csv.Writer) error {
        var headers []string
        for _, row := range data {
                if len(headers) == 0 {
                        headers = getHeaders(row)
                        err := w.Write(headers)
                        if err != nil {
                                log.Printf("Error writing headers: %v", err)
                                return err
                        }
                }
                converted := convertRowValuesToString(row)
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

func convertRowValuesToString(row map[string]interface{}) []string {
        var converted []string
        for _, value := range row {
                converted = append(converted, fmt.Sprintf("%v", value))
        }
        return converted
}

func getHeaders(row map[string]interface{}) []string {
        var headers []string
        for header := range row {
                headers = append(headers, header)
        }
        return headers
}
