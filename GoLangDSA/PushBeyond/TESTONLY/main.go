package main

import "fmt"


type NodeList struct {
	val int // The only data holds
	next *NodeList // for chaining
}

func InsertBeginning(nodeList *NodeList, val int) *NodeList {
	newNode := &NodeList{val:val} // like new node() in Java; Go holds location explicitly
	if nodeList == nil {
		return newNode
	} else {
		newNode.next = nodeList
		return newNode
	}
}


func PrintNode(nodeList *NodeList) {
	for nodeList != nil {
		fmt.Print(nodeList.val, " -> ")
		nodeList = nodeList.next
	}
	fmt.Println("nil")
}

func main(){

	var nodeList *NodeList = nil
	nodeList = InsertBeginning(nodeList, 9)
	nodeList = InsertBeginning(nodeList, 8)
	nodeList = InsertBeginning(nodeList, 7)
	nodeList = InsertBeginning(nodeList, 6)
	PrintNode(nodeList)



}