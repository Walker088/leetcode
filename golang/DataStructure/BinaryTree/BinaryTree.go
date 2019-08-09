package BinaryTree

import "fmt"

//TreeNode defines a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//NewBinaryTree returns a Binary Tree instance base on the parameters
func NewBinaryTree(params ...interface{}) *TreeNode {
	if len(params) < 0 {
		return &TreeNode{}
	}
	return &TreeNode{}
}

//GetHeight returns the height of a Binary Tree
func (bt *TreeNode) GetHeight() int {
	if bt == nil {
		return 0
	}
	lefth := bt.Left.GetHeight()
	righth := bt.Right.GetHeight()
	if lefth > righth {
		return lefth + 1
	}
	return righth + 1
}

//PrintByPostOrder prints the elements in a Binary Tree by Postorder (Left-Right-Root)
func (bt *TreeNode) PrintByPostOrder() {
	if bt == nil {
		return
	}
	bt.Left.PrintByPostOrder()
	bt.Right.PrintByPostOrder()
	fmt.Print(bt.Val, " ")
}

//PrintByPreOrder prints the elements in a Binary Tree by Preorder (Root-Left-Right)
func (bt *TreeNode) PrintByPreOrder() {
	if bt == nil {
		return
	}
	fmt.Print(bt.Val, " ")
	bt.Left.PrintByPostOrder()
	bt.Right.PrintByPostOrder()
}

//PrintByInOrder prints the elements in a Binary Tree by Inorder (Left-Root-Right)
func (bt *TreeNode) PrintByInOrder() {
	if bt == nil {
		return
	}
	bt.Left.PrintByPostOrder()
	fmt.Print(bt.Val, " ")
	bt.Right.PrintByPostOrder()
}

//PostOrderArray returns an array which follows the postorder (Left-Right-Root)
func (bt *TreeNode) PostOrderArray() []int {
	if bt == nil {
		return nil
	}
	Lst := bt.Left.PostOrderArray()
	Lst = append(Lst, bt.Right.PostOrderArray()...)
	Lst = append(Lst, bt.Val)
	return Lst
}

//PreOrderArray returns an array which follows the preorder (Root-Left-Right)
func (bt *TreeNode) PreOrderArray() []int {
	if bt == nil {
		return nil
	}
	Lst := []int{bt.Val}
	Lst = append(Lst, bt.Left.PreOrderArray()...)
	Lst = append(Lst, bt.Right.PreOrderArray()...)
	return Lst
}

//InOrderArray returns an array which follows the inorder (Left-Root-Right)
func (bt *TreeNode) InOrderArray() []int {
	if bt == nil {
		return nil
	}
	Lst := bt.Left.InOrderArray()
	Lst = append(Lst, bt.Val)
	Lst = append(Lst, bt.Right.InOrderArray()...)
	return Lst
}
