package main

func topKFrequent(nums []int, k int) []int {
    hashMap := make(map[int]int)
    for _, num := range nums{hashMap[num]++}

    result := make([][]int, len(nums)+1)
    for num, freq := range hashMap {
        result[freq] = append(result[freq], num)
    }

    var finalResult []int
    for i := len(result) - 1; i >= 0; i-- {
        for _, val := range result[i]{
            if len(finalResult) == k {break}
            finalResult = append(finalResult, val)
        }
    }
    
    return finalResult
}