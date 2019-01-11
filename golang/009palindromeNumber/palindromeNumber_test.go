package palindrome2integer

import "testing"

func Test_isPalindrome(t *testing.T) {
	var input_1, output_1 = 121, true
	if res := isPalindrome(input_1); res != output_1 {
		t.Error("input string:", input_1, "output string should be:", output_1, "real output is:", res)
	}
}
