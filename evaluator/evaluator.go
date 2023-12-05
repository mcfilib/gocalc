package evaluator

import (
	"strconv"

	"github.com/mcfilib/gocalc/parser"
)

// Eval reduces the AST to a value
func eval(node *parser.Node) int {
	var result int

	if node.Atom != nil {
		i, err := strconv.Atoi(string(node.Atom.Value))
		if err != nil {
			panic(err)
		}

		return i
	}

	if node.List != nil {
		switch string(node.List[0].Atom.Value) {
		case "+":
			for _, arg := range node.List[1:] {
				result = result + eval(arg)
			}

			return result
		}
	}

	return result
}
