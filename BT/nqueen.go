package BT

import "fmt"

//for col := 0; col < n; col++ {
//if isValid(board, row, col, n) {
//board[row][col] = 'Q'
//backtrack(row + 1)
//board[row][col] = '.'
//}
//}这个我有点感到新奇，明明是双重循环来更新board的，这里却只用一层循环，把另外一层循环的行数当做选择来使用了，这种技巧很巧妙啊，还有其他地方用到这种技巧吗
package main

import (
"fmt"
"strings"
)
//这种技巧在解决组合问题和排列问题时非常有用。它通过递归函数的深度来模拟循环的嵌套，避免了显式的多重循环，使代码更加简洁易懂。这里有几个常见的应用场景：
//
//全排列问题：
//
//给定一个没有重复数字的序列，返回其所有可能的全排列。
//组合问题：
//
//从一组数中选择 k 个数的所有组合。
//子集问题：
//
//给定一组不包含重复数字的整数数组，返回该数组的所有可能子集。


func solveNQueens(n int) [][]string {
	results := [][]string{}
	board := make([][]byte, n)
	for i := range board {
		board[i] = make([]byte, n)
		for j := range board[i] {
			board[i][j] = '.'
		}
	}
	var backtrack func(row int)
	backtrack = func(row int) {
		if row == n {
			result := make([]string, n)
			for i := range board {
				result[i] = string(board[i])
			}
			results = append(results, result)
			return
		}
		for col := 0; col < n; col++ {//可以看到回溯问题中常常使用递归深度的选项来代替掉循环，方便掌握递归深度
			if isValid(board, row, col, n) {
				board[row][col] = 'Q'
				backtrack(row + 1)////可以看到回溯问题中常常使用递归深度的选项来代替掉循环，方便掌握递归深度
				board[row][col] = '.'
			}
		}
	}
	backtrack(0)
	return results
}

func isValid(board [][]byte, row, col, n int) bool {
	for i := 0; i < row; i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 'Q' {
			return false
		}
	}
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] == 'Q' {
			return false
		}
	}
	return true
}

func main() {
	n := 8
	solutions := solveNQueens(n)
	for _, solution := range solutions {
		for _, row := range solution {
			fmt.Println(row)
		}
		fmt.Println()
	}
}
