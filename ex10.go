package main

// https://leetcode.com/problems/power-of-two/

// Given an integer n, return true if it is a power of two. Otherwise, return false.
//
// An integer n is a power of two, if there exists an integer x such that n == 2x.

func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}

	for ; n != 1; n = n / 2 {
		if n%2 != 0 {
			return false
		}
	}

	return true
}

// [[0,0,0,null,null,0,0,null,null,0,0],[0,0,0,null,null,0,0,0,0],[0,0,0,0,0,null,null,null,null,0,0],[0,0,0,0,0,null,null,0,0]]
// [[0,0,0,null,null,0,0,null,null,0,0],[0,0,0,null,null,0,0,0,0],[0,0,0,0,0,0,0],[0,0,0,0,0,null,null,null,null,0,0],[0,0,0,0,0,null,null,0,0]]
