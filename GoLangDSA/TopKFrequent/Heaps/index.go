package main

// MaxHeap struct has a slice that holds the array

type MaxHeap struct{
	array []int
}

// Insert adds an element to the heap
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

// Extract returns the largest key, and removes it from the heap

func (h *MaxHeap) maxHeapifyUp(index int) {

}

// Get the parent index
func parent(i int) int {
	return ((i - 1) /2)
}
// Get the left child
func left(i int) int {
	return 2 * i + 1
}
// Get the left child
func right(i int) int {
	return 2 * i + 2
}

// Swap keys in the array
func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}


func main() {

}