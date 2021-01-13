package main

func merge(nums1 []int, m int, nums2 []int, n int)  {
	for m > 0 && n > 0 {
		if nums1[m - 1] >= nums2[n - 1] {
			nums1[m + n - 1] = nums1[m - 1]
			m--
		} else {
			nums1[m + n - 1] = nums2[n - 1]
			n--
		}
	}

	for ;m >= 1; m-- {
		nums1[m + n - 1] = nums1[m - 1]
	}
	for ;n >= 1; n-- {
		nums1[m + n - 1] = nums2[n - 1]
	}
}

func main() {
	nums1 := []int{0}
	merge(nums1, 0, []int{1}, 1)
}
