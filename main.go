package main

import "fmt"
package main

import "fmt"

// DutchNationalFlag sorts the array such that all elements less than the pivot come first,
// followed by all elements equal to the pivot, followed by all elements greater than the pivot.
func DutchNationalFlag(nums []int, pivot int) {
	low, mid, high := 0, 0, len(nums)-1

	// Traverse the array with the mid pointer (or probe)
	for mid <= high {
		if nums[mid] < pivot {
			// Swap elements at low and mid pointers, increment both pointers
			nums[low], nums[mid] = nums[mid], nums[low]
			low++
			mid++
		} else if nums[mid] > pivot {
			// Swap elements at mid and high pointers, decrement high pointer
			nums[mid], nums[high] = nums[high], nums[mid]
			high--
		} else {
			// Move mid pointer forward if element is equal to the pivot
			mid++
		}
	}
}

func main() {
	nums := []int{2, 2, 1, 0, 2, 1, 0}
	pivot := 1
	fmt.Println("Unsorted array:", nums)
	DutchNationalFlag(nums, pivot)
	fmt.Println("Sorted array:", nums)
}

// DutchNationalFlag sorts the array such that all elements less than the pivot come first,
// followed by all elements equal to the pivot, followed by all elements greater than the pivot.
//func DutchNationalFlag(nums []int, pivot int) {
//	low, mid, high := -1, -1, len(nums)-1
//
//	//for mid <= high {
//	//	if nums[mid] < pivot {
//	//		nums[low], nums[mid] = nums[mid], nums[low]
//	//		low++
//	//		mid++
//	//	} else if nums[mid] > pivot {
//	//		nums[mid], nums[high] = nums[high], nums[mid]
//	//		high--//hign是上界限，low是下接线，mid是探针，采取交换不是覆盖，又是最右边的是没判断过，所以探针不会右移，左边已经判断过，右移
//	//	} else {
//	//		mid++
//	//	}
//	//}
//	i := 0
//	j := high
//	for i <= j {
//		if nums[i] < pivot {
//			low++
//			nums[i], nums[low] = nums[low], nums[i]
//		}
//		if nums[j] < pivot {
//
//		}
//	}
//}

func main() {
	nums := []int{2, 3, 1, 0, 2, 1, 0}
	pivot := 1
	fmt.Println("Unsorted array:", nums)
	dutchflag(nums, pivot)
	fmt.Println("Sorted array:", nums)
}

func quicks(nums []int, low, high int) {
	if low < high {
		i := pt(nums, low, high)
		quicks(nums, low, i-1)
		quicks(nums, i+1, high)
	}
}
func pt(nums []int, low, high int) int {
	pivot := nums[high]
	ilow := low - 1 //双色旗
	for i := low; i <= high-1; i++ {
		if nums[i] < pivot {
			ilow++
			nums[i], nums[ilow] = nums[ilow], nums[i]
		}
	}
	nums[ilow+1], nums[high] = nums[high], nums[ilow+1]
	return ilow + 1
}
func dutchflag(nums []int, pivot int) { //三色旗，以pivot划分递增分区
	ilow := -1
	pb := 0  //应该吧ihgh看做最后一个未定元素，ilow看做小于pivot区间的最后一个元素
	ihigh := len(nums) - 1 //ilow表示小于pivot区间最后一个值，ihigh表示大于pivot区间的前一个值，这样遵循快排分区点的逻辑规范，我感觉是最好理解的
	for pb <= ihigh {
		if nums[pb] < pivot {
			ilow++
			nums[ilow], nums[pb] = nums[pb], nums[ilow]
			pb++ //该探针位置已经处理过
		} else if nums[pb] > pivot {
			nums[ihigh], nums[pb] = nums[pb], nums[ihigh]
			ihigh-- //换过来的pb位未处理
		} else {
			pb++ //nums[pb]==pivot
		}

	}
}
