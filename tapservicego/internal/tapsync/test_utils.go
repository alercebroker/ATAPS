package tapsync

import (
	"bufio"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"

	"golang.org/x/net/html"
)

func sendTestQuery(query string, service *TapSyncService) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(
		"POST",
		"/sync",
		strings.NewReader(query),
	)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	service.Router.ServeHTTP(w, req)
	return w
}

func getColumnNames(v interface{}) []string {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	var fieldNames []string

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		columnTag := field.Tag.Get("db")
		if columnTag != "" {
			columnName := strings.Split(columnTag, ",")[0]
			fieldNames = append(fieldNames, columnName)
		} else {
			fieldNames = append(fieldNames, field.Name)
		}
	}

	return fieldNames
}

func parseHTMLTable(doc *html.Node, data *[]string, tag string) {
	var traverse func(n *html.Node, tag string) *html.Node
	traverse = func(n *html.Node, tag string) *html.Node {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == html.TextNode && c.Parent.Data == tag {
				*data = append(*data, c.Data)
			}
			res := traverse(c, tag)
			if res != nil {
				return res
			}
		}
		return nil
	}
	traverse(doc, tag)
}

func parseTextTable(doc string, data *[]string, headers *[]string) error {
	scanner := bufio.NewScanner(strings.NewReader(doc))
	isHeader := false
	for scanner.Scan() {
		line := scanner.Text()
		isComment := strings.HasPrefix(line, "#")
		if isComment {
			if isHeader {
				h := strings.Split(line, "|")
				for _, header := range h {
					header = strings.Replace(header, "#", "", -1)
					*headers = append(*headers, strings.TrimSpace(header))
				}
			}
			isHeader = strings.Contains(line, "Headers")
		} else {
			row := strings.Split(line, "|")
			for _, cell := range row {
				*data = append(*data, strings.TrimSpace(cell))
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
