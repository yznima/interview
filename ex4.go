package main

// https://leetcode.com/problems/binary-gap/

// Given a positive integer N, find and return the longest distance between
// two consecutive 1's in the binary representation of N.

// If there aren't two consecutive 1's, return 0.

func binaryGap(N int) int {
	start, end, dist := -1, -1, 0
	for i := uint(0); i < 64; i++ {
		if N&(1<<i) == 0 {
			continue
		}

		if start == -1 {
			start = int(i)
		} else if end == -1 {
			end = int(i)
			dist = end - start
		} else {
			start, end = end, int(i)
			if t := end - start; t > dist {
				dist = t
			}
		}
	}

	return dist
}
