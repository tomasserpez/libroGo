// vm/vm_test.go

package vm

import (
	"fmt"
	"testing"
	"tomlang/ast"
	"tomlang/lexer"
	"tomlang/object"
	"tomlang/parser"
)

func parse(input string) *ast.Program {
	l := lexer.New(input)
	p := parser.New(l)
	return p.ParseProgram()
}

func testIntegerObject(expected int64, actual object.Object) error {
	result, ok := actual.(*object.Integer)
	if !ok {
		return fmt.Errorf("ERROR: El objeto no es Integer. Se obtuvo=%T (%+v)", actual, actual)
	}

	if result.Value != expected {
		return fmt.Errorf("ERROR: El objeto tiene valor erroneo. Se obtuvo=%d, se esperaba=%d", result.Value, expected)
	}

	return nil
}

type vmTestCase struct {
	input    string
	expected interface{}
}

func runVmTests(t *testing.T, tests []vmTestCase) {
	t.Helper()

	for _, tt := range tests {
		program := parse(tt.input)
	}
}
