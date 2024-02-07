package main

import "strings"

/* N 皇后 II */

func main() {

}

func totalNQueens(n int) int {
	ans := [][]string{}
	board := make([]string, n)
	for i := 0; i < n; i++ {
		board[i] = strings.Repeat(".", n)
	}
	var backtrack func(board []string, row int)
	backtrack = func(board []string, row int) {
		if row == len(board) {
			temp := make([]string, len(board))
			copy(temp, board)
			ans = append(ans, temp)
			return
		}

		n := len(board[row])
		for col := 0; col < n; col++ {
			if !isValid(board, row, col) {
				// 排除不合法选择
				continue
			}
			// 做选择
			board[row] = board[row][:col] + "Q" + board[row][col+1:]
			// 进入下一行决策
			backtrack(board, row+1)
			// 撤销选择
			board[row] = board[row][:col] + "." + board[row][col+1:]
		}
	}
	backtrack(board, 0)
	return len(ans)
}

func isValid(board []string, row, col int) bool {
	n := len(board)
	// 检查列是否有皇后互相冲突
	for i := 0; i <= row; i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}
	// 检查右上方是否有皇后互相冲突
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] == 'Q' {
			return false
		}
	}
	// 检查左上方是否有皇后互相冲突
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 'Q' {
			return false
		}
	}
	return true
}
