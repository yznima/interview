package main

// https://leetcode.com/problems/sort-integers-by-the-power-value

import (
	"sort"
)

func getKth(lo int, hi int, k int) int {
	vps := make(ValPowers, hi-lo+1)
	mem := make(map[int]int, hi-lo+1)
	for i := lo; i <= hi; i++ {
		vps[i-lo] = ValPower{Val: i, Power: power(i, mem)}
	}

	sort.Sort(vps)
	return vps[k-1].Val
}

type ValPower struct {
	Val, Power int
}

type ValPowers []ValPower

func (v ValPowers) Len() int { return len(v) }
func (v ValPowers) Less(i, j int) bool {
	if v[i].Power < v[j].Power {
		return true
	} else if v[i].Power == v[j].Power {
		return v[i].Val < v[j].Val
	} else {
		return false
	}
}

func (v ValPowers) Swap(i, j int) { v[i], v[j] = v[j], v[i] }

func power(n int, mem map[int]int) int {
	if n == 1 {
		return 0
	}

	if mem[n] > 0 {
		return mem[n]
	}

	var steps int
	if n%2 == 0 {
		steps = 1 + power(n/2, mem)
	} else {
		steps = 1 + power(3*n+1, mem)
	}

	mem[n] = steps
	return steps
}
