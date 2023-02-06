package main

// Given an array A, we may rotate it by a non-negative integer K so that the
// array becomes A[K], A[K+1], A{K+2], ... A[A.length - 1], A[0], A[1], ..., A[K-1].
// Afterward, any entries that are less than or equal to their index are worth 1 point.

// For example, if we have [2, 4, 1, 3, 0], and we rotate by K = 2, it becomes
// [1, 3, 0, 2, 4].  This is worth 3 points because 1 > 0 [no points],
// 3 > 1 [no points], 0 <= 2 [one point], 2 <= 3 [one point], 4 <= 4 [one point].

// Over all possible rotations, return the rotation index K that corresponds to
// the highest score we could receive.  If there are multiple answers, return the
// smallest such index K.

func bestRotation(A []int) int {
	n := len(A)
	intervals := make([]int, n)
	for i, a := range A {
		start := (i + 1)
		end := start + n - a

		if start < n {
			intervals[start]++

			if end < n {
				intervals[end]--
			}
		}

		if a <= i {
			intervals[0]++

			end2 := i - a + 1
			if end2 < n {
				intervals[end2]--
			}
		}
	}

	var sum, score, index int
	for i, j := range intervals {
		if sum += j; sum > score {
			score = sum
			index = i
		}
	}

	return index
}
