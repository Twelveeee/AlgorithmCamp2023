package level1

// 合并两个有序数组
// https://leetcode.cn/problems/merge-sorted-array/
// 解法1 合并数组后排序
// 解法2 双指针temp数组
// 解法3 逆向双指针存入nums1 数组
func merge(nums1 []int, m int, nums2 []int, n int) {
	index1, index2 := m-1, n-1
	tail := m + n - 1
	curr := 0
	for index1 >= 0 && index2 >= 0 && tail > 0 {
		if nums1[index1] > nums2[index2] {
			curr = nums1[index1]
			index1--
		} else {
			curr = nums2[index2]
			index2--
		}

		nums1[tail] = curr
		tail--
	}

	for index1 >= 0 {
		nums1[tail] = nums1[index1]
		tail--
		index1--
	}
	for index2 >= 0 {
		nums1[tail] = nums2[index2]
		tail--
		index2--
	}
}
