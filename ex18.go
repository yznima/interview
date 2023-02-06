package main

import "math"

// https://leetcode.com/problems/minimum-falling-path-sum

// Given an n x n array of integers matrix, return the minimum sum of any falling path through matrix.
//
// A falling path starts at any element in the first row and chooses the element in the next row that is
// either directly below or diagonally left/right. Specifically, the next element from position (row, col)
// will be (row + 1, col - 1), (row + 1, col), or (row + 1, col + 1).

func minFallingPathSum(matrix [][]int) int {
	if len(matrix) == 1 {
		return matrix[0][0]
	}

	acc := make([][]int, len(matrix))
	for i, _ := range acc {
		l := make([]int, len(matrix[i]))
		for j := 0; j < cap(l); j++ {
			if i == 0 {
				l[j] = matrix[i][j]
			} else {
				l[j] = matrix[i][j] + minFallingPath(acc, i, j)
			}
		}

		acc[i] = l
	}

	m := math.MaxInt
	for _, v := range acc[len(acc)-1] {
		if v < m {
			m = v
		}
	}

	return m
}

func minFallingPath(matrix [][]int, i, j int) int {
	return min(
		val(matrix, i-1, j-1),
		val(matrix, i-1, j),
		val(matrix, i-1, j+1),
	)
}

func val(matrix [][]int, i, j int) int {
	if i >= len(matrix) || j >= len(matrix) || (i) < 0 || (j) < 0 {
		return math.MaxInt
	}

	return matrix[i][j]
}

func min(i, j, k int) int {
	if i <= j && i <= k {
		return i
	} else if j <= k && j <= i {
		return j
	} else {
		return k
	}
}
