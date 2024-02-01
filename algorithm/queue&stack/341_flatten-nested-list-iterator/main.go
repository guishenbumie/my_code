package main

/* 扁平化嵌套列表迭代器 */

func main() {

}

// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (this NestedInteger) IsInteger() bool {
	return false
}

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (this NestedInteger) GetInteger() int {
	return 0
}

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (this *NestedInteger) Add(elem NestedInteger) {}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (this NestedInteger) GetList() []*NestedInteger {
	return nil
}

type NestedIterator struct {
	l []NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	l := make([]NestedInteger, 0)
	for _, x := range nestedList {
		l = append(l, *x)
	}
	return &NestedIterator{l}
}

func (this *NestedIterator) Next() int {
	res := this.l[0].GetInteger()
	this.l = this.l[1:]
	return res
}

func (this *NestedIterator) HasNext() bool {
	for len(this.l) > 0 && !this.l[0].IsInteger() {
		first := this.l[0].GetList()
		this.l = this.l[1:]
		for i := len(first) - 1; i >= 0; i-- {
			this.l = append([]NestedInteger{*first[i]}, this.l...)
		}
	}
	return len(this.l) > 0
}
