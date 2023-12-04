package lexer

import (
	"unicode"
)

type tType uint

const (
	Parens tType = iota
	Op
	Int
	Space
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
		token := lexParens(cursor, input)
		if token != nil {
			tokens = append(tokens, token)

			continue
		}

		token = lexOp(cursor, input)
		if token != nil {
			tokens = append(tokens, token)

			continue
		}

		token = lexSpace(cursor, input)
		if token != nil {
			tokens = append(tokens, token)

			continue
		}

		token = lexInt(cursor, input)
		if token != nil {
			tokens = append(tokens, token)

			continue
		}

		cursor++
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

func lexSpace(cursor int, input []rune) *Token {
	r := input[cursor]

	if unicode.IsSpace(r) {
		cursor++

		return &Token{Value: []rune{r}, Type: Space}
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
