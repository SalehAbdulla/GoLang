package main

import (
	"fmt"
)

func main() {
	nums := []int {1,2,2,3,3,3}
	k := 2
	solution := topKFrequent(nums, k)
	fmt.Println(solution)
}

// ----- MaxHeap Struct
type Element struct {
	num  int
	freq int
}

type MaxHeap struct {
	array []Element
}

// Insert & Extract
func (h *MaxHeap) Insert(node Element) {
	// MaxHeap means parents are bigger than their childs
	// The insertion process is to add it at the end of the tree
	// then heapifyUp

	h.array = append(h.array, node)  // add
	h.MaxHeapifyUp(len(h.array) - 1) // HeapifyUp

}

func (h *MaxHeap) Extract() Element {
	
	if len(h.array) == 0 {
		fmt.Println("Cannot Extract from empty Heap")
		return Element{}
	}
	
	extracted := h.array[0]

	// get last leap to the top
	// then compare it with largest child
	// keep the swapping process
	lastIndex := len(h.array) - 1
	h.array[0] = h.array[lastIndex]
	h.array = h.array[:lastIndex] // last leaf not included 

	// HeapifyDown Process starts here
	h.MaxHeapifyDown(0)


	return extracted
}



// MaxHeapifyUp and MaxHeapifyDown
func (h *MaxHeap) MaxHeapifyUp(lastIndex int) {
	for lastIndex > 0 && h.array[lastIndex].freq > h.array[getParent(lastIndex)].freq {
		h.swap(lastIndex, getParent(lastIndex))
		lastIndex = getParent(lastIndex)
	}
}

func (h *MaxHeap) MaxHeapifyDown(firstIndex int) {
	leftChild, rightChild := getLeftChild(firstIndex), getRightChild(firstIndex)
	lastIndex := len(h.array) - 1
	childToCompare := 0

	for firstIndex < lastIndex {

		// if right child does not exist
		if rightChild > firstIndex {
			childToCompare = leftChild
		} else if h.array[leftChild].freq > h.array[rightChild].freq{
			childToCompare = leftChild
		} else {
			childToCompare = rightChild
		}

		// Swapping
		if h.array[firstIndex].freq < h.array[childToCompare].freq {
			h.swap(firstIndex, childToCompare)
			firstIndex = childToCompare
			leftChild, rightChild = getLeftChild(firstIndex), getRightChild(firstIndex)
		} else { // Heapify Satisfied
			return
		}
	}
}

// Helper methods
func getParent(node int) int {
	return (node - 1) / 2
}
func getLeftChild(node int) int {
	return 2*node + 1
}
func getRightChild(node int) int {
	return 2*node + 2
}
func (h *MaxHeap) swap(node1, node2 int) {
	h.array[node1], h.array[node2] = h.array[node2], h.array[node1]
}


func topKFrequent(nums []int, k int) []int {
	
	// hashMap 
	hashMap := make(map[int]int)
	for _, num := range nums {
		hashMap[num]++
	}
	// Convert it into maxHeap
	h := MaxHeap{}
	for num, freq := range hashMap {
		h.Insert(Element{num, freq})
	}

	// get top k freq
	result := []int{}
	for i:= 0; i < k; i++ {
		result = append(result, h.Extract().num)
	}

	// return result
	return result
}