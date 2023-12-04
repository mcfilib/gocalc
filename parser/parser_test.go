package parser_test

import (
	"testing"

	"github.com/mcfilib/gocalc/lexer"
	"github.com/mcfilib/gocalc/parser"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	input := lexer.Lex([]rune("(+ 1 2 3 4)"))
	require.NotNil(t, input)

	got := parser.Parse(input)
	require.NotNil(t, got)

	want := &parser.AST{
		First: &lexer.Token{
			Type:  lexer.Op,
			Value: []rune("+"),
		},
		Rest: []*parser.AST{},
	}

	require.Equal(t, want, got)
}
