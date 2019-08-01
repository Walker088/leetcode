package longestSubStr

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
Suppose the chart set is ASCll, then we can use a [128]int array to store the whole charactors
Then by the slide window concept, go through the input string with two pointer: i, j.
When ever the s[j] equals to s[i] then reset the start point at j and
*/
func lengthOfLongestSubstring(s string) int {
	res, sp := 0, 0     // res: result, sp: start point
	table := [128]int{} // charset table
	for i := range table {
		table[i] = -1
	}
	for ep, c := range s { // ep: end point
		if table[c] >= sp { // when s[ep] is repeated on right hand side
			sp = table[c] + 1
		}
		table[c] = ep
		res = max(res, ep-sp+1)
	}
	return res
}
