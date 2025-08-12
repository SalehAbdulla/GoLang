package piscine


func topKFrequent(nums []int, k int) []int {
    hashMap := make(map[int]int)
    for _, n := range nums {
        hashMap[n]++
    }
    
    var mapToSlice [][]int
    for num, occur := range hashMap {
        mapToSlice = append(mapToSlice, []int{num, occur})
    }

    // [2, 3],[6, 1],[3, 4],[1, 2],[0, 2]
    // [[],[6],[1, 0],[2],[3]] // zero will never be used
    // then loop < k {pop last k elements}

    result := make([][]int, len(nums)+1)
    for _, pairs := range mapToSlice {
        getPosition := pairs[1]
        result[getPosition] = append(result[getPosition], pairs[0])
    }

    var finalResult []int

    for i := len(result)-1; i >= 0; i-- {
        for _, num := range result[i] {
            if len(finalResult) == k {break}
            finalResult = append(finalResult, num)
        }
    }
    

    return finalResult
}