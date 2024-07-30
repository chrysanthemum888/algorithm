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
