package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

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
func canJump2(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	maxReach := 0
	for i := 0; i < len(nums) && i <= maxReach; i++ {
		maxReach = max(maxReach, i+nums[i])
		if maxReach >= len(nums)-1 {
			return true
		}
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	minPrice := prices[0]
	maxP := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else {
			maxP = max(maxP, prices[i]-minPrice)
		}
	}

	return maxP
}
func maxProfit2(prices []int) int {

	total := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			total += prices[i] - prices[i-1]
		}
	}
	return total
}
func largestNumber2(nums []int) string {
	if len(nums) == 1 {
		return strconv.Itoa(nums[0])
	}
	maxstring := strconv.Itoa(nums[0])

	for i := 1; i < len(nums); i++ {
		now := nums[i]
		pre, _ := strconv.Atoi(maxstring)
		if numsconcat(pre, now) > numsconcat(now, pre) {
			maxstring = nums2string(pre, now)
		} else {
			maxstring = nums2string(now, pre)
		}
	}
	return maxstring
}
func numsconcat(a, b int) int {
	ia := strconv.Itoa(a)
	ib := strconv.Itoa(b)
	mid := ia + ib
	z, _ := strconv.Atoi(mid)
	return z
}
func nums2string(a, b int) string {
	ia := strconv.Itoa(a)
	ib := strconv.Itoa(b)
	return ia + ib
}
func largestNumber(nums []int) string {
	numstr := make([]string, len(nums))
	for i, v := range nums {
		numstr[i] = strconv.Itoa(v)
	}
	sort.Slice(numstr, func(i, j int) bool {
		if numstr[i]+numstr[j] > numstr[j]+numstr[i] {
			return true
		}
		return false
	})
	mid := strings.Join(numstr, "")

	re := regexp.MustCompile(`^0*`)
	res := re.ReplaceAllString(mid, "")

	if res == "" {
		res = "0"
	}
	return res

}
