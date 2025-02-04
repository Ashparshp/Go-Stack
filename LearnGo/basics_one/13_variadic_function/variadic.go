package main

import "fmt"

func sum(nums ...int) int {
	sum := 0

	for _, num := range nums {
		sum += num;
	}

	return sum
}

func main() {

	fmt.Println(1, 2, 3, 4, "Hello")

	fmt.Println(sum(1, 2, 3, 4 ,5, 6))

	numbers := []int{1, 2, 3, 4, 5}

	fmt.Println(sum(numbers...))
	
}