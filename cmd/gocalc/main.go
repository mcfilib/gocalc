package main

import (
	"fmt"

	"github.com/mcfilib/gocalc/lexer"
)

func main() {
	result := lexer.Lex([]rune("(+ 1 2 3 4)"))

	fmt.Println(result)
}
