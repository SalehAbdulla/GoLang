package main

import "fmt"

// Top K Frequent Element using MinHeap

// Input: nums = [1,2,2,3,3,3], k = 2
// Output: [2,3]

// hashMap number -> occurances
// Convert it into MinHeap, why? we will extract small-values and keep k element only - saving space and time
// save the result into an []result and return solution

// Create the structure, since the hashmap looks like int -> int, thus
type Element struct {
	num int
	freq int
}

type MinHeap struct {
	array []Element
	k int
}

// --- MinHeap Methods

// Insert and Extract
func (h *MinHeap) Insert(node Element) {
	h.array = append(h.array, node)
	h.MinHeapifyUp(len(h.array) - 1)
	// Make sure our heap does not exceed k size
	for (len(h.array) > h.k) {
		h.Extract()
	}
}

func (h *MinHeap) Extract() Element {
	// Avoid Panic
	if len(h.array) == 0 {
		fmt.Println("Cannot Extract empty Heap")
		return Element{}
	}
	extracted := h.array[0]

	// Jump, update, heapifyDown
	h.array[0] = h.array[len(h.array) - 1] // lastNode becomes firstNode
	h.array = h.array[:len(h.array) - 1]
	h.MinHeapifyDown(0)

	return extracted

}

// MinHeapifyUp & MinHeapifyDown
func (h *MinHeap) MinHeapifyUp(lastIndex int) {
	// MinHeap means - parents are smaller than their childs, thus
		// child < parent
	for len(h.array) > 0 && h.array[lastIndex].freq < h.array[getParent(lastIndex)].freq {
		h.swap(lastIndex, getParent(lastIndex))
		lastIndex = getParent(lastIndex)
	}
}

func (h *MinHeap) MinHeapifyDown(firstIndex int) {
	// MinHeap Means parents are smaller than their childs, thus
	lastIndex := h.getSize()
	leftChild, rightChild := getLeftChild(firstIndex), getRightChild(firstIndex)
	childToCompare := 0
	// Which Child to compare with
	// while loop firstIndex <= lastIndex
	for firstIndex <= lastIndex {
		if leftChild > lastIndex { // End of the tree
			break
		}
		
		// In case right leaf not Exists (OutOfBound)
		if rightChild > lastIndex {
			childToCompare = leftChild
			// Left leaf is smaller 
		} else if h.array[leftChild].freq < h.array[rightChild].freq {
			childToCompare = leftChild
			// Right leaf is smaller
		} else {
			childToCompare = rightChild
		}

		// Swapping process
		if h.array[childToCompare].freq < h.array[firstIndex].freq {
			h.swap(childToCompare, firstIndex)
			firstIndex = childToCompare
			leftChild, rightChild = getLeftChild(firstIndex), getRightChild(firstIndex)
		} else {
			// MinHeap Satidfied - Exit
			return
		}

	}

}




// Helpers Methods
func getParent(index int)int{return (index - 1)/2}
func getLeftChild(index int)int{return 2*index + 1}
func getRightChild(index int)int{return 2*index + 2}
func (h *MinHeap) swap(index1, index2 int) {
	h.array[index1], h.array[index2]  = h.array[index2], h.array[index1]
}
func (h *MinHeap) getSize() int {return len(h.array) -1}



func topKFrequent(nums []int, k int) []int {
	// Create hashMap
	hashMap := make(map[int]int)
	for _, num := range nums {
		hashMap[num]++
	}
	// Convert it into MinHeap
	h := &MinHeap{k:k}
	for num, freq := range hashMap {
		h.Insert(Element{num:num, freq:freq})
	}
	// save it and return
	result := []int{}
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