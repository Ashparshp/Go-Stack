package main

import (
	"fmt"
)

func add(a int, b int) int {
	return a + b
}

func substract(a, b int) int {
	return a - b
}

func getLanguages() (string, string, string, bool) {
	return "Go", "Java", "TS", true
}

func backFunction(fn func(a int) int) int {
	return fn(5)
}

func returnBack() func(b int) int {
	return func (b int) int {
		return b * b * b
		
	}
}

func main() {

	result := add(5, 3)
	fmt.Println(result)

	result = substract(5, 3)
	fmt.Println(result)

	println(getLanguages())

	fn := func(a int) int {
		return a * a
	}

	resultBack := backFunction(fn)
	fmt.Println(resultBack)

	
	fmt.Println(returnBack()(5))


	

}