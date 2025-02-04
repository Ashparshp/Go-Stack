package main

import "fmt"

func changeNum(num int) {
	num = 5
	fmt.Println("In changeNum", num)
}


func changeNumRefrence(num *int) {
	*num = 5
	fmt.Println("In changeNumRefrence", *num)
}


func changeSlice(nums []int) {
	nums = append(nums, 6)
	fmt.Println("In changeSlice", nums)
}


func changeSliceRefrence(nums *[]int) {
	*nums = append(*nums, 6)
	fmt.Println("In changeSlice", *nums)
}

func main() {

	num := 1

	changeNum(num)
	fmt.Println("After changeNum", num)

	changeNumRefrence(&num)
	fmt.Println("After changeNumRefrence", num)


	var nums = []int{1, 2, 3, 4, 5}
	changeSlice(nums)
	fmt.Println("After changeSlice", &nums)

	changeSliceRefrence(&nums)
	fmt.Println("After changeSliceRefrence", nums)
}