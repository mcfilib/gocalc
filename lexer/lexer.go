package lexer

type tType uint

const (
	Parens tType = iota
	Int
	Op
)

type Token struct {
	Value string
	Type  tType
}

func Lex(input []rune) []*Token {
	return []*Token{}
}
