package ast

import (
	"testing"
	"zero/token"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "hello"},
					Value: "hello",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "world"},
					Value: "world",
				},
			},
		},
	}

	if program.String() != "let hello = world;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
