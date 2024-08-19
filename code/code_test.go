// code/code_test.go

package code

import "testing"

func TestMake(t *testing.T){
	tests := []struct {
		op			Opcode
		operands 	[]int
		expected	[]byte
	}{
		{OpConstant, []int{65534}, []byte{byte(OpConstant),255,254}},
	}

	for _, tt := range tests {
		instruction := Make(tt.op, tt.operands...)

		if len(instruction) != len(tt.expected){
			t.Errorf("ERROR: La instrucción tuvo el largo erroneo. Se esperaba=%d, se obtuvo=%d", len(tt.expected), len(instruction))
		}

		for i, b := range tt.expected {
			if instruction[i] != tt.expected[i]{
				t.Errorf("ERROR: Byte erroneo en la psición %d. Se esperaba=%d, se obtuvo=%d", i, b, instruction[i])
			}
		}
	}
}

