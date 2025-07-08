package main

import "fmt"

type MaxHeap struct {
	array []int
}

func (h *MaxHeap) Insert(number int) {
	h.array = append(h.array, number)
	h.MaxHeapifyUp(len(h.array) - 1)
}

// Heapify Methods
func (h *MaxHeap) MaxHeapifyUp(lastIndex int) {
	// lastIndex > 0 and the child is greater than its parent: then swap
	for lastIndex > 0 && h.array[lastIndex] > h.array[getParent(lastIndex)] {
		h.swap(lastIndex, getParent(lastIndex))
		lastIndex = getParent(lastIndex)
	}
}

func (h *MaxHeap) MaxHeapifyDown(firstIndex int) {

	// Child to compare
	childToCompare := 0
	leftChild, rightChild := getLeftChild(firstIndex), getRightChild(firstIndex)
	lastIndex := len(h.array) - 1

		// In case there is only left leaf
	if rightChild > lastIndex {
		childToCompare = leftChild
		// In case left leaf is greater than right leaf
	} else if h.array[leftChild] > h.array[rightChild] {
		childToCompare = leftChild
		// In case right leaf is greater than left leaf
	} else {
		childToCompare = rightChild
	}

		// Then We can Swap
	if h.array[firstIndex] < h.array[childToCompare] {
		h.swap(firstIndex, childToCompare)
		firstIndex = childToCompare
		leftChild, rightChild = getLeftChild(firstIndex), getRightChild(firstIndex)
	} else {
		return
	}
	
}

// Helpers Methods
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

// Extract Method
func (h *MaxHeap) Extract() int {
	extracted := h.array[0]
	lastIndex := len(h.array) - 1

	if len(h.array) == 0 {
		fmt.Println("Can't Extract empty heap")
		return -1
	}

	// take the lastElement and place it to the top
	h.array[0] = h.array[lastIndex]

	// Update the array, getting rid of last element
	h.array = h.array[:lastIndex]

	// Rearange the tree from top to bottom
	h.MaxHeapifyDown(0)

	return extracted
}

func main() {}
