package binarysearch

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	n1, n2 := len(nums1), len(nums2)
	N := n1 + n2

	start := 0
	end := n1

	for start <= end {

		cut1 := start + (end-start)/2 // mid -> cut1
		cut2 := N/2 - cut1

		l1 := math.MinInt
		if cut1 > 0 {
			l1 = nums1[cut1-1]
		}

		l2 := math.MinInt
		if cut2 > 0 {
			l2 = nums2[cut2-1]
		}

		r1 := math.MaxInt
		if cut1 < n1 {
			r1 = nums1[cut1]
		}

		r2 := math.MaxInt
		if cut2 < n2 {
			r2 = nums2[cut2]
		}

		// check if cut is valid
		if l1 <= r2 && l2 <= r1 {
			// if N is odd
			if N%2 != 0 {
				return float64(min(r1, r2))
			} else {
				// N is even
				return float64(max(l1, l2)+min(r1, r2)) / 2
			}
		} else if l1 > l2 {
			end = cut1 - 1
		} else {
			start = cut1 + 1
		}
	}

	return -1
}
