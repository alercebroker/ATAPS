package parsers

import (
	"html/template"
	"io"
)

func ParseHTML(data []map[string]interface{}, writer io.Writer) error {
	const tpl = `<!DOCTYPE html>
<html>
	<head>
		<title>Results</title>
	</head>
	<body>
		<h1>Results</h1>
		<table>
			<tr>
				{{- range $key, $value := index . 0 }}
				<th>{{$key}}</th>
				{{- end }}
			</tr>
			{{- range $index, $row := .}}
			<tr>
				{{- range $key, $value := $row }}
				<td>{{$value}}</td>
				{{- end }}
			</tr>
			{{- end}}
		</table>
	</body>
</html>
`
	t, err := template.New("html").Parse(tpl)
	if err != nil {
		return err
	}
	err = t.Execute(writer, data)
	if err != nil {
		return err
	}
	return nil
}
