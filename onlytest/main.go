package main

func twoSum(numbers []int, target int) []int {
	s, e := 0, len(numbers)-1
	for s < e {
		sum := numbers[s] + numbers[e]
		if sum == target {
			return []int{s + 1, e + 1}
		} else if sum < target {
			s++
		} else {
			e--
		}
	}
	return []int{}
}
