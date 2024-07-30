package main

import (
	"fmt"
	"math"
)

func divide(dividend int, divisor int) int {
	// 处理特殊情况
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32 // 处理溢出情况
	}

	// 记录结果符号
	negative := (dividend < 0) != (divisor < 0)

	// 使用绝对值
	dividend = abs(dividend)
	divisor = abs(divisor)

	result := 0
	// 逐步减去 divisor 的倍数
	for dividend >= divisor {
		temp := divisor
		multiplier := 1
		for dividend >= (temp << 1) {
			temp <<= 1
			multiplier <<= 1
		}
		dividend -= temp
		result += multiplier
	}

	// 根据符号调整结果
	if negative {
		result = -result
	}

	return result
}

// 辅助函数：返回绝对值
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

//func main() {
//	fmt.Println(divide(10, 3))           // 输出: 3
//	fmt.Println(divide(7, -3))           // 输出: -2
//	fmt.Println(divide(-2147483648, -1)) // 输出: 2147483647
//}
package main

import "fmt"

func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	//height := []int{0,1,0,2,1,0,1,3,2,1,2,1}
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	totalWater := 0

	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				totalWater += leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				totalWater += rightMax - height[right]
			}
			right--
		}
	}

	return totalWater
}

func main() {
//	height := []int{0,1,0,2,1,0,1,3,2,1,2,1}
//	fmt.Println(trap(height))  // 输出 6



}
