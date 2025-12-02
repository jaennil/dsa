package main

func removeDuplicates(nums []int) int {
    left, right := 0, 1
    eq := false
    for right < len(nums) {
        if nums[left] == nums[right] {
            right++
            eq = true
        } else {
            left++
            if eq {
                left++
                eq = false
            }
            nums[left] = nums[right]
            right++
        }
    }

    return left+1
}

func main() {
	println(removeDuplicates([]int{0,0,1,1,1,1,2,3,3}))
}
