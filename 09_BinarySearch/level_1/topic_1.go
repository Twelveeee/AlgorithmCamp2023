package level1

// 二分查找
func binarySearchV1(list []int, target int) int {
	if len(list) == 0 {
		return -1
	}

	low, high := 0, len(list)-1
	for high >= low {
		mid := (high + low) / 2
		val := list[mid]
		if val == target {
			return mid
		} else if val > target {
			high = mid - 1
		} else if val < target {
			low = mid + 1
		}

	}
	return -1
}
