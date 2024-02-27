package main

/* 交错字符串 */

func main() {

}

func isInterleave(s1 string, s2 string, s3 string) bool {
	n, m, t := len(s1), len(s2), len(s3)
	if (n + m) != t {
		return false
	}
	f := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		f[i] = make([]bool, m+1)
	}
	f[0][0] = true
	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			p := i + j - 1
			if i > 0 {
				f[i][j] = f[i][j] || (f[i-1][j] && s1[i-1] == s3[p])
			}
			if j > 0 {
				f[i][j] = f[i][j] || (f[i][j-1] && s2[j-1] == s3[p])
			}
		}
	}
	return f[n][m]
}

//func isInterleave(s1 string, s2 string, s3 string) bool {
//	m, n := len(s1), len(s2)
//	// 如果长度对不上，必然不可能
//	if m+n != len(s3) {
//		return false
//	}
//	// 备忘录，其中 -1 代表未计算，0 代表 false，1 代表 true
//	memo := make([][]int, m+1)
//	for i := range memo {
//		memo[i] = make([]int, n+1)
//		for j := range memo[i] {
//			memo[i][j] = -1
//		}
//	}
//
//	return dp(s1, 0, s2, 0, s3, memo)
//}
//
//// 定义：计算 s1[i..] 和 s2[j..] 是否能组合出 s3[i+j..]
//func dp(s1 string, i int, s2 string, j int, s3 string, memo [][]int) bool {
//	k := i + j
//	// base case，s3 构造完成
//	if k == len(s3) {
//		return true
//	}
//	// 查备忘录，如果已经计算过，直接返回
//	if memo[i][j] != -1 {
//		return memo[i][j] == 1
//	}
//
//	res := false
//	// 如果，s1[i] 可以匹配 s3[k]，那么填入 s1[i] 试一下
//	if i < len(s1) && s1[i] == s3[k] {
//		res = dp(s1, i+1, s2, j, s3, memo)
//	}
//	// 如果，s1[i] 匹配不了，s2[j] 可以匹配，那么填入 s2[j] 试一下
//	if j < len(s2) && s2[j] == s3[k] {
//		res = res || dp(s1, i, s2, j+1, s3, memo)
//	}
//	// 如果 s1[i] 和 s2[j] 都匹配不了，则返回 false
//	// 将结果存入备忘录
//	memo[i][j] = 0
//	if res {
//		memo[i][j] = 1
//	}
//
//	return res
//}
