package main

import (
	"fmt"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	for n != 0 {
		if m != 0 && nums1[m-1] > nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}

}
func main() {
	int1 := []int{1, 2, 3, 0, 0, 0}
	merge(int1, 3, []int{2, 5, 6}, 3)
	fmt.Println(int1)

}
