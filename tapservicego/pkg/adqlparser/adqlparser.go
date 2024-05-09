package adqlparser

import (
)

func Parse(adql string) (*ADQLGrammar, error) {
	parsedadql, err := parser.ParseString("", adql)
	if err != nil {
		return nil, err
	}
	return parsedadql, nil
}
