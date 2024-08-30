// code/code_test.go

package code

import "testing"

func TestMake(t *testing.T) {
	tests := []struct {
		op       Opcode
		operands []int
		expected []byte
	}{
		{OpConstant, []int{65534}, []byte{byte(OpConstant), 255, 254}},
		{OpAdd, []int{}, []byte{byte(OpAdd)}},
	}

	for _, tt := range tests {
		instruction := Make(tt.op, tt.operands...)

		if len(instruction) != len(tt.expected) {
			t.Errorf("ERROR: La instrucción tuvo el largo erroneo. Se esperaba=%d, se obtuvo=%d", len(tt.expected), len(instruction))
		}

		for i, b := range tt.expected {
			if instruction[i] != tt.expected[i] {
				t.Errorf("ERROR: Byte erroneo en la psición %d. Se esperaba=%d, se obtuvo=%d", i, b, instruction[i])
			}
		}
	}
}

func TestInstructionString(t *testing.T) {
	instructions := []Instructions{
		//Make(OpConstant, 1),
		Make(OpAdd),
		Make(OpConstant, 2),
		Make(OpConstant, 65535),
	}

	expected := `0000 OpAdd
0001 OpConstant 2
0004 OpConstant 65535
`

	concatted := Instructions{}
	for _, ins := range instructions {
		concatted = append(concatted, ins...)
	}
	if concatted.String() != expected {
		t.Errorf("ERROR: las instrucciones estan formateadas de forma erronea. \nSe esperaba=%q\nSe obtuvo=%q", expected, concatted.String())
	}
}

func TestReadOperands(t *testing.T) {
	tests := []struct {
		op        Opcode
		operands  []int
		bytesRead int
	}{
		{OpConstant, []int{65535}, 2},
	}

	for _, tt := range tests {
		instruction := Make(tt.op, tt.operands...)

		def, err := Lookup(byte(tt.op))
		if err != nil {
			t.Fatalf("ERROR: definición no encontrada: %q\n", err)
		}

		operandsRead, n := ReadOperands(def, instruction[1:])
		if n != tt.bytesRead {
			t.Fatalf("ERROR: n erroneo. Se esperaba=%d, Se obtuvo=%d", tt.bytesRead, n)
		}

		for i, want := range tt.operands {
			if operandsRead[i] != want {
				t.Errorf("ERROR: Operando erroneo. Se esperaba=%d, Se obtuvo=%d", want, operandsRead[i])
			}
		}
	}
}
