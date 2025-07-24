package main


func main(){

}


// Create a struct

type Element struct {
	num int
	freq int
}

type MaxHeap struct {
	array []Element
}

// ----------

// Insert And Extract Methods

func (h *MaxHeap) Insert(node Element) {
	// append to the array -> goes to the end
	h.array = append(h.array, node)
	// Heapify starts from last Index
	h.MaxHeapifyUp(len(h.array) - 1)
}

// MaxHeapifyUp and MaxHeapifyDown

func (h *MaxHeap) MaxHeapifyUp(lastIndex int){
	// compare with its parent
	// MaxHeapify Means the parent is bigger than its child
	for getParent(lastIndex) < lastIndex {
		// Keep Swapping
		h.swap(getParent(lastIndex), lastIndex)
	}

}

// Helper Methods
func getParent(node int) int {
	return (node - 1) * 2
}

func getLeftChild(node int) int {
	return node * 2 + 1
}

func getRightChild(node int) int {
	return node * 2 + 2
}

func (h *MaxHeap) swap(node1, node2 int) {
	h.array[node1], h.array[node2] = h.array[node2], h.array[node1]
}

func topKFrequent(nums []int, k int) []int {
	// Create a hashMap number -> Occurance
	hashMap := make(map[int]int)
	for _, num := range nums {
		hashMap[num]++
	}

	// Convert it into MaxHeap

	// Extract top k

	// return result
	return []int{}
}
