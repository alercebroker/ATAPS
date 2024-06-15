package parsers

import (
	"bytes"
	"fmt"
)

func ParseText(data []map[string]interface{}) string {
	var textResult bytes.Buffer
	parseTextData(data, &textResult)
	return textResult.String()
}

func parseTextData(data []map[string]interface{}, buffer *bytes.Buffer) {
	fmt.Fprintln(buffer, "# Results:")
	fmt.Fprintln(buffer, "#")
	fmt.Fprintln(buffer, "# Headers:")
	headers := getHeaders(data[0])
	header_line := "# "
	for i, header := range headers {
		if i == len(headers)-1 {
			header_line += fmt.Sprintf("%v", header)
			break
		}
		header_line += fmt.Sprintf("%v | ", header)
	}
	fmt.Fprintln(buffer, header_line)
	for _, row := range data {
		converted := convertRowValuesToString(row, headers)
		row_line := ""
		for i, value := range converted {
			if i == len(converted)-1 {
				row_line += fmt.Sprintf("%v", value)
				break
			}
			row_line += fmt.Sprintf("%v | ", value)
		}
		fmt.Fprintln(buffer, row_line)
	}
}
