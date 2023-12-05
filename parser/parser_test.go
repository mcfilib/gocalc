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

	want := []*parser.Node{
		{
			Atom: nil,
			List: []*parser.Node{
				{Atom: &lexer.Token{Value: []rune("+"), Type: lexer.Op}},
				{Atom: &lexer.Token{Value: []rune("1"), Type: lexer.Int}},
				{Atom: &lexer.Token{Value: []rune("2"), Type: lexer.Int}},
				{Atom: &lexer.Token{Value: []rune("3"), Type: lexer.Int}},
				{Atom: &lexer.Token{Value: []rune("4"), Type: lexer.Int}}},
		},
	}

	_, got := parser.Parse(0, input)
	require.Equal(t, want, got)
}
