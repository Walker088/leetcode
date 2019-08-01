package longestSubStr

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	var input1, output1 = "abcdaefg", 7
	if res := lengthOfLongestSubstring(input1); res != output1 {
		t.Error("input:", input1, "output should be:", output1, "real output is:", res)
	}
}
