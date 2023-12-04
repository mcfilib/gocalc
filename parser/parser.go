package parser

import (
	"github.com/mcfilib/gocalc/lexer"
)

type AST struct {
	First *lexer.Token
	Rest  []*AST
}

// Parse takes a list of tokens and returns an AST
func Parse(cursor int, input []*lexer.Token) (int, *AST) {
	var ast *AST

	for cursor < len(input) {
		token := input[cursor]

		if token.Type == lexer.Parens && string(token.Value) == "(" {
			// recurse
		}

		if token.Type == lexer.Parens && string(token.Value) == ")" {
			// finish
		}
	}

	return cursor, ast
}
