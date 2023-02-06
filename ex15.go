package main

import "fmt"

func kthSmallest(root *TreeNode, k int) int {
	finder := &kthSmalletFinder{k: k, smallest: nil, n: 0}
	finder.Traverse(root)
	return finder.smallest.Val
}

type kthSmalletFinder struct {
	k        int
	smallest *TreeNode
	n        int
}

func (k *kthSmalletFinder) Traverse(t *TreeNode) bool {
	if t == nil {
		return false
	}

	found := k.Traverse(t.Left)
	if found {
		return found
	}

	found = k.Observe(t)
	if found {
		return found
	}

	found = k.Traverse(t.Right)
	if found {
		return found
	}

	return false
}

func (k *kthSmalletFinder) Observe(t *TreeNode) bool {
	if t == nil {
		return false
	}

	fmt.Println(t.Val, k.n, k.k)
	if k.n >= k.k {
		return true
	} else {
		k.smallest = t
		k.n++
		return k.n >= k.k
	}
}
