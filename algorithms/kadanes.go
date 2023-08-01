package main

import "fmt"

func main() {
	fmt.Println("Hello, World")
	input := [8]int{-2,-3,4,-1,-2,1,5,-3}
	fmt.Println(maxSum(input))
}

func maxSum(array [8]int) int {
	for index, value := range(array) {
		fmt.Println(index, value)
	}
	// remove
	return 0
}
