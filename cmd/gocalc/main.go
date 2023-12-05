package main

import (
	"fmt"
	"os"

	"github.com/mcfilib/gocalc/evaluator"
	"github.com/mcfilib/gocalc/lexer"
	"github.com/mcfilib/gocalc/parser"
)

func main() {
	_, ast := parser.Parse(0, lexer.Lex([]rune(os.Args[1])))

	fmt.Println(evaluator.Eval(ast))
}
