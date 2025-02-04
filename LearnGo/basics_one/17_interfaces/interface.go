package main

import "fmt"

type paymenter interface {
	payment(amount float32)
	refund(amount float32, account string)
}

type payment struct{
	// gateway paytm
	gateway paymenter
}

func (p payment) makePayment(amount float32) {
	paymentGateway := razorpay{}
	paymentGateway.payment(amount)

	// paymentGateway := paytm{}
	// paymentGateway.payment(amount)

	p.gateway.payment(amount)
}

type razorpay struct{}

func (r razorpay) payment(amount float32) {
	fmt.Println("Payment done using razorpay", amount)
}

func (r razorpay) refund(amount float32, account string) {
	fmt.Println("Refund done using razorpay", amount, account)
}

// type paytm struct{}

// func (p paytm) payment(amount float32) {
// 	fmt.Println("Payment done using paytm", amount)
// }

// type test struct{}

// func (t test) payment(amount float32) {
// 	fmt.Println("Payment done using test", amount)
// }


func main() {

	newPayment := payment{}
	newPayment.makePayment(1000)

	paytmtGateway := razorpay{}
	// paytmtGateway := paytm{}
	// paytmtGateway := test{}
	paymentOne := payment{gateway: paytmtGateway}

	paymentOne.makePayment(1000)



}

// Method signature should be same as interface