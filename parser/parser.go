package parser

import (
	"github.com/mcfilib/gocalc/lexer"
)

type Node struct {
	Atom *lexer.Token
	List []*Node
}

// Parse takes a list of tokens and returns an AST
func Parse(cursor int, input []*lexer.Token) (int, []*Node) {
	var ast []*Node

	for cursor < len(input) {
		token := input[cursor]
		cursor++

		if token.Type == lexer.Parens && string(token.Value) == "(" {
			newCursor, subTree := Parse(cursor, input)

			ast = append(ast, &Node{List: subTree})
			cursor = newCursor

			continue
		}

		if token.Type == lexer.Parens && string(token.Value) == ")" {
			return cursor, ast
		}

		ast = append(ast, &Node{Atom: token})
	}

	return cursor, ast
}
