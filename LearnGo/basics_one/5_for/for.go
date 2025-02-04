package main

import "fmt"

func main() {
	i := 1

	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for num := 0; num < 11; num++ {
		fmt.Println(num)
	
	}

	for skip := 0; skip < 5; skip++ {
		if skip == 1 {
			continue
		}
		fmt.Println(skip)
	}

	for time := range [24]int{} {
		fmt.Println(time+1)
	}

}