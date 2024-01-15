package main

/* 加油站 */

func main() {

}

func canCompleteCircuit(gas []int, cost []int) int {
	currSum, sum := 0, 0
	start := 0
	for i := 0; i < len(gas); i++ {
		currSum += gas[i] - cost[i]
		sum += gas[i] - cost[i]
		if currSum < 0 {
			start = i + 1
			currSum = 0
		}
	}
	if sum < 0 {
		return -1
	}
	return start
}
