package main

import "fmt"

type MaxHeap struct {
	array []int
}

func (h *MaxHeap) insert(child int) {
	h.array = append(h.array, child)
	h.maxHeapifyUp(len(h.array) - 1)
}

func (h *MaxHeap) maxHeapifyUp(lastIndex int) {
	for lastIndex != 0 && h.array[getParent(lastIndex)] < h.array[lastIndex] {
		h.swap(getParent(lastIndex), lastIndex)
		lastIndex = getParent(lastIndex)
	}
}

// Helpers
func getParent(child int) int {
	return (child - 1) / 2
}

func getLeftChild(parent int) int {
	return 2*parent + 1
}

func getRightChild(parent int) int {
	return 2*parent + 2
}

func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

// --- Extract Method
func (h *MaxHeap) Extract() int {

	// -- Prevent Panic
	if len(h.array) == 0 {
		fmt.Println("Cannot extract with 0 length")
		return -1
	}

	// get top element
	extracted := h.array[0]

	// take last element, and put it at the top
	lastIndex := len(h.array) - 1
	h.array[0] = h.array[lastIndex]

	// Update the Slice with getting rid of last leaf
	h.array = h.array[:lastIndex]

	// top to bottom -- compare & swap method
	h.MaxHeapifyDown(0)

	return extracted
}

func (h *MaxHeap) MaxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	leftChild, rightChild := getLeftChild(index), getRightChild(index)
	childToCompare := 0


	// loop while index has at least one child
	for leftChild <= lastIndex {
		// left child is only child
		if leftChild == lastIndex {
			childToCompare = leftChild
			// left child greater than right child
		} else if h.array[leftChild] > h.array[rightChild] {
			childToCompare = leftChild
			// right child grater than left child
		} else {
			childToCompare = rightChild
		}

		// Compare and swap process - parent is smaller than its child
		if h.array[index] < h.array[childToCompare] {
			// swap the index to go parent place
			h.swap(index, childToCompare)
			// update the index, to compare again
			index = childToCompare
			leftChild, rightChild = getLeftChild(index), getRightChild(index)
			// else - parent is larger than its child - stop heapify
		} else {
			return
		}

	}

}

func main() {
	h := &MaxHeap{}
	h.insert(5)
	h.insert(7)
	h.insert(9)
	h.insert(19)
	fmt.Println(h.array)
	fmt.Println(h.array[getLeftChild(1)])
}
