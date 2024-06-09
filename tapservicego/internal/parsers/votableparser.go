package parsers

import (
	"ataps/pkg/votable"
	"encoding/xml"
	"fmt"
	"sort"
	"strings"
)

func CreateVOTable(data []map[string]interface{}) (votable.VOTable, error) {
	result := votable.VOTable{
		Version: "1.4",
		Xmlns:   "http://www.ivoa.net/xml/VOTable/v1.4",
		Resource: votable.Resource{
			Type:  "results",
			Infos: []votable.Info{{Name: "QUERY_STATUS", Value: "OK"}},
			Tables: []votable.Table{
				{
					Name:        "results",
					Description: "Results of the query",
				},
			},
		},
	}
	result.Resource.Tables[0].Fields = addFields(data)
	result.Resource.Tables[0].Data.TableData = addTableData(data)
	return result, nil
}

func addFields(data []map[string]interface{}) []votable.Field {
	if len(data) == 0 {
		return []votable.Field{}
	}
	fields := []votable.Field{}
	var keys []string
	for key := range data[0] {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		// TODO: add datatype, unit, description, etc.
		fields = append(fields, votable.Field{Name: key})
	}
	return fields
}

func addTableData(data []map[string]interface{}) votable.TableData {
	tableData := votable.TableData{}
	for _, row := range data {
		tableData.Rows = append(tableData.Rows, votable.Row{Columns: addColumns(row)})
	}
	return tableData
}

func addColumns(row map[string]interface{}) []votable.Column {
	columns := []votable.Column{}
	keys := make([]string, 0, len(row))
	for key := range row {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		columns = append(columns, votable.Column{Value: fmt.Sprintf("%v", row[key])})
	}
	return columns
}

func VOTableToXML(votable votable.VOTable) (string, error) {
	var xmlBuilder strings.Builder
	encoder := xml.NewEncoder(&xmlBuilder)
	xmlBuilder.WriteString(xml.Header)
	encoder.Indent("", "\t")
	err := encoder.Encode(votable)
	if err != nil {
		return "", err
	}
	return xmlBuilder.String(), nil
}
