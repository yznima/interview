// https://leetcode.com/problems/delete-and-earn/

// Given an array nums of integers, you can perform operations on the array.

// In each operation, you pick any nums[i] and delete it to earn nums[i] points.
// After, you must delete every element equal to nums[i] - 1 or nums[i] + 1.

// You start with 0 points. Return the maximum number of points you can earn by applying such operations.

func deleteAndEarn(nums []int) int {
    set := make(map[int]int, len(nums))
    ints := make([]int, 0, len(nums))
    for _, n := range nums {
        if _, exist := set[n]; exist {
            set[n]++
        } else {
            set[n] = 1
            ints = append(ints, n)
        }
    }
    
    sort.Ints(ints)
    avoid, using, prev := 0, 0, -1
    for _, n := range ints {
        max := max(avoid, using)
        if prev + 1 == n {
            avoid, using = max, avoid + n*set[n]
        } else {
            avoid, using = max, max + n*set[n]
        }
        prev = n
    }
    
    return max(avoid, using)
}

func max(i1 int, i2 int) int {
    if i1 > i2 {
        return i1
    } else {
        return i2
    }
}