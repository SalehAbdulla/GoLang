package main

import "fmt"

type node struct {
	left, right *node
	data        int
}

func insert(root *node, ele int) *node {
	if root == nil {
		return &node{data: ele}
	}

	if ele < root.data {
		root.left = insert(root.left, ele)
	} else {
		root.right = insert(root.right, ele)
	}

	return root
}

func PrintTree(root *node) {
	if root == nil {
		return
	}
	fmt.Println(root.data)
	PrintTree(root.left)
	PrintTree(root.right)
}

func main() {
	var root *node
	insert(root, 4)
	insert(root, 2)
	insert(root, 6)
	insert(root, 1)
	insert(root, 3)
	insert(root, 5)
	insert(root, 7)
	PrintTree(root)
}
