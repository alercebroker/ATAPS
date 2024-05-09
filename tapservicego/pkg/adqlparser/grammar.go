package adqlparser

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

type ADQLGrammar struct {
	QuerySpecification *QuerySpecification `parser:"'SELECT' @@"`
}

type QuerySpecification struct {
	SetQuantifier *SetQuantifier `parser:"@@*"`
	SetLimit      *SetLimit      `parser:"@@*"`
	SelectList    *SelectList    `parser:"@@"`
	// TableExpression *TableExpression `  @@`
}

type SetQuantifier struct {
	Quantifier string `parser:"@('ALL' | 'DISTINCT')"`
}

type SetLimit struct {
	Limit int `parser:"'TOP' @Number"`
}

type SelectList struct {
	Asterisk      bool             `parser:" @'*'"`
	SelectSublist []*SelectSublist `parser:"| @@ ( ',' @@ )*"`
}

type SelectSublist struct {
	DerivedColumn *DerivedColumn `parser:"@@"`
	// Qualifier *Qualifier `| @@`
	// Period *Period `| @@`
	Asterisk bool `parser:"| @'*'"`
}

type DerivedColumn struct {
	ValueExpression *ValueExpression `parser:"@@"`
	// AsClause *AsClause `@@?`
}

type ValueExpression struct {
	// NumericValueExpression *NumericValueExpression `@@`
	StringValueExpression *StringValueExpression `parser:"@@"`
	// GeometryValueExpression *GeometryValueExpression `| @@`
}

type StringValueExpression struct {
	CharacterValueExpression *CharacterValueExpression `parser:"@@"`
}

type CharacterValueExpression struct {
	CharacterFactor *CharacterFactor `parser:"@@"`
	Concatenation   []*Concatenation `parser:"@@*"`
}

type Concatenation struct {
	Operator        string           `parser:"@ConcatOp"`
	CharacterFactor *CharacterFactor `parser:"@@"`
}

type CharacterFactor struct {
	CharacterPrimary *CharacterPrimary `parser:"@@"`
	// StringValueFunction *StringValueFunction `| @@`
}

type CharacterPrimary struct {
	ValueExpressionPrimary *ValueExpressionPrimary `parser:"@@"`
	// StringValueFunction *StringValueFunction `| @@`
}

type ValueExpressionPrimary struct {
	// UnsignedValueSpecification *UnsignedValueSpecification `@@`
	ColumnReference *ColumnReference `parser:"@@"`
	// SetFunctionSpecification *SetFunctionSpecification `| @@`
	// ValueExpressionWithinParens *ValueExpressionWithinParens `| @@`
}

type ColumnReference struct {
	FullName []string `parser:"@Ident ( '.' @Ident )*"`
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
	parser = participle.MustBuild[ADQLGrammar](
		participle.Lexer(lex),
		participle.Unquote("String"),
		participle.CaseInsensitive("Keyword"),
	)
)
