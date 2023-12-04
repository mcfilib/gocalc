package lexer_test

import (
	"testing"

	"github.com/mcfilib/gocalc/lexer"
	"github.com/stretchr/testify/require"
)

func TestLex(t *testing.T) {
	result := lexer.Lex([]rune("(+ 1 2 3 4)"))
	require.NotNil(t, result)
	require.Equal(t, []*lexer.Token{}, result)
}
