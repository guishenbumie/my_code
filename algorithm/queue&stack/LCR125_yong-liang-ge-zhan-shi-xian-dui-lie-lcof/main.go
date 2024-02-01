package main

/* 图书整理 II */

func main() {

}

type CQueue struct {
	s1 []int
	s2 []int
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.s1 = append(this.s1, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.s2) == 0 {
		for len(this.s1) > 0 {
			this.s2 = append(this.s2, this.s1[len(this.s1)-1])
			this.s1 = this.s1[:len(this.s1)-1]
		}
	}
	if len(this.s2) == 0 {
		return -1
	}
	x := this.s2[len(this.s2)-1]
	this.s2 = this.s2[:len(this.s2)-1]
	return x
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
