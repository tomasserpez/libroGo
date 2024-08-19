// compiler/compiler.go

package compiler
import (
	"tomlang/ast"
	"tomlang/code"
	"tomlang/object"
)

type Compiler struct {
	instruction code.Instructions
	constants		[]objects.Object
}

func New() *Compiler{
	return &Compiler{
		instructions:	code.Instructions{},
		constants:		[]object.Object{},
	}
}

func (c *Compiler) Compile(node ast.Node) error{
	return nil
}

func (c *Compiler) Bytecode() *Bytecode {
	return &Bytecode{
		Instructions:	c.instruction,
		Constants:		c.constants,
	}
}

type Bytecode struct {
	Instructions	code.Instructions
	Constants		[]object.Object
}
