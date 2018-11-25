// https://leetcode.com/problems/length-of-longest-fibonacci-subsequence

// A sequence X_1, X_2, ..., X_n is fibonacci-like if:

// n >= 3
// X_i + X_{i+1} = X_{i+2} for all i + 2 <= n

// Given a strictly increasing array A of positive integers forming a sequence,
// find the length of the longest fibonacci-like subsequence of A.  If one does 
// not exist, return 0.

// (Recall that a subsequence is derived from another sequence A by deleting 
// any number of elements (including none) from A, without changing the order 
// of the remaining elements.  For example, [3, 5, 8] is a subsequence of
// [3, 4, 5, 6, 7, 8].)
 
func lenLongestFibSubseq(A []int) int {
    set := make(map[int]int, len(A))
    for i, a := range A {
        set[a] = i
    }
    
    var longest int
    for m := 0; m < len(A) - 1; m++ {
        for n := m + 1; n < len(A); n++ {
        
            length, i, j := 0, m, n
            for j < len(A) {
                index, ok := set[A[i] + A[j]]
                if !ok {
                    break
                }

                if length == 0 {
                    length += 2
                }

                length++
                if length > longest {
                    longest = length
                }

                i, j = j, index
            } 
        }
    }
    
    return longest
}