package array

import "fmt"

func MergeSortArray(nums1 []int, m int, nums2 []int, n int) {
	t := m + n - 1
	m--
	n--
	for ; t >= 0; t-- {
		if m >= 0 && n >= 0 && nums1[m] > nums2[n] {
			nums1[t] = nums1[m]
			m--
		} else if n >= 0 {
			nums1[t] = nums2[n]
			n--
		}
	}
	fmt.Println(nums1)
}
