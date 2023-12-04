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
		{Value: "(", Type: lexer.Parens},
		{Value: "+", Type: lexer.Op},
		{Value: "1", Type: lexer.Int},
		{Value: "2", Type: lexer.Int},
		{Value: "3", Type: lexer.Int},
		{Value: "4", Type: lexer.Int},
		{Value: ")", Type: lexer.Parens}}

	require.Equal(t, want, got)
}
