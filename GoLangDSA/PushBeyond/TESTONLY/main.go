package main

import "fmt"

type NodeList struct {
	Val  int
	next *NodeList
}

func Insert(nodeList *NodeList, val int) *NodeList {
	newNode := &NodeList{Val: val}
	if nodeList == nil {
		return newNode
	} else {
		// I'm Taking the prev list, and chain it with new one
		newNode.next = nodeList
		return newNode
	}
}

func PrintNode(nodeList *NodeList) {
	for nodeList.next != nil {
		fmt.Print(nodeList.Val, " -> ")
		nodeList = nodeList.next
	}
	fmt.Println("nil")
}

func main() {
	nodeList := &NodeList{}
	nodeList = Insert(nodeList, 1)
	nodeList = Insert(nodeList, 2)
	nodeList = Insert(nodeList, 3)
	nodeList = Insert(nodeList, 4)
	PrintNode(nodeList)
}

