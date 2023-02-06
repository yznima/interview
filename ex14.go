package main

// https://leetcode.com/problems/generate-parentheses/

// Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

func generateParenthesis(n int) []string {
	return generateParenthesis2("(", n-1, n)
}

func generateParenthesis2(str string, l, r int) []string {
	if r == 0 {
		return []string{str}
	} else if l == 0 {
		return generateParenthesis2(str+")", l, r-1)
	} else if r > l {
		return append(generateParenthesis2(str+"(", l-1, r), generateParenthesis2(str+")", l, r-1)...)
	} else {
		return generateParenthesis2(str+"(", l-1, r)
	}
}
