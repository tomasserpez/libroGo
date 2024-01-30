// ast/ast_test.go
package ast

import (
	"testing"
	"tomprueba/token"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "miVariable"},
					Value: "miVariable",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "miOtraVariable"},
					Value: "miOtraVariable",
				},
			},
		},
	}
	if program.String() != "let miVariable = miOtraVariable;" {
		t.Errorf("program.String() error. Se obtuvo=%q", program.String())
	}
}
