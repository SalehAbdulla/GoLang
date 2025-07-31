package main

import (
    "fmt"
    "os"
    "strconv"
    "strings"
)

func GetPairs(nums []int, target int) []int {
    pairs := []int{}
    used := make(map[int]bool)
    for i := 0; i < len(nums); i++ {
        if used[i] {
            continue
        }
        for j := i + 1; j < len(nums); j++ {
            if used[j] {
                continue
            }
            if nums[i]+nums[j] == target {
                pairs = append(pairs, i, j)
                used[i], used[j] = true, true
                break
            }
        }
    }
    return pairs
}

func ParseNums(nums string) []int {
    if len(nums) < 2 || nums[0] != '[' || nums[len(nums)-1] != ']' {
        return nil
    }

    trimmed := strings.Trim(nums, "[]")
    parts := strings.Split(trimmed, ",")
    result := []int{}

    for _, part := range parts {
        part = strings.TrimSpace(part)
        if part == "" {
            continue
        }
        num, err := strconv.Atoi(part)
        if err != nil {
            fmt.Println("Invalid number:", part)
            return nil
        }
        result = append(result, num)
    }

    return result
}

func main() {
    args := os.Args[1:]
    if len(args) != 2 {
        fmt.Println("Invalid input.")
        return
    }

    nums := ParseNums(args[0])
    if nums == nil {
        fmt.Println("Invalid input.")
        return
    }

    target, err := strconv.Atoi(args[1])
    if err != nil {
        fmt.Println("Invalid target sum.")
        return
    }

    pairs := GetPairs(nums, target)
    if len(pairs) == 0 {
        fmt.Println("No pairs found.")
    } else {
        fmt.Printf("Pairs with sum %d: %v\n", target, pairs)
    }
}
