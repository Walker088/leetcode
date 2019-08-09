package PrintBinaryTree

import (
	"math"
	"strconv"

	bt "github.com/Walker088/leetcode/golang/DataStructure/BinaryTree"
)

func updateValues(n *bt.TreeNode, array [][]string, ht int, st int, ed int) {
	if n == nil {
		return
	}
	mid := (st + ed) / 2
	array[ht][mid] = strconv.Itoa(n.Val)
	updateValues(n.Left, array, ht+1, st, mid-1)
	updateValues(n.Right, array, ht+1, mid+1, ed)
}

func getHeight(n *bt.TreeNode) float64 {
	if n == nil {
		return 0
	}
	lefth := getHeight(n.Left)
	righth := getHeight(n.Right)
	if lefth > righth {
		return float64(lefth + 1)
	}
	return float64(righth + 1)
}

//PrintTree returns a 2D array of the given binay tree
func PrintTree(root *bt.TreeNode) [][]string {
	if root == nil {
		return append([][]string{}, []string{"None"})
	}
	ht := getHeight(root)
	wd := int(math.Pow(2, ht)) - 1
	res := [][]string{}
	for i := 0; i < int(ht); i++ {
		tmp := []string{}
		for j := 0; j < wd; j++ {
			tmp = append(tmp, "")
		}
		res = append(res, tmp)
	}
	updateValues(root, res, 0, 0, wd)
	return res
}
