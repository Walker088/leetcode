package UniqueBST

/*
NumTrees returns the number of unique bst
*/
func NumTrees(n int) int {
	//fmt.Println("NumTrees")
	dpt := make([]int, n+1) //dynamic programming table
	dpt[0], dpt[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dpt[i] += dpt[j-1] * dpt[i-j]
		}
	}
	return dpt[n]
}
