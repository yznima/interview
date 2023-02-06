package main

// https://leetcode.com/problems/fibonacci-number

// The Fibonacci numbers, commonly denoted F(n) form a sequence, called the Fibonacci sequence, such that each number is the sum of the two preceding ones, starting from 0 and 1. That is,
//
// F(0) = 0, F(1) = 1
// F(n) = F(n - 1) + F(n - 2), for n > 1.
// Given n, calculate F(n).

func fib(n int) int {
	return fibM(n, make(map[int]int, n))
}

func fibM(n int, mem map[int]int) int {
	if f, ok := mem[n]; ok {
		return f
	}

	var f int
	switch n {
	case 0:
		f = 0
	case 1:
		f = 1
	default:
		f = fib(n-1) + fib(n-2)
	}
	mem[n] = f
	return f
}
