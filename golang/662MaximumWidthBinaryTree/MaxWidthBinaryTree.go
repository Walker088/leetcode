package MaxWidthBinaryTree

import bt "github.com/Walker088/leetcode/golang/DataStructure/BinaryTree"

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func dfs(n *bt.TreeNode, id int, d int, lefts *[]int) int { // d: depth
	if n == nil {
		return 0
	}
	if d == len(*lefts) {
		*lefts = append(*lefts, id)
	}
	//fmt.Println("tree, id, depth, lefts, current res: ", n, id, d, lefts, id-(*lefts)[d]+1)
	return max(id-(*lefts)[d]+1, max(dfs(n.Left, id*2, d+1, lefts), dfs(n.Right, id*2+1, d+1, lefts)))
}

func widthOfBinaryTree(root *bt.TreeNode) int {
	var lefts []int
	return dfs(root, 1, 0, &lefts)
}
