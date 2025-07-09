package main

import "fmt"

// Input: nums = [1,2,2,3,3,3], k = 2
// Output: [2,3]

// 1) Create a HashMap number -> Occurances
// 2) Convert it into Max-Heap
// 3) Extract top k number

// ------------- MaxHeap -------------

// Because the structure of the hashMap is map[int]int, so its gonna be like

// Structure
type Element struct {
	num  int
	freq int
}

type MaxHeap struct {
	array []Element
}

// Insert & Extract Methods

func (h *MaxHeap) Insert(node Element) {
	h.array = append(h.array, node)
	h.MaxHeapifyUp(len(h.array) - 1) // Because its gonna be added at the bottom, we need to make the adjusments from the bottom to the top
}

// This Method gonna extract and return top element, and get rid of it
func (h *MaxHeap) Extract() Element {
	extracted := h.array[0]

	// Avoid index out of Range Panic
	if len(h.array) == 0 {
		fmt.Println("Cannot Extract from empty Heap")
		return Element{}
	}

	// Take lastNode place it to the top
	lastIndex := len(h.array) - 1
	h.array[0] = h.array[lastIndex]

	// Update Binary Tree
	h.array = h.array[:lastIndex]

	// Heapify from top to bottom - Compare with childs, then keep swaping
	h.MaxHeapifyDown(0)

	return extracted
}

// MaxHeapifyUp & MaxHeapifyDown
func (h *MaxHeap) MaxHeapifyUp(lastIndex int) {
	// MaxHeap=parentBigger, so while child is bigger keep swapping
	// Gonna keep swap any child thats bigger than its parent
	for h.array[lastIndex].freq > h.array[getParent(lastIndex)].freq {
		h.swap(lastIndex, getParent(lastIndex))
		lastIndex = getParent(lastIndex)
	}

}

func (h *MaxHeap) MaxHeapifyDown(firstIndex int) {

	leftChild, rightChild := getLeftChild(firstIndex), getRightChild(firstIndex)
	lastIndex := len(h.array) - 1
	childToCompare := 0

	// which child to compare with ?, looping from topNode till lastNode, swap any vaiolants

	for firstIndex <= lastIndex {

		// if no children, then break
		if (leftChild > lastIndex) {
			break
		}

		// In case there is only left leaf (Out Of Bound)
		if rightChild > lastIndex {
			childToCompare = leftChild
		} else if h.array[leftChild].freq > h.array[rightChild].freq {
			childToCompare = leftChild
		} else {
			childToCompare = rightChild
		}

		// Keep Swapping parents with childs from top to bottom
		if h.array[childToCompare].freq > h.array[firstIndex].freq {
			h.swap(childToCompare, firstIndex)
			firstIndex = childToCompare
			leftChild, rightChild = getLeftChild(firstIndex), getRightChild(firstIndex)
		} else {
			return
		}
	}

}

// Helper methods
func getParent(index int) int     { return (index - 1) / 2 }
func getLeftChild(index int) int  { return 2*index + 1 }
func getRightChild(index int) int { return 2*index + 2 }
func (h *MaxHeap) swap(index1, index2 int) {
	h.array[index1], h.array[index2] = h.array[index2], h.array[index1]
}

func topKFrequent(nums []int, k int) []int {
	// number -> Occurance
	hashMap := make(map[int]int)
	for _, num := range nums {
		hashMap[num]++
	}

	// Convert into MaxHeap
	h := &MaxHeap{}
	for num, freq := range hashMap {
		h.Insert(Element{num: num, freq: freq})
	}

	// Extract top k num
	var result []int
	for i := 0; i < k; i++ {
		result = append(result, h.Extract().num)
	}

	return result
}

func main() {
	nums := []int{1, 2, 2, 3, 3, 3}
	k := 2

	solution := topKFrequent(nums, k)
	fmt.Println(solution)

}
