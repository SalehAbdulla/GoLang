package main

import "fmt"

// Top k Frequent Element

// Input: nums = [1,2,2,3,3,3], k = 2
// Output: [2,3]
// In this solution we are gonna use MaxHeap instead of MinHeap for Learning purposes

// the structure of saving the hashMap int -> int, thus
type Element struct {
	num  int
	freq int
}

type MaxHeap struct {
	array []Element
}

// --- Insert and Extract
func (h *MaxHeap) Insert(node Element) {
	h.array = append(h.array, node)
	lastIndex := len(h.array) - 1
	h.MaxHeapifyUp(lastIndex) // Max Heapify Adjustment
}

func (h *MaxHeap) Extract() Element {
	
	// Avoid Empty Array Extract Panic
	if len(h.array) == 0 {
		fmt.Println("Cannot extract empty array")
		return Element{}
	}
	
	extracted := h.array[0]
	// Jump, get rid and update
	lastIndex := len(h.array) - 1
	h.array[0] = h.array[lastIndex]
	h.array = h.array[:lastIndex]

	// Heapify Process from top to bottom -- compare then swap
	h.MaxHeapifyDown(0)

	return extracted
}

// --- MaxHeapifyUp & MaxHeapifyDown
func (h *MaxHeap) MaxHeapifyUp(lastIndex int) {
	// while the child.freq is greater than its parent.freq -- keep swaping
	for len(h.array) > 0 && h.array[lastIndex].freq > h.array[getParent(lastIndex)].freq {
		h.swap(lastIndex, getParent(lastIndex)) // Swap
		lastIndex = getParent(lastIndex)        // Update
	}
}

func (h *MaxHeap) MaxHeapifyDown(firstIndex int) {
	leftChild, rightChild := getLeftChild(firstIndex), getRightChild(firstIndex)
	lastIndex := len(h.array) - 1
	childToCompare := 0
	// compare curr node with its bigger child
	// loop from the top to the bottom

	for firstIndex <= lastIndex {
		// No Child case - Break
		if leftChild > lastIndex {
			break
		}

		// No Right Child - proceed with left child
		if rightChild > lastIndex {
			childToCompare = leftChild
		} else if h.array[leftChild].freq > h.array[rightChild].freq {
			childToCompare = leftChild
		} else {
			childToCompare = rightChild
		}

		// Swapping process
		if h.array[firstIndex].freq < h.array[childToCompare].freq {
			h.swap(firstIndex, childToCompare)
			firstIndex = childToCompare
			leftChild, rightChild = getLeftChild(firstIndex), getRightChild(firstIndex)
		} else {
			// Heapify Satisfied - exit function
			return
		}

	}

}

// --- Helper Methods
func getParent(index int) int     { return (index - 1) / 2 }
func getLeftChild(index int) int  { return 2*index + 1 }
func getRightChild(index int) int { return 2*index + 2 }
func (h *MaxHeap) swap(index1, index2 int) {
	h.array[index1], h.array[index2] = h.array[index2], h.array[index1]
}

func topKFrequent(nums []int, k int) []int {
	// invidable hashMap
	hashMap := make(map[int]int)
	for _, num := range nums {
		hashMap[num]++
	}

	// Turn it into MaxHeap
	h := &MaxHeap{}
	for num, freq := range hashMap {
		h.Insert(Element{num: num, freq: freq})
	}

	// .pop top k num elements
	var result []int
	for i := 0; i < k; i++ {
		result = append(result, h.Extract().num)
	}

	return result
}


func main() {
	nums := []int {1,2,2,3,3,3}
	k := 2
	solution := topKFrequent(nums, k)
	fmt.Println(solution)
}