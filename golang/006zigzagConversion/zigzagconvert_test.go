package zigzag

import "testing"

func Test_convert(t *testing.T) {
	var input_1, output_1, numRows_1 = "PAYPALISHIRING", "PAHNAPLSIIGYIR", 3
	var input_2, output_2, numRows_2 = "A", "A", 4
	var input_3, output_3, numRows_3 = "ABCDE", "ABCED", 4
	if res := convert(input_1, numRows_1); res != output_1 {
		t.Error("input string:", input_1, "output string should be:", output_1, "real output is:", res)
	}
	if res := convert(input_2, numRows_2); res != output_2 {
		t.Error("input string:", input_2, "output string should be:", output_2, "real output is:", res)
	}
	if res := convert(input_3, numRows_3); res != output_3 {
		t.Error("input string:", input_3, "output string should be:", output_3, "real output is:", res)
	}
}
