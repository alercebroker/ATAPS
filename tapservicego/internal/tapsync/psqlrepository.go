package tapsync

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// GetDB creates a new database connection
// using the provided URL
// and returns a pointer to the connection
// or an error if one occurs.
// The URL should be in the format:
// postgresql://user:password@host:port/database
func GetDB(url string) (*sql.DB, error) {
	db, err := sql.Open("pgx", url)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}


// HandleSQLQuery executes the provided query
// on the provided database connection
// and returns a slice of maps representing the rows
// or an error if one occurs.
// The maps are keyed by the column names
// and the values are the column values.
// The query should be a valid SQL query string.
func HandleSQLQuery(query string, db *sql.DB) ([]map[string]interface {}, error) {
	// Execute the query
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	// Defer closing the rows
	defer rows.Close()
	// Get the column names from the rows
	columns, err := rows.Columns()
	if err != nil {
		log.Printf("Error getting columns: %v", err)
		return nil, err
	}
	// create a slice of maps to hold each row
	var results []map[string]interface{}
	// Create a slice of interfaces to represent each column,
	// and a slice of pointers to each item in the interfaces slice.
	// This is necessary because the Scan function requires pointers
	values := make([]interface{}, len(columns))
	pointers := make([]interface{}, len(columns))
	// Initialize the pointers slice with the addresses of the values slice
	for i := range values {
		pointers[i] = &values[i]
	}
	for rows.Next() {
		// Scan the row into the pointers slice
		err := rows.Scan(pointers...)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}
		// Create a map to hold the row data
		rowMap := make(map[string]interface{})
		// Iterate over the columns and add the data to the map
		for i, colName := range columns {
			val := values[i]
			rowMap[colName] = val
		}
		// Append the map to the results slice
		results = append(results, rowMap)
	}
	// Check for errors during the iteration
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	// Return the results
	return results, nil
}

