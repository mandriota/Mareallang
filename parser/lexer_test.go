package parser

import (
	"fmt"
	"strings"
	"testing"
)

const TEST = `
(out (et 1e6 (_ x nl)))
`

func TestLexer_Ast(t *testing.T) {
	l := &Lexer{in: TEST}
	if v := l.Lex(); v != nil {
		t.Log(traversal("", v))
	}
}

func traversal(ind string, node *Node) string {
	var out string
	for i, el := range node.Component {
		out += fmt.Sprintf("\n%s%d) Type: %s; Literal: %v", ind, i, Tokens[el.Typ], el.Val)
		if el.Component != nil {
			out += traversal(strings.Repeat("\t", len(ind)+1), el)
		}
	}
	return out
}
