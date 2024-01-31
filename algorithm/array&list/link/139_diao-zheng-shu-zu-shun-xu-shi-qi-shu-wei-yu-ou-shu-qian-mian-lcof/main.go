package main

/* 训练计划 I */

func main() {

}

func trainingPlan(actions []int) []int {
	slow, fast := 0, 0
	for fast < len(actions) {
		if actions[fast]%2 == 1 {
			actions[slow], actions[fast] = actions[fast], actions[slow]
			slow++
		}
		fast++
	}
	return actions
}
