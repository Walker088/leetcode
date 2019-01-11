package roman2Int

import "testing"

func Test_romanToInt(t *testing.T) {
	var input_1, output_1 = "LVIII", 58
	if res := romanToInt(input_1); res != output_1 {
		t.Error("input string:", input_1, "output string should be:", output_1, "real output is:", res)
	}
}
