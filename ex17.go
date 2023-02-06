package main

// https://leetcode.com/problems/flatten-nested-list-iterator

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

type Iterator interface {
	Next() int
	HasNext() bool
}

type NestedInteger struct {
}

func (i NestedInteger) IsInteger() bool {
	panic("implement me")
}

func (i NestedInteger) GetList() []*NestedInteger {
	panic("implement me")
}

func (i NestedInteger) GetInteger() int {
	panic("implement me")
}

func Constructor(nestedList []*NestedInteger) Iterator {
	its := make([]Iterator, 0, len(nestedList))
	for _, l := range nestedList {
		if l.IsInteger() {
			its = append(its, NewIntegerIterator(l))
		} else {
			its = append(its, Constructor(l.GetList()))
		}
	}

	return &NestedIterator{its: its, i: 0}
}

type NestedIterator struct {
	its []Iterator
	i   int
}

func (this *NestedIterator) Next() int {
	if this.i >= len(this.its) {
		panic("next without value")
	}

	return this.its[this.i].Next()
}

func (this *NestedIterator) HasNext() bool {
	for ; this.i < len(this.its); this.i++ {
		if this.its[this.i].HasNext() {
			return true
		}
	}

	return false
}

type IntegerIterator struct {
	val  int
	done bool
}

func NewIntegerIterator(i *NestedInteger) Iterator {
	if !i.IsInteger() {
		panic("invalid nested integer")
	}

	return &IntegerIterator{val: i.GetInteger(), done: false}
}

func (this *IntegerIterator) Next() int {
	if this.done {
		panic("next without value")
	}
	this.done = true
	return this.val
}

func (this *IntegerIterator) HasNext() bool {
	return !this.done
}
