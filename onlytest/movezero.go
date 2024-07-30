package main

import (
	"sort"
	"strings"
)

func moveZeroes(nums []int) {
	ilow := -1 //永远指向最后一个非零区间元素
	pb := 0

	for pb < len(nums) {
		if nums[pb] != 0 {
			ilow++
			nums[pb], nums[ilow] = nums[ilow], nums[pb]
			pb++
		} else {
			pb++
		}

	}
}

func maxArea(height []int) int {
	maxa := 0

	length := len(height)

	for s := 0; s < length-1; s++ {
		for e := s + 1; e < length; e++ {
			l := e - s
			h := min(height[s], height[e])
			maxa = max(maxa, l*h)
		}
	}
	return maxa
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func reverseVowels(s string) string {
	st := 0
	ed := len(s) - 1
	mid := []rune{}
	for i := 0; i < len(s); i++ {
		mid = append(mid, rune(s[i]))
	}
	for st < ed {
		for !isvowal(rune(mid[st])) {
			st++
		}
		for !isvowal(rune(mid[ed])) {
			ed--
		}
		if st < ed { //事实上我发现一点，再双指针交换分区时，往往忽略了检查下标是否合理，这会导致错误
			mid[st], mid[ed] = mid[ed], mid[st]
			st++
			ed--
		}
	}
	return string(mid)
}
func isvowal(a rune) bool {
	switch a {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false

	}
}
func minSubArrayLen(target int, nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	cursum := 0
	count := 0
	for i := 0; i < len(nums); i++ {
		cursum = cursum + nums[i]
		count++
		if cursum >= target {
			return count
		}
	}
	return 0
}
func lengthOfLastWord(s string) int {
	ss := strings.TrimSpace(s)
	pre := -1
	for i := 0; i < len(ss); i++ {
		if ss[i] == uint8(' ') {
			pre = i
		}
	}
	return len(ss) - pre - 1
}
