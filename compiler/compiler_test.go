// compiler/compiler_test.go

package compiler

import (
	"fmt"
	"testing"
	"tomlang/ast"
	"tomlang/code"
	"tomlang/lexer"
	"tomlang/object"
	"tomlang/parser"
)

func parse(input string) *ast.Program {
	l := lexer.New(input)
	p := parser.New(l)
	return p.ParseProgram()
}

type compilerTestCase struct {
	input                string
	expectedConstants    []interface{}
	expectedInstructions []code.Instructions
}

func TestIntegerArithmetic(t *testing.T) {
	tests := []compilerTestCase{
		{
			input:             "1 + 2",
			expectedConstants: []interface{}{1, 2},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OpConstant, 1),
				code.Make(code.OpAdd),
			},
		},
	}
	runCompilerTests(t, tests)
}

func runCompilerTests(t *testing.T, tests []compilerTestCase) {
	t.Helper()

	for _, tt := range tests {
		program := parse(tt.input)

		compiler := New()
		err := compiler.Compile(program)
		if err != nil {
			t.Fatalf("ERROR: Error de compilacion: %s", err)
		}

		bytecode := compiler.Bytecode()

		err = testInstructions(tt.expectedInstructions, bytecode.Instructions)
		if err != nil {
			t.Fatalf("ERROR: error en el testInstructions: %s", err)
		}

		err = testConstants(t, tt.expectedConstants, bytecode.Constants)
		if err != nil {
			t.Fatalf("ERROR: error en el testConstants: %s", err)
		}
	}
}

func testInstructions(
	expected []code.Instructions,
	actual code.Instructions,
) error {
	concatted := concatInstructions(expected)

	if len(actual) != len(concatted) {
		return fmt.Errorf("ERROR: instrucción con largo erroneo. \nSe esperaba=%q\nSe obtuvo=%q",
			concatted, actual)
	}

	for i, ins := range concatted {
		if actual[i] != ins {
			return fmt.Errorf("ERROR: instruccion erronea en %d. \nSe esperaba=%q\nSe obtuvo=%q", i, concatted, actual)
		}
	}
	return nil
}

func concatInstructions(s []code.Instructions) code.Instructions {
	out := code.Instructions{}

	for _, ins := range s {
		out = append(out, ins...)
	}
	return out
}

func testConstants(
	t *testing.T,
	expected []interface{},
	actual []object.Object,
) error {
	if len(expected) != len(actual) {
		return fmt.Errorf("ERROR: Numero erroneo de constantes. Se obtuvo=%d, Se esperaba=%d", len(actual), len(expected))
	}

	for i, constant := range expected {
		switch constant := constant.(type) {
		case int:
			err := testIntegerObject(int64(constant), actual[i])
			if err != nil {
				return fmt.Errorf("ERROR: constant %d - falló testIntegerObject: %s", i, err)
			}
		}
	}
	return nil
}

func testIntegerObject(expected int64, actual object.Object) error {
	result, ok := actual.(*object.Integer)
	if !ok {
		return fmt.Errorf("ERROR: El objeto no es un integer. Se obtuvo=%T (%+v", actual, actual)
	}
	if result.Value != expected {
		return fmt.Errorf("ERROR: el objeto tiene el valor erroneo. Se obtuvo=%d, Se esperaba=%d", result.Value, expected)
	}

	return nil
}
