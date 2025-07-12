package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int {1,2,2,3,3,3}
	k := 2
	solution1 := topKFrequentSort(nums, k)
	fmt.Println(solution1)

	solution2 := topKFrequentMaxHeap(nums, k)
	fmt.Println(solution2)

}

// struct MaxHeap

type Element struct {
	num int
	freq int
}

type MaxHeap struct {
	array []Element
}

// -------------

// Insert & Extract Methods

func (h *MaxHeap) Insert(node Element) {
	h.array = append(h.array, node)
	lastIndex := len(h.array) - 1

	// We can put this into a method - MaxHeapifyUp
	h.MaxHeapifyUp(lastIndex)

} 

func (h *MaxHeap) Extract() Element {

	if (len(h.array) == 0) {
		fmt.Println("Cannot Extract Empty heap")
		return Element{}
	}
	
	extracted := h.array[0]

	// take last Element and place it to the top
	lastIndex := len(h.array) - 1
	h.array[0] = h.array[lastIndex]
	h.array = h.array[:lastIndex] // Get rid of last element as well


	// Heapify from top to bottom
	h.MaxHeapifyDown(0)

	return extracted
}


// MaxHeapifyUp & MaxHeapifyDown

func (h *MaxHeap) MaxHeapifyUp(lastIndex int) {

	// lastIndex must not reach 0
	for lastIndex > 0 && h.array[lastIndex].freq > h.array[getParent(lastIndex)].freq {
		h.swap(lastIndex, getParent(lastIndex))
		lastIndex = getParent(lastIndex)
	}
}

func (h *MaxHeap) MaxHeapifyDown(firstIndex int) {
	leftChild, rightChild := getLeftChild(firstIndex), getRightChild(firstIndex)
	childToCompare := 0
	lastIndex := len(h.array) - 1
	// while firstIndex <= lastIndex, keep the loop

	for firstIndex <= lastIndex {

		// Here we decide which child to compare with
		if rightChild > firstIndex { // right child does not exist, Out Of Bound
			childToCompare = leftChild
		} else if h.array[leftChild].freq > h.array[rightChild].freq {
			childToCompare = leftChild
		} else {
			childToCompare = rightChild
		}
	
		// while loop parent < child then swap
		if h.array[firstIndex].freq < h.array[childToCompare].freq {
			h.swap(firstIndex, childToCompare)
			firstIndex = childToCompare
			leftChild, rightChild = getLeftChild(firstIndex), getRightChild(firstIndex)
		}

	}

}


// Helper Methods
func getParent(node int) int {return (node - 1) / 2}
func getLeftChild(node int) int {return 2 * node + 1}
func getRightChild(node int) int {return 2 * node + 2}
func (h *MaxHeap) swap(index1, index2 int) {
	h.array[index1], h.array[index2] = h.array[index2], h.array[index1]
}

func topKFrequentMaxHeap(nums []int, k int) []int {

	// Creating a hashMap -- number -> occurances
	hashMap := make(map[int]int)
	for _, num := range nums {
		hashMap[num]++
	}

	// Insert the Occurances into MaxHeap
	h := &MaxHeap{}
	for key, value := range hashMap {
		h.Insert(Element{key, value})
	}

	var result []int

	// Extract top k elements
	for i := 0; i < k; i++ {
		topElement := h.Extract().num

		result = append(result, topElement) // remove and return top element.num
	}

	return result
}


// Input: nums = [1,2,2,3,3,3], k = 2
// Output: [2,3]

// This solution by using Normal Sorting
func topKFrequentSort(nums []int, k int) []int {

	// Create a hashMap number -> Occurance
	hashMap := make(map[int]int)
	for _, num := range nums {
		hashMap[num]++
	}

	// convert it into a slice
	var toArray [][2]int
	for key, value := range hashMap {
		toArray = append(toArray, [2]int{key, value})
	}

	// sort by occurances
	sort.Slice(toArray, func(i, j int) bool {
		// Decsinding order
		return toArray[i][1] > toArray[j][1]
	})

	// return top k occurance
	var result []int
	for i := 0; i < k; i++ {
		result = append(result, toArray[i][0])
	}

	// fmt.Println(result)

	return result
}