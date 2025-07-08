package main

import (
	"fmt"
)

type Element struct {
	num int
	freq int
}

type MaxHeap struct {
	array[]Element
}

// Insert and Extract methods
func (h *MaxHeap) Insert(node Element) {
	h.array = append(h.array, node)
	h.MaxHeapifyUp(len(h.array) - 1)
}

func (h *MaxHeap) Extract() Element {
	extracted := h.array[0]
	lastIndex := len(h.array) - 1

	if (len(h.array) == 0) {
		fmt.Println("Cannot Extract from Empty Heap")
		return Element{}
	}

	// take last element, put it at the top
	h.array[0] = h.array[lastIndex]
	// Update the array, get rid of last element
	h.array = h.array[:lastIndex]

	h.MaxHeapifyDown(0)

	return extracted
}

// MaxHeapifyUp and MaxHeapifyDown

func (h *MaxHeap) MaxHeapifyUp(lastIndex int) {
	// Compare with its parent and then swap
	for h.array[lastIndex].freq > h.array[getParent(lastIndex)].freq {
		h.swap(lastIndex, getParent(lastIndex))
		lastIndex = getParent(lastIndex)
	}
}

func (h *MaxHeap) MaxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1

	for {
		leftChild, rightChild := getLeftChild(index), getRightChild(index)
		largest := index

		if leftChild <= lastIndex && h.array[leftChild].freq > h.array[largest].freq {
			largest = leftChild
		}
		if rightChild <= lastIndex && h.array[rightChild].freq > h.array[largest].freq {
			largest = rightChild
		}

		if largest == index {
			// Heap property is satisfied
			break
		}

		h.swap(index, largest)
		index = largest
	}
}


// Helper methods
func getParent(index int) int { return (index - 1) /2 }
func getLeftChild(index int) int { return 2 * index + 1 }
func getRightChild(index int) int { return 2 * index + 2 }
func (h *MaxHeap) swap(i1, i2 int) { h.array[i1], h.array[i2] = h.array[i2], h.array[i1] }



func TopKFrequent(nums []int, k int) []int{
	// Hash make number -> Occurance
	hashMap := make(map[int]int)
	for _, num := range nums {
		hashMap[num]++
	}
	// Convert it into MaxHeap and add the element
	h := &MaxHeap{}
	for num, freq := range hashMap {
		h.Insert(Element{freq: freq, num: num})
	}

	// Extract top k num elements
	result := []int{}
	for i := 0; i < k; i++ {
		result = append(result, h.Extract().num)
	}

	return result
}




func main(){
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	fmt.Println(TopKFrequent(nums, k)) // Output: [1 2]
}

