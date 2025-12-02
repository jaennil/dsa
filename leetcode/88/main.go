package main

func main() {

}

// nums1 4 5 6 0 0 0
// nums2 1 2 3

// nums1 4 5 6 7 0 0  m 4
// nums2 1 2 3  n 3

// func merge(nums1 []int, m int, nums2 []int, n int)  {
// 	p1, p2 := 0, 0
// 	for i := range m + n { // 
// 		if nums1[p1] <= nums2[p2] {
// 			nums1[i] = nums1[p1]
// 			p1++
// 		} else {
// 			nums2[p2], nums1[i] = nums1[i], nums2[p2]
// 		}
// 	}
// }

// 1 2 3 0 0 0
// 2 5 6

func merge(nums1 []int, m int, nums2 []int, n int)  {
	p1, p2 := m-1, n-1
	k := m+n-1
	for p2 > -1 {
		if nums1[p1] > nums2[p2] {
			nums1[k] = nums1[p1]
			p1--
		} else {
			nums1[k] = nums2[p2]
			p2--
		}
	}
}

