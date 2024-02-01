package main

/* 最小栈 */

func main() {

}

type MinStack struct {
	stk    []int
	minStk []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.stk = append(this.stk, val)
	if len(this.minStk) == 0 || this.minStk[len(this.minStk)-1] > val {
		this.minStk = append(this.minStk, val)
	} else {
		this.minStk = append(this.minStk, this.minStk[len(this.minStk)-1])
	}
}

func (this *MinStack) Pop() {
	this.stk = this.stk[:len(this.stk)-1]
	this.minStk = this.minStk[:len(this.minStk)-1]
}

func (this *MinStack) Top() int {
	return this.stk[len(this.stk)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStk[len(this.minStk)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
