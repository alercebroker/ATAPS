package tapsync

import (
	"ataps/pkg/votable"
	"strconv"
)

// getErrorVOTable returns a VOTable with an error message
func getErrorVOTable(err error, code int) votable.VOTable {
	errorMessage := err.Error()
	return votable.VOTable{
		Version: "1.4",
		Xmlns:   "http://www.ivoa.net/xml/VOTable/v1.4",
		Resource: votable.Resource{
			Type: "results",
			Infos: []votable.Info{
				{Name: "QUERY_STATUS", Value: "ERROR"},
				{Name: "ERROR_DETAIL", Description: errorMessage},
				{Name: "ERROR_CODE", Value:	strconv.Itoa(code)},
			},
		},
	}
}
