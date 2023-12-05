package evaluator_test

import (
	"testing"

	"github.com/mcfilib/gocalc/evaluator"
	"github.com/mcfilib/gocalc/lexer"
	"github.com/mcfilib/gocalc/parser"
	"github.com/stretchr/testify/require"
)

func TestEval(t *testing.T) {
	tokens := lexer.Lex([]rune("(+ 1 2 3 (+ 4 5 6 (+ 7 8 9)))"))
	require.NotNil(t, tokens)

	_, ast := parser.Parse(0, tokens)
	require.NotNil(t, ast)

	got := evaluator.Eval(ast)
	require.Equal(t, 45, got)
}
