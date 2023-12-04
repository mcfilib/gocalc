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

	got := parser.Parse(0, input)
	require.NotNil(t, got)

	want := &parser.AST{
		First: &lexer.Token{
			Type:  lexer.Op,
			Value: []rune("+"),
		},
		Rest: []*parser.AST{
			{
				First: &lexer.Token{
					Type:  lexer.Int,
					Value: []rune("1")},
			},
			{
				First: &lexer.Token{
					Type:  lexer.Int,
					Value: []rune("2")},
			},
			{
				First: &lexer.Token{
					Type:  lexer.Int,
					Value: []rune("3")},
			},
			{
				First: &lexer.Token{
					Type:  lexer.Int,
					Value: []rune("4")},
			},
		},
	}

	require.Equal(t, want, got)
}
