package main

import (
	"fmt"
	"time"
)

func main() {
	ordersChan := make(chan Order)
	resultsChan := make(chan string)

	go supplier(ordersChan)
	go warehouse(ordersChan, resultsChan)

	for msg := range resultsChan {
		fmt.Println(msg)
	}
}

type Order struct {
	ID       int
	Quantity int
}

func supplier(ordersChan chan<- Order) {
	for i := 1; i <= 5; i++ {
		order := Order{ID: i, Quantity: 1}
		ordersChan <- order
		time.Sleep(300 * time.Millisecond)
	}
	close(ordersChan)
}

func warehouse(ordersChan <-chan Order, resultsChan chan<- string) {
	stock := 10

	for order := range ordersChan {
		if order.Quantity <= stock {
			stock -= order.Quantity
			time.Sleep(500 * time.Millisecond)
			resultsChan <- fmt.Sprintf("Заказ %d обработанных, осталось %d элементов", order.ID, stock)
		} else {
			resultsChan <- fmt.Sprintf("Заказ %d отклонен: недостаточно товаров", order.ID)
		}
	}
	close(resultsChan)
}
