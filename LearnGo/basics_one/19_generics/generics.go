package main

/*
func printSlice(items []int) {

	for _, item := range items {
		println(item)
	}

}

func printSlice(items []string) {

	for _, item := range items {
		println(item)
	}

}

func printSlice[T any](items []T) {

	for _, item := range items {
		println(item)
	}

}

func printSlice[T int | string | bool](items []T) {

	for _, item := range items {
		println(item)
	}

}

func printSlice[T comparable](items []T) {

	for _, item := range items {
		println(item)
	}

}

type stack struct {
	element []int
}
*/

type stack[T any] struct {
	element []T
}

func main() {

	/*
		nums := []int{1, 2, 3, 4, 5}
		printSlice(nums)

		languages := []string{"Go", "Python", "Java", "C++"}

		printSlice(languages)

		stack := stack{element: []int{1, 2, 3, 4, 5}}
		fmt.Println(stack)
	*/

	flag := []bool{true, false, true, false, true}
	stack := stack[bool]{element: flag}
	println(stack)

}
