package main

import "fmt"

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	for i := 0; i < len(nums); i++ {
		for j := i; j < i+nums[i] && j < len(nums); j++ {
			if i+nums[j] >= len(nums)-1 {
				return true
			} else {
				nextjump := i + nums[j] //
				for nextjump < len(nums)-1 && nextjump != 0 {
					for k := nextjump; k < nextjump+nums[nextjump]; k++ {
						//这里无用功太多了，重复量大，这个多了一次遍历了啊，怎么去剪枝
						if k+nums[k] >= len(nums)-1 {
							return true
						} else {
							nextjump = max(nextjump, k+nums[k])
						}
					}
				}
			}

		}
	}
	return false
}
func main() {
	t := []int{2, 3, 1, 1, 4}
	fmt.Println(canJump(t))
}
func climbStairs(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i < len(dp); i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
func generate(numRows int) [][]int {
	if numRows == 1 {
		return [][]int{{1}}
	}
	if numRows == 2 {
		return [][]int{{1}, {1, 1}}
	}
	es := [][]int{{1}, {1, 1}}
	for i := 3; i <= numRows; i++ {
		mid := make([]int, i)
		for j := 0; j < i; j++ {
			if j == 0 || j == i-1 {
				mid[j] = 1
			} else {
				mid[j] = es[i-2][j-1] + es[i-2][j]
			}

		}
		es = append(es, mid)
	}
	return es
}
func rob(nums []int) int {
     if 
}