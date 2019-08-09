package ConsBinaryTree

import (
	"fmt"
	"testing"
)

func Test_BuildTree(t *testing.T) {
	fmt.Println("Test_BuildTree")
	io := []int{9, 3, 15, 20, 7}
	po := []int{9, 15, 7, 20, 3}

	bt := BuildTree(io, po)
	fmt.Println("Whole Tree:")
	bt.WholeTreePrint("Root", bt.GetHeight())
}
