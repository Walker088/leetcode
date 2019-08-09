package BinaryTree

import (
	"fmt"
	"testing"
)

func Test_NewBinaryTree(t *testing.T) {
	fmt.Println("Test Binary Tree Data Structure")
	bt := NewBinaryTree()
	fmt.Println("height: ", bt.GetHeight())

	bt.PrintByInOrder()
	fmt.Println(bt.InOrderArray())

	bt.PrintByPostOrder()
	fmt.Println(bt.PostOrderArray())

	bt.PrintByPreOrder()
	fmt.Println(bt.PreOrderArray())
}
