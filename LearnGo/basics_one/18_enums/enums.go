package main

import "fmt"

type OrderStatus int

const (
	Pending OrderStatus = iota
	Processing
	Shipped
	Delivered
)

func changeOrderStatus(status OrderStatus) {

	fmt.Println("Order status is: ", status)

}

type OrderStatusStr string

const (
	PendingStr    OrderStatusStr = "Pending"
	ProcessingStr OrderStatusStr = "Processing"
	ShippedStr    OrderStatusStr = "Shipped"
	DeliveredStr  OrderStatusStr = "Delivered"
)

func changeOrderStatusStr(status OrderStatusStr) {

	fmt.Println("Order status is: ", status)

}

func main() {

	changeOrderStatus(Pending)
	changeOrderStatus(Delivered)

	changeOrderStatusStr(PendingStr)
	changeOrderStatusStr(ProcessingStr)
	changeOrderStatusStr(ShippedStr)
	changeOrderStatusStr(DeliveredStr)

}
