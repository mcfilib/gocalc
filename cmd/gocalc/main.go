package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/mcfilib/gocalc/lexer"
)

func main() {
	result := lexer.Lex([]rune("(+ 1 2 3 4)"))

	spew.Dump(result)
}
