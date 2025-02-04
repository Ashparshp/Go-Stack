package main

import (
	"fmt"
	"time"
)


type order struct {
	orderId    int
	customerId int
	amount     float64
	status     string
	createdAt  time.Time
}

func (o *order) changeStatus(status string) {
	o.status = status

}

func (o order) getAmount() float64 {
	return o.amount
}

func new (orderId int, customerId int, amount float64, status string, createdAt time.Time) *order {
	return &order{
		orderId:    orderId,
		customerId: customerId,
		amount:     amount,
		status:     status,
		createdAt:  createdAt,
	}
}

type customer struct {
	customerId int
	name       string
	email      string
}

type neworder struct {
	orderId   int
	amount    float64
	status    string
	createdAt string
	customer
}

func main() {

	order1 := order{
		orderId:    1,
		customerId: 1,
		amount:     100.0,
		status:     "Pending",
		createdAt:  time.Now(),
	}

	order2 := order{
		orderId:    2,
		customerId: 2,
		amount:     200.0,
		status:     "Completed",
		createdAt:  time.Now(),
	}

	order3 := order{
		orderId:    3,
		customerId: 3,
		amount:     300.0,
		status:     "Cancelled",
		createdAt:  time.Now().Local().Truncate(time.Minute),
	}

	orders := []order{order1, order2, order3}

	for _, order := range orders {
		fmt.Println("Order ID:", order.orderId)
		fmt.Println("Customer ID:", order.customerId)
		fmt.Println("Amount:", order.amount)
		fmt.Println("Status:", order.status)
		fmt.Println("Created At:", order.createdAt)
	}

	order1.changeStatus("Completed")
	fmt.Println("Order 1 Status:", order1.status)

	fmt.Println("Order 1 Amount:", order1.getAmount())

	order4 := new(4, 4, 400.0, "Pending", time.Now())
	fmt.Println("Order 4:", order4)

	language := struct {
		name string
		version float64
	}{
		name: "Go",
		version: 1.15,
	}

	fmt.Println("Language:", language)

	order5 := neworder{
		orderId:   5,
		amount:    500.0,
		status:    "Pending",
		createdAt: time.Now().String(),
		customer: customer{
			customerId: 4,
			name:       "John Doe",
			email:      "ashparsh@gmail.com",
		},
	}

	fmt.Println("Order 5:", order5)
}