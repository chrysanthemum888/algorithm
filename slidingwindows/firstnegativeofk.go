package main

import (
	"math"
)

func first_ng_st_of_k(nums []int, k int) []int {
	dq := make([]int, 0)
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ { //总结窗口题三部曲，（遍历数组）1.排除窗口外的元素，2.窗口内操作，3.维护结果集合
		if len(dq) > 0 && dq[0] <= i-k+1 {
			dq = dq[0:]
		}
		if nums[i] < 0 {
			dq = append(dq, nums[i])
		}
		if i >= k-1 {
			if len(dq) > 0 {
				res = append(res, dq[0])
			} else {
				res = append(res, 0)
			}
		}

	}
	return res
}

// dq:=make([]int,0)
// fmt.Println(dq[0])--访问空切片会出错
// 给定三个整数X、Y和P，任务是找到最小窗口大小K，使这个大小范围[X,Y]中的每个窗口至少有P个素数。
func isPrime(tg int) bool {
	if tg <= 1 {
		return false
	}
	if tg == 2 {
		return true
	}
	if tg%2 == 0 {
		return false
	}
	for i := 3; i <= int(math.Sqrt(float64(tg))); i += 2 {
		if tg%i == 0 {
			return false
		}
	}
	return true
}

func Pprime2(x, y, p int) int {
	left, right := x, x
	midcount := 0
	clen := 0
	minlen := y - x + 1 + 1
	dq := make([]int, 0)
	for right <= y {
		if isPrime(right) {
			midcount++
			dq = append(dq, right)
		}
		right++
		clen = right - left
		if midcount == p {
			minlen = min(minlen, clen)
			//left++//这步处理出问题了，无法直接找到窗口内第一个素数的位置，这里必须改进成双端队列来记录窗口的第一个素数位置
			left = dq[0] + 1
			midcount--
			dq = dq[1:]
		}

	}
	if minlen == y-x+1+1 {
		return -1
	}
	return minlen
}

// 给定一个大小为N的队列和一个整数K，返回所有大小为K的窗口中不同数字的数量。
//
// 例子：
//
// 输入： arr[] = {1, 2, 1, 3, 4, 2, 3}, K = 4
// 输出： 3 4 4 3
