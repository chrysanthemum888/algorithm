package main

import "fmt"

func k_distinct_quantity(nums []int, k int) []int { //这个不知道错哪里
	left, right := 0, 0
	imap := make(map[int]int)
	ucount := 0
	es := make([]int, 0)
	//for i := 0; i < k; i++ {
	//	imap[nums[i]]=true
	//}
	//right=k
	for right < len(nums) {
		_, ok := imap[nums[right]]
		if !ok {
			imap[nums[right]] = 1
			ucount++
			if right-left+1 == k {
				es = append(es, ucount)
			}
		} else {
			imap[nums[right]]++
		}
		if right >= k-1 {

			left_val := nums[left]
			if imap[left_val] > 0 {
				imap[left_val]--
				if imap[left_val] == 0 {
					delete(imap, left_val)
					ucount--
				}
			}
			left++
		}
		right++

	}
	return es
}
func main() {
	nums := []int{1, 2, 1, 2, 3}
	k := 3
	result := k_distinct_quantity(nums, k)
	fmt.Println(result) // Output: [2, 2, 3]
}
