package main

import "fmt"

type Solution struct {
	Values []int
}

func findSolutions(candidates []int) []Solution {
	var solutions []Solution
	var currentSolution []int
	var backtrack func(int)

	backtrack = func(start int) {
		if isSolution(currentSolution) {
			solutions = append(solutions, Solution{Values: append([]int(nil), currentSolution...)}) //这个是为了防止共享
		} else {
			for i := start; i < len(candidates); i++ {
				if isValid(currentSolution, candidates[i]) {
					currentSolution = append(currentSolution, candidates[i])   //选择当前结果+下面的进一步选择
					backtrack(i + 1)                                           //进一步选择
					currentSolution = currentSolution[:len(currentSolution)-1] // 回溯
				}
			}
		}
	}

	backtrack(0)
	return solutions
}

func isSolution(solution []int) bool { //模板都确定了，问题是怎么样能够去进行,判断solution和选择是否合法，也就是写着两个函数
	// 判断当前solution是否为有效解
	return true // 假设这里有一个具体的判断逻辑
}

func isValid(solution []int, candidate int) bool {
	// 判断候选值是否可以被选中
	return true // 假设这里有一个具体的判断逻辑
}

func main() {
	candidates := []int{1, 2, 3, 4}
	solutions := findSolutions(candidates)
	for _, sol := range solutions {
		fmt.Println(sol.Values)
	}
}
