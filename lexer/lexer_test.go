package lexer_test

import (
	"testing"

	"github.com/mcfilib/gocalc/lexer"
	"github.com/stretchr/testify/require"
)

func TestLex(t *testing.T) {
	got := lexer.Lex([]rune("(+ 1 2 3 4)"))
	require.NotNil(t, got)

	want := []*lexer.Token{
		{Value: []rune("("), Type: lexer.Parens},
		{Value: []rune("+"), Type: lexer.Op},
		{Value: []rune("1"), Type: lexer.Int},
		{Value: []rune("2"), Type: lexer.Int},
		{Value: []rune("3"), Type: lexer.Int},
		{Value: []rune("4"), Type: lexer.Int},
		{Value: []rune(")"), Type: lexer.Parens}}

	require.Equal(t, want, got)
}
