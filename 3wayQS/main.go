package main

import "fmt"

func three_way_qs(nums []int, low, high int) {
	if low < high {
		s, e := partition(nums, low, high)

		three_way_qs(nums, low, s-1)
		three_way_qs(nums, e+1, high)
	}

}

func partition(nums []int, low, high int) (s, e int) {
	pivot := nums[high]
	ilow := low - 1
	pb := low         //ilow是什么呢，是最后一个小于的元素，pb最后是等于的元素的下一个，原因是pb《=ihigh被破坏
	ihigh := high - 1 //比pivot大的数的分区的前一个，也就是ihigh是最后一个未定元素，high是已经定了，最后处理时应该和等于pivot的最后一个坐标交换
	if low <= high {
		for pb <= ihigh { //
			if nums[pb] < pivot {
				ilow++
				nums[pb], nums[ilow] = nums[ilow], nums[pb]
				pb++ //pb已经处理过
			} else if nums[pb] > pivot {
				nums[pb], nums[ihigh] = nums[ihigh], nums[pb]
				ihigh-- //未定区间前移，但pb未处理
			} else {
				pb++ //和pivot相等,
			}
		} //最后pb指向等于pivot区间的最后一个元素，这时候只需将第一个大于pivot的值和原来的high交换即可
		//注意防止数组越界

		nums[pb], nums[high] = nums[high], nums[pb]
		return ilow + 1, pb
	}
	return -1, -1
}

// 另外一种实现可以少了数组越界判断，也就是按我之前的习惯for pb <= ihigh 写成<即可，这样的话pb

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	fmt.Println("Unsorted array:", nums)
	three_way_qs(nums, 0, len(nums)-1)
	fmt.Println("Sorted array:", nums)
}

//package main
//
//import "fmt"
//
//func threeWayQuickSort(nums []int, low, high int) {
//	if low < high {
//		s, e := partition(nums, low, high)
//		threeWayQuickSort(nums, low, s-1)
//		threeWayQuickSort(nums, e+1, high)
//	}
//}
//
//func partition(nums []int, low, high int) (int, int) {
//	pivot := nums[high]
//	ilow := low - 1 // 小于pivot的区域
//	pb := low       // 当前探针
//	ihigh := high   // 大于pivot的区域起点
//
//	for pb < ihigh {
//		if nums[pb] < pivot {
//			ilow++
//			nums[ilow], nums[pb] = nums[pb], nums[ilow]
//			pb++
//		} else if nums[pb] > pivot {
//			ihigh--
//			nums[ihigh], nums[pb] = nums[pb], nums[ihigh]
//		} else {
//			pb++
//		}
//	}
//	nums[ihigh], nums[high] = nums[high], nums[ihigh]
//	return ilow + 1, ihigh
//}
//
//func main() {
//	nums := []int{2, 0, 2, 1, 1, 0}
//	fmt.Println("Unsorted array:", nums)
//	threeWayQuickSort(nums, 0, len(nums)-1)
//	fmt.Println("Sorted array:", nums)
//}
