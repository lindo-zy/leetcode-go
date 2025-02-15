package main

import "fmt"

func main() {

	nums1 := []int{1, 3}
	nums2 := []int{2}
	res := findMedianSortedArrays(nums1, nums2)
	fmt.Println(res)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	total := m + n
	if total%2 == 1 {
		return float64(find(nums1, nums2, total/2+1))
	} else {
		return float64(find(nums1, nums2, total/2)+find(nums1, nums2, total/2+1)) / 2.0
	}
}
func find(nums1 []int, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
		if index1 == len(nums1) {
			return nums2[index2+k-1]
		}
		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}
		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}
		newIndex1 := min(index1+k/2-1, len(nums1)-1)
		newIndex2 := min(index2+k/2-1, len(nums2)-1)
		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
		if pivot1 <= pivot2 {
			k -= newIndex1 - index1 + 1
			index1 = newIndex1 + 1
		} else {
			k -= newIndex2 - index2 + 1
			index2 = newIndex2 + 1
		}
	}

}
