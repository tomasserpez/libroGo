// evaluator/builtins.go

package evaluator

import "tomlang/object"

var builtins = map[string]*object.Builtin{
	"len": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("Cantidad erronea de argumentos. Se obtuvo=%d, Se esperaba=1", len(args))
			}
			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return newError("Argumento para 'len' no soportado, se obtuvo=%s", args[0].Type())
			}
		},
	},
}
