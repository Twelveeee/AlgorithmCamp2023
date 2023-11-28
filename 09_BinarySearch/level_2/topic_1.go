package level2

// 山脉数组的峰顶索引
// https://leetcode.cn/problems/peak-index-in-a-mountain-array/

func peakIndexInMountainArray(arr []int) int {
	low, high := 1, len(arr)-2

	for low <= high {
		mid := (low + high) / 2
		leftVal, midVal, righVal := arr[mid-1], arr[mid], arr[mid+1]
		if leftVal < midVal && righVal < midVal {
			return mid
		} else if leftVal < midVal && righVal > midVal {
			low = mid + 1
		} else if leftVal > midVal && righVal < midVal {
			high = mid - 1
		}
	}

	return 0
}
