package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	if len(nums2) == 0 {
		return
	}

	p1 := 0
	p2 := 0
	for p1 < len(nums1) {
		if nums2[p2] <= nums1[p1] {
			for i := len(nums1) - 1; i > p1; i-- {
				nums1[i] = nums1[i-1]
			}
			nums1[p1] = nums2[p2]
			p2++
			if p2 >= n {
				break
			}
		}

		p1++
	}

	// insert the rest
	// add where the p1 pointer is at...
	for p2 < n {
		nums1[p2+m] = nums2[p2]
		p2++
	}
}

func merge2(nums1 []int, m int, nums2 []int, n int) {
	i, j := m-1, n-1
	for k := m + n - 1; j >= 0; k-- {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
	}
}
