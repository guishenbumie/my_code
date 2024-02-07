package main

/* 括号生成 */

func main() {

}

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	ans := []string{}
	path := []byte{}
	var dfs func(left, right int)
	dfs = func(left, right int) {
		if left > right {
			return
		}
		if left < 0 || right < 0 {
			return
		}
		if left == 0 && right == 0 {
			ans = append(ans, string(path))
			return
		}

		path = append(path, '(')
		dfs(left-1, right)
		path = path[:len(path)-1]

		path = append(path, ')')
		dfs(left, right-1)
		path = path[:len(path)-1]
	}
	dfs(n, n)
	return ans
}
