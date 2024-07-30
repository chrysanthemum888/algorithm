package main

func canReach(s string, minJump int, maxJump int) bool {
	dp := make(map[int]bool, len(s))
	dp[0] = true
	for i := 1; i < len(s); i++ {
		if s[i] == '1' {
			dp[i] = false
			continue
		}
		if i-minJump >= 0 {
			flag := 0
			for k := i - minJump; k >= 0 && k >= i-maxJump; k-- {
				if dp[k] == true {
					dp[i] = true
					flag = 1
					break
				}
			}
			if flag == 0 {
				dp[i] = false
			}
		}
	}
	return dp[len(s)-1]
}
func maxSlidingWindow(nums []int, k int) []int {
	deque := []int{}  // 双端队列，存储元素的索引
	result := []int{} // 存储每个滑动窗口的最大值

	for i := 0; i < len(nums); i++ {
		// 移除不在当前滑动窗口内的元素索引
		if len(deque) > 0 && deque[0] == i-k {
			deque = deque[1:]
		}

		// 维护队列的单调递减性
		// 移除队列中所有比当前元素小的元素
		for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[i] {
			deque = deque[:len(deque)-1]
		}

		// 将当前元素的索引添加到队列
		deque = append(deque, i)

		// 当窗口的大小达到 k 时，记录当前窗口的最大值
		if i >= k-1 {
			result = append(result, nums[deque[0]])
		}
	}

	return result
}
