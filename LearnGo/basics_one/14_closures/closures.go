package main

import "fmt"

func counter() func() int {
	var count int = 0

	return func() int {
		count += 1
		return count
	}
}



func giftCard() func(spend int) int {
	amount := 1000

	return func(spend int) int {
		amount -= spend
		return amount
	}

}


func main() {

	increment := counter()

	fmt.Println(increment())
	fmt.Println(increment())


	spend := giftCard()
	spend_s := giftCard()

	println(spend(100))
	println(spend(200))

	println(spend_s(100))
	println(spend_s(200))

}