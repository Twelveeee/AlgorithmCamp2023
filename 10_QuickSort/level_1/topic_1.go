package level1

// 快速排序 填坑法
func quickSortV1(nums []int) []int {
	if len(nums) <= 1 {
		return nums

	}

	p := partionV1(nums)
	quickSortV1(nums[:p])
	quickSortV1(nums[p+1:])
	return nums
}

func partionV1(nums []int) int {
	low, high := 0, len(nums)-1

	pVal := nums[low]
	for low < high {
		for low < high && nums[high] >= pVal {
			high--
		}
		nums[low] = nums[high]

		for low < high && nums[low] <= pVal {
			low++
		}
		nums[high] = nums[low]
	}
	nums[low] = pVal
	return low
}

// 快速排序 顺序遍历法
func quickSortV2(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	p := partionV2(nums)
	quickSortV2(nums[:p])
	quickSortV2(nums[p+1:])
	return nums

}

func partionV2(nums []int) int {
	low, high := 0, len(nums)-1
	pVal := nums[low]
	i := low

	for j := low + 1; j <= high; j++ {
		if nums[j] < pVal {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}

	nums[i] = pVal
	return i
}
