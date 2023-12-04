package lexer

import (
	"fmt"
	"unicode"
)

type tType uint

const (
	Parens tType = iota
	Op
	Int
)

type Token struct {
	Value []rune
	Type  tType
}

// Lex takes a slice of runes and performs tokenization
func Lex(input []rune) []*Token {
	tokens := make([]*Token, 0, len(input))

	cursor := 0
	for cursor < len(input) {
		if unicode.IsSpace(input[cursor]) {
			cursor++
			continue
		}

		token := lexParens(cursor, input)
		if token != nil {
			tokens = append(tokens, token)
			cursor++
			continue
		}

		token = lexOp(cursor, input)
		if token != nil {
			tokens = append(tokens, token)
			cursor++
			continue
		}

		token = lexInt(cursor, input)
		if token != nil {
			tokens = append(tokens, token)
			cursor++
			continue
		}

		panic(fmt.Sprintf("unknown token: %c (pos %d)", input[cursor], cursor))
	}

	return tokens
}

func lexParens(cursor int, input []rune) *Token {
	r := input[cursor]

	if r == '(' || r == ')' {
		cursor++

		return &Token{Value: []rune{r}, Type: Parens}
	}

	return nil
}

func lexOp(cursor int, input []rune) *Token {
	r := input[cursor]

	if r == '+' {
		cursor++

		return &Token{Value: []rune{r}, Type: Op}
	}

	return nil
}

func lexInt(cursor int, input []rune) *Token {
	var value []rune

	for cursor < len(input) {
		r := input[cursor]

		if unicode.IsDigit(r) {
			value = append(value, r)
			cursor++
			continue
		}

		break
	}

	if len(value) == 0 {
		return nil
	}

	return &Token{Value: value, Type: Int}
}
