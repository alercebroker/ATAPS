package parsers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseText(t *testing.T) {
	data := []map[string]interface{}{
		{"name": "Alice", "age": 30, "city": "New York"},
		{"name": "Bob", "age": 25, "city": "Los Angeles"},
	}
	expected := `# Results:
#
# Headers:
# age | city | name
30 | New York | Alice
25 | Los Angeles | Bob
`
	result := ParseText(data)
	assert.Equal(t, expected, result)
}
