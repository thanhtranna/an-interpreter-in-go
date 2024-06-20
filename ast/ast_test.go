package ast

import (
	"testing"

	"github.com/thanhtranna/an-interpreter-in-go/token"
)

func Test_String(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{token.Let, "let"},
				Name: &Identifier{
					Token: token.Token{token.Identifier, "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{token.Identifier, "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != "let myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
