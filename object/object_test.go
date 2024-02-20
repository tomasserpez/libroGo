// object/object_test.go

package object

import "testing"

func TestStringHashKey(t *testing.T) {
	hello1 := &String{Value: "Hello World"}
	hello2 := &String{Value: "Hello World"}
	diff1 := &String{Value: "Mi nombre es Tomas"}
	diff2 := &String{Value: "Mi nombre es Tomas"}
	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("Cadenas con el mismo contenido y diferentes claves")
	}
	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("Cadenas con el mismo contenido y diferentes claves")
	}
	if hello1.HashKey() == diff1.HashKey() {
		t.Errorf("Cadenas con diferente contenido y misma clave")
	}
}
