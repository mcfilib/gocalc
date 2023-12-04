package parser

import "github.com/mcfilib/gocalc/lexer"

type AST struct {
	First *lexer.Token
	Rest  []*AST
}

// Parse takes a list of tokens and returns an AST
func Parse(input []*lexer.Token) *AST {
	return &AST{}
}
