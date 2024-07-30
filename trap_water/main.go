package trap_water

import "fmt"

package main

import "fmt"
//if height[left] < height[right]这里面为什么这么执着于判断两个指针所指的高度哪个更高
//前提设置是left和right分别是未定区间的左右边界，在指针判断时候，对于高的一段一律跳过，
//这其中包括了两端的较高端的跳过（这个跳过是不做考虑，left和right指针并不移动），还包括在较矮的一端跳过比之前高的

//这个逻辑很漂亮，在高的那端肯定能堵住第的那端的雨水，所以只需考虑低的那端和之前最低点之间能不能存雨水，这样直接把考虑的事情减少了一半了
//核心思想
//堵住雨水的原理：
//
//无论左侧还是右侧，较高的一端总能“堵住”雨水。只有较矮的一端决定了当前能存储的水量，因为水总是从低处积累起。
//减少考虑范围：
//
//每次只处理较矮一侧的指针，不用担心较高一侧，这样直接减少了一半的计算量。
//详细逻辑解析
//左侧高度较矮
//当 height[left] < height[right] 时：
//
//左侧较矮，说明右侧可以提供一个“高墙”来阻挡水，因此当前的水位取决于左侧的高度。
//只需要更新左侧最大高度 leftMax，并计算当前左侧位置能存储的水量。
//然后将左指针右移，以便继续检查下一个位置。
//右侧高度较矮
//当 height[left] >= height[right] 时：
//
//右侧较矮，说明左侧可以提供一个“高墙”来阻挡水，因此当前的水位取决于右侧的高度。
//只需要更新右侧最大高度 rightMax，并计算当前右侧位置能存储的水量。
//然后将右指针左移，以便继续检查下一个位置。
func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}

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
			left++//处理过后未定区间指针增加，left是指向未定区间的左闭，right同理，这个和快排的分区--我的习惯ilow指向low-1为小于区间最后一个不一样
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
	height := []int{0,1,0,2,1,0,1,3,2,1,2,1}
	fmt.Println(trap(height))  // 输出 6
}
