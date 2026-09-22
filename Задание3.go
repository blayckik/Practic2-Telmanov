package main

import "fmt"

type Order struct {
	id          int
	items       []int
	total       float64
	address     string
	isCompleted bool
}

func addOrder(orders map[int]Order, newOrder Order) {
	orders[newOrder.id] = newOrder
}

func main() {
	orders := make(map[int]Order)

	order1 := Order{
		id:          1,
		items:       []int{101, 102, 103},
		total:       2500.50,
		address:     "ул. Пушкина, д. 32",
		isCompleted: false,
	}

	addOrder(orders, order1)
	fmt.Printf("%+v\n", orders)
}