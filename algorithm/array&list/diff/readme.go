package main

//差分数组的主要适用场景是频繁对原始数组的某个区间的元素进行增减

func main() {

}

// 核心代码
type Diff struct {
	diff []int
}

// 构造函数
func Constructor(nums []int) Diff {
	if len(nums) <= 0 {
		return Diff{}
	}
	diff := make([]int, len(nums))
	diff[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		diff[i] = nums[i] - nums[i-1]
	}
	return Diff{diff: diff}
}

// 给闭区间[i,j]加v（可以是负数）
func (d *Diff) incr(i, j, v int) {
	d.diff[i] += v
	//当j+1>=diff.length时说明是对i之后的所有数据进行修改，diff还是保持不变的
	if j+1 < len(d.diff) {
		d.diff[j+1] -= v
	}
}

func (d *Diff) result() []int {
	res := make([]int, len(d.diff))
	res[0] = d.diff[0]
	for i := 1; i < len(d.diff); i++ {
		res[i] = res[i-1] + d.diff[i]
	}
	return res
}
