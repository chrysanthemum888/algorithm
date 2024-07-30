package main

import "fmt"

func minCost(cost [][]int) int {
	MAXX := len(cost)
	Maxy := len(cost[0])
	dp := make([][]int, MAXX+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, Maxy+1)
	}
	for i := 1; i <= MAXX; i++ {
		for j := 1; j <= Maxy; j++ {
			dp[i][j] = cost[i-1][j-1] + min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1])
		}
	}
	return dp[MAXX][Maxy]
}
func main() {
	set := []int{3, 34, 4, 12, 5, 2}
	sum := 9
	fmt.Println(isSubSum(set, sum)) // 输出最小路径和
}

func isSubSum(nums []int, sum int) bool {
	MX := len(nums)
	MY := sum
	var dp [][]bool
	for i := 0; i <= MX; i++ {
		dp[i] = make([]bool, MY+1)
		dp[i][0] = true
	}
	for i := 1; i <= MY; i++ {
		dp[0][i] = false
	}
	for i := 1; i <= MX; i++ {
		for j := 1; j <= MY; j++ {
			dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[j-1]]
		}
	}
	return dp[MX][MY]
}
func findSolution(){
	var res collection_for_recording
	if isRES(){
		addtoRes()
	}else{
		for val:= candidates {
			if select(val){
				findSolution_with_val()
				backtrack()
			}
		}
	}
	return
}
