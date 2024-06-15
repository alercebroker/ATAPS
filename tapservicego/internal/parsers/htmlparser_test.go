package parsers

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseHTML(t *testing.T) {
	data := []map[string]interface{}{
		{"name": "Alice", "age": 30},
		{"name": "Bob", "age": 25},
	}
	var htmlResult bytes.Buffer
	err := ParseHTML(data, &htmlResult)
	expected := `<!DOCTYPE html>
<html>
	<head>
		<title>Results</title>
	</head>
	<body>
		<h1>Results</h1>
		<table>
			<tr>
				<th>age</th>
				<th>name</th>
			</tr>
			<tr>
				<td>30</td>
				<td>Alice</td>
			</tr>
			<tr>
				<td>25</td>
				<td>Bob</td>
			</tr>
		</table>
	</body>
</html>
`
	assert.Nil(t, err)
	assert.Equal(t, expected, htmlResult.String())
}
