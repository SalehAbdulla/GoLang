package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 2, 3, 3, 3}
	k := 2
	solution := topKFrequentElement(nums, k)
	fmt.Println(solution)
}

//Input: nums = [1,2,2,3,3,3], k = 2
// Output: [2,3]

// MinHeap Struct
type Element struct {
	num  int
	freq int
}

type MinHeap struct {
	array []Element
	k     int
}

// ---- Insert & Extract

func (h *MinHeap) Insert(node Element) {
	h.array = append(h.array, node)
	h.MinHeapifyUp(len(h.array) - 1)
	// This will save memory, prevent unecessary saving
	if len(h.array) > h.k {
		h.Extract()
	}
}

func (h *MinHeap) Extract() Element {
	extracted := h.array[0]

	// Remove top element
	lastIndex := len(h.array) - 1
	// take lastElement, place it into the first
	h.array[0] = h.array[lastIndex]
	// Update the entire Heap
	h.array = h.array[:lastIndex]

	// Heapify process from top to bottom
	h.MinHeapifyDown(0)

	return extracted
}

// --- MinHeapifyUp() & MinHeapifyDown
func (h *MinHeap) MinHeapifyUp(lastIndex int) {
	// MinHeap - parent is smaller than its childs
	// child is smaller than its parent
	for lastIndex > 0 && h.array[lastIndex].freq < h.array[getParent(lastIndex)].freq {
		h.swap(lastIndex, getParent(lastIndex))
		lastIndex = getParent(lastIndex)
	}
}

func (h *MinHeap) MinHeapifyDown(index int) {
	lastIndex := len(h.array) - 1

	for {
		left := getLeftChild(index)
		right := getRightChild(index)
		smallest := index

		if left <= lastIndex && h.array[left].freq < h.array[smallest].freq {
			smallest = left
		}
		if right <= lastIndex && h.array[right].freq < h.array[smallest].freq {
			smallest = right
		}

		if smallest != index {
			h.swap(index, smallest)
			index = smallest
		} else {
			break
		}
	}
}

// --- Helper Methods
func getParent(node int) int     { return (node - 1) / 2 }
func getLeftChild(node int) int  { return 2*node + 1 }
func getRightChild(node int) int { return 2*node + 2 }
func (h *MinHeap) swap(index1, index2 int) {
	h.array[index1], h.array[index2] = h.array[index2], h.array[index1]
}

func topKFrequentElement(nums []int, k int) []int {
	// 1) Create a hashMap number -> Occurance
	hashMap := make(map[int]int)
	for _, num := range nums {
		hashMap[num]++
	}

	// 2) Convert it into MinHeap
	h := &MinHeap{k: k}
	for num, freq := range hashMap {
		h.Insert(Element{num: num, freq: freq})
	}

	// 3) Extract top k Element then return
	var result []int
	for i := 0; i < k; i++ {
		result = append(result, h.Extract().num)
	}

	return result
}
