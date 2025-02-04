package main

import (
	"fmt"
	"slices"
)

func main() {

	var nums = []int{}

	fmt.Println(nums == nil)
	fmt.Println(len(nums))

	var numbers = make([]int, 2, 5)
	fmt.Println(numbers)
	fmt.Println(numbers == nil)
	fmt.Println(cap(numbers))

	numbers = append(numbers, 1)
	fmt.Println(numbers)

	numbers = append(numbers, 2)
	fmt.Println(numbers)

	numbers = append(numbers, 3)
	fmt.Println(numbers)

	numbers = append(numbers, 4)
	fmt.Println(numbers)

	fmt.Println(cap(numbers))
	fmt.Println(len(numbers))

	numbers[0] = 3
	fmt.Println(numbers)

	shortHand := []int{}
	fmt.Println(shortHand)
	fmt.Println(cap(shortHand))
	fmt.Println(len(shortHand))

	shortHand = append(shortHand, 1)
	fmt.Println(shortHand)

	var duplicate = make([]int, len(numbers))
	copy(duplicate, numbers)
	fmt.Println(numbers, duplicate)

	fmt.Println(numbers[0:2])
	fmt.Println(numbers[:2])
	fmt.Println(numbers[0:])

	fmt.Println(slices.Equal(numbers, nums))

	var slice2D = [][]int{}
	fmt.Println(slice2D)

	slice2D = append(slice2D, [][]int{{1, 2}}...)
	fmt.Println(slice2D)

}