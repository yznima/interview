// https://leetcode.com/problems/subarray-sum-equals-k/

// Given an array of integers and an integer k, you need to find the total
// number of continuous subarrays whose sum equals to k.

func subarraySum(nums []int, k int) int {
    m := make(map[int]int, len(nums))
    m[0] = 1

    count := 0
    sum := 0
    for _, num := range nums {
        sum += num
        s := sum - k
        count += m[s]
        m[sum]++
    }
    
    return count
}

/** TOO SLOW
func subarraySum(nums []int, k int) int {
    count := 0
    for i := 0; i < len(nums); i++ {
        sum := 0
        for j := i; j < len(nums); j++ {
            sum += nums[j]
            if sum == k {
                count++
            }
        }
    }
    return count
}
*/
