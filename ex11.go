package main

// https://leetcode.com/problems/all-possible-full-binary-trees

// Given an integer n, return a list of all possible full binary trees with n nodes. Each node of each tree in the answer must have Node.val == 0.
//
// Each element of the answer is the root node of one possible tree. You may return the final list of trees in any order.
//
// A full binary tree is a binary tree where each node has exactly 0 or 2 children.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func allPossibleFBT(n int) []*TreeNode {
	if n%2 == 0 {
		return []*TreeNode{}
	}

	return allPossibleFBTRecursive(n, make(map[int][]*TreeNode, n))
}

func allPossibleFBTRecursive(n int, mem map[int][]*TreeNode) []*TreeNode {
	if trees, ok := mem[n]; ok {
		return trees
	}

	var trees []*TreeNode

	switch n {
	case 1:
		trees = []*TreeNode{&TreeNode{}}
	case 3:
		trees = []*TreeNode{&TreeNode{Left: &TreeNode{}, Right: &TreeNode{}}}
	default:
		trees = make([]*TreeNode, 0, 2*n)
		for x := 1; x <= n-2; x = x + 2 {
			lefts := allPossibleFBTRecursive(x, mem)
			rights := allPossibleFBTRecursive(n-x-1, mem)
			for _, l := range lefts {
				for _, r := range rights {
					trees = append(trees, &TreeNode{
						Right: r,
						Left:  l,
					})
				}
			}
		}
	}

	mem[n] = trees
	return trees
}
