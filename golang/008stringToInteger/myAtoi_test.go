package myAtoi

import "testing"

func Test_myAtoi(t *testing.T) {

	var s1, r1 = "     -123456789", -123456789
	var s2, r2 = "4193 with words", 4193
	var s3, r3 = "words and 987", 0
	var s4, r4 = "-91283472332", -2147483648
	var s5, r5 = "3.14159", 3
	var s6, r6 = "  + 0 123", 0
	var s7, r7 = " 0000000000012345678", 12345678

	if atoi_r := myAtoi(s1); atoi_r != r1 {
		t.Error("input1:", s1, "valid output:", r1, "real output:", atoi_r)
	}
	if atoi_r := myAtoi(s2); atoi_r != r2 {
		t.Error("input3:", s2, "valid output:", r2, "real output:", atoi_r)
	}
	if atoi_r := myAtoi(s3); atoi_r != r3 {
		t.Error("input3:", s3, "valid output:", r3, "real output:", atoi_r)
	}
	if atoi_r := myAtoi(s4); atoi_r != r4 {
		t.Error("input4:", s4, "valid output:", r4, "real output:", atoi_r)
	}
	if atoi_r := myAtoi(s5); atoi_r != r5 {
		t.Error("input5:", s5, "valid output:", r5, "real output:", atoi_r)
	}
	if atoi_r := myAtoi(s6); atoi_r != r6 {
		t.Error("input6:", s6, "valid output:", r6, "real output:", atoi_r)
	}
	if atoi_r := myAtoi(s7); atoi_r != r7 {
		t.Error("input6:", s7, "valid output:", r7, "real output:", atoi_r)
	}
}
