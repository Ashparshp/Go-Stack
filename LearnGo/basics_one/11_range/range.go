package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}

	for index, num := range nums {
		fmt.Println(index, " ", num)
	}
	
	mapOne := map[string]int {"First:": 1, "Second:": 2, "Third:": 3, "Fourth:": 4, "Fifth:": 5}

	for k, v := range mapOne {
		fmt.Println(k, v)
	}

	for keys := range mapOne {
		fmt.Println(keys)
	}

	for i, c := range "_ƀGolang" {
		fmt.Println("unicode:", c, string(c))
		fmt.Println("byte:", i)
	}

	

}