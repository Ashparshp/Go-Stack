package main

import (
	"fmt"
)

func main() {

	var nums [4]int

	fmt.Println(len(nums))

	nums[0] = 1
	fmt.Println(nums[0])

	fmt.Println(nums)

	var b [4]bool
	fmt.Println(b)

	var names [4]string
	names[0] = "Go"
	fmt.Println(names)

	list := [3]int{1, 2, 3}
	fmt.Println(list)

	list2D := [2][2]int{{1, 2}, {5, 6}}
	fmt.Println(list2D)
	
}