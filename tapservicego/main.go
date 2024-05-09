package main

import (
	"fmt"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/alecthomas/repr"
)

type ColumnReference struct {
	Qualifier *Qualifier `parser:"(@@ '.')?"`
	ColumnName string `parser:"@Ident"`
}

type Qualifier struct {
	TableName *TableName `parser:"@@"`
	CorrelationName *string `| @Ident`
}

type TableName struct {
	SchemaName *SchemaName `parser:"(@@ '.')?"`
	Identifier string `parser:"@Ident"`
}

type SchemaName struct {
	CatalogName string `parser:"(@Ident '.')?"`
	UnqualifiedSchemaName string `parser:"@Ident"`
}

var (
	// ADQLLexer is the lexer for the ADQL grammar.
	lex = lexer.MustSimple([]lexer.SimpleRule{
		{`Keyword`, `(?i)\b(SELECT|FROM|TOP|DISTINCT|ALL|WHERE|GROUP|BY|HAVING|UNION|MINUS|EXCEPT|INTERSECT|ORDER|LIMIT|OFFSET|TRUE|FALSE|NULL|IS|NOT|ANY|SOME|BETWEEN|AND|OR|LIKE|AS|IN)\b`},
		{`Ident`, `[a-zA-Z_][a-zA-Z0-9_]*`},
		{`Number`, `[-+]?\d*\.?\d+([eE][-+]?\d+)?`},
		{`String`, `'[^']*'|"[^"]*"`},
		{`Operators`, `<>|!=|<=|>=|[-+*/%,.()=<>]`},
		{"whitespace", `\s+`},
		{"ConcatOp", `\|\|`},
		{"Period", `\.`},
	})
	parser = participle.MustBuild[ColumnReference](
		participle.Lexer(lex),
		participle.Unquote("String"),
		participle.CaseInsensitive("Keyword"),
	)
)

func main() {
    inputs := []string{"a", "a.b", "a.b.c", "a.b.c.d"}
    for _, input := range inputs {
	parsed, err := parser.ParseString("", input)
	if err != nil {
	    fmt.Println(err)
	}
	repr.Println(parsed, repr.Indent("  "), repr.OmitEmpty(true))
    }
}
