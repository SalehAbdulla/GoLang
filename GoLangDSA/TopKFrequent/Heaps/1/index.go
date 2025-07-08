package main

import "fmt"

// create MaxHeap Struct
type MaxHeap struct {
	array []int
}

// Insert Method
func (h *MaxHeap) Insert(index int) {
	h.array = append(h.array, index)
	h.MaxHeapifyUp(len(h.array) - 1)
}

// MaxHeapifyUp & MaxHeapifyDown
func (h *MaxHeap) MaxHeapifyUp(lastIndex int) {

	for lastIndex > 0 && h.array[lastIndex] > h.array[getParent(lastIndex)] {
		h.swap(lastIndex, getParent(lastIndex))
		lastIndex = getParent(lastIndex)
	}

}

func (h *MaxHeap) MaxHeapifyDown(firstIndex int) {

	lastIndex := len(h.array) - 1
	// Keep track of left and right child
	leftChild, rightChild := getLeftChild(firstIndex), getRightChild(firstIndex)

	childToCompare := 0

	// while leftChild <= lastIndex 
	for leftChild <= lastIndex {
		// Decide which to compare with
		// if rightChild > lastIndex means there is no right element
		if rightChild > lastIndex {
			childToCompare = leftChild
		} else if h.array[leftChild] > h.array[rightChild] {
			childToCompare = leftChild
		} else {
			childToCompare = rightChild
		}

		// then we can swap
		if h.array[firstIndex] < h.array[childToCompare] {
			h.swap(firstIndex, childToCompare)
			firstIndex = childToCompare
			leftChild, rightChild = getLeftChild(firstIndex), getRightChild(firstIndex)
		} else {
			// else its on the correct place - return;
			return
		}

	}

}

// Helper method for heapify methods
func getParent(index int) int {
	return (index - 1) / 2
}
func getLeftChild(index int) int {
	return 2*index + 1
}
func getRightChild(index int) int {
	return 2*index + 2
}
func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

// Extract Method - return first element and delete it from the top

// We can only delete top element
// last element will take top element place
// from top to bottom, we will compare its child, and swap

func (h *MaxHeap) extract() int {

	if len(h.array) == 0 {
		fmt.Println("Cannot extract from empty heap")
		return -1
	}

	extracted := h.array[0]
	l := len(h.array) - 1

	h.array[0] = h.array[l]
	h.array = h.array[:l] // Update the slice, get rid of lastElement

	h.MaxHeapifyDown(0)

	return extracted
}

func main() {
	h := &MaxHeap{}
	h.Insert(9)
	h.Insert(19)
	h.Insert(29)
	h.Insert(39)
	h.Insert(49)
	h.Insert(59)
	fmt.Println(h.array)
	fmt.Println(h.extract())
	fmt.Println(h.array)
}

