package lexer_test

import (
	"testing"

	"github.com/mcfilib/gocalc/lexer"
	"github.com/stretchr/testify/require"
)

func TestLex(t *testing.T) {
	lexer.Lex()

	require.Nil(t, nil)
}