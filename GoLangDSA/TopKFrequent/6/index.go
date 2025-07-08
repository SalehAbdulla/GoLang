package main

import "fmt"

// Step 1: Element Struct 
type Element struct {
	freq int
	num  int
}

type MaxHeap struct {
	array []Element
}

// Step 2: Insert & Extract Methods

func (h *MaxHeap) Insert(element Element) {
	h.array = append(h.array, element)
	h.MaxHeapifyUp(len(h.array) - 1)
}

func (h *MaxHeap) Extract() Element {
	if len(h.array) == 0 {
		fmt.Println("Cannot extract from empty heap")
		return Element{}
	}

	extracted := h.array[0]
	last := len(h.array) - 1
	h.array[0] = h.array[last]
	h.array = h.array[:last]
	h.MaxHeapifyDown(0)

	return extracted
}

// Step 3: MaxHeapifyUp & Down

func (h *MaxHeap) MaxHeapifyUp(index int) {
	for index > 0 && h.array[index].freq > h.array[getParent(index)].freq {
		h.swap(index, getParent(index))
		index = getParent(index)
	}
}

func (h *MaxHeap) MaxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	childToCompare := 0
	leftChild, rightChild := getLeftChild(index), getRightChild(index)

	for leftChild <= lastIndex {

		if rightChild > lastIndex {
			childToCompare = leftChild
		} else if h.array[leftChild].freq > h.array[rightChild].freq {
			childToCompare = leftChild
		} else {
			childToCompare = rightChild
		}

		if h.array[index].freq < h.array[childToCompare].freq {
			h.swap(index, childToCompare)
			index = childToCompare
			leftChild = getLeftChild(index)
			rightChild = getRightChild(index)
		} else {
			return
		}
	}
}

// Step 4: Helper methods

func getParent(index int) int       { return (index - 1) / 2 }
func getLeftChild(index int) int    { return 2*index + 1 }
func getRightChild(index int) int   { return 2*index + 2 }
func (h *MaxHeap) swap(i1, i2 int)  { h.array[i1], h.array[i2] = h.array[i2], h.array[i1] }

// Step 5: TopKFrequentElement

func TopKFrequent(nums []int, k int) []int {
	// 1. Count frequencies
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	// 2. Build max heap using our custom struct
	h := &MaxHeap{}
	for num, freq := range freqMap {
		h.Insert(Element{freq: freq, num: num})
	}

	// 3. Extract top k frequent elements
	result := []int{}
	for i := 0; i < k; i++ {
		result = append(result, h.Extract().num)
	}

	return result
}

// Test
func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	fmt.Println(TopKFrequent(nums, k)) // Output: [1 2]
}
