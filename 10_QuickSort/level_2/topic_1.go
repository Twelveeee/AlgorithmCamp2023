package level2

// 数组中的第K个最大元素
// https://leetcode.cn/problems/kth-largest-element-in-an-array/description/

func findKthLargest(nums []int, k int) int {
	return subFindKthLargest(nums, len(nums)-k-1, 0, len(nums)-1)
}

func subFindKthLargest(nums []int, k, low, high int) int {
	if low == high {
		return nums[low]
	}

	pVal := nums[low]
	i, j := low-1, high+1
	for i < j {
		for i++; nums[i] < pVal; i++ {
		}
		for j--; nums[j] > pVal; j-- {
		}

		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	p := j
	if p <= k {
		return subFindKthLargest(nums, k, p+1, high)
	}

	return subFindKthLargest(nums, k, low, p)

}
