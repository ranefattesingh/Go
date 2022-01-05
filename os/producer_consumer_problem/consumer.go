package main

import "fmt"

type Consumer struct {
	msgs *chan int
}

func NewConsumer(msgs *chan int) *Consumer {
	return &Consumer{msgs: msgs}
}

func (c *Consumer) consume() {
	fmt.Println("consume: STARTED")
	for {
		msg := <-*c.msgs
		fmt.Println("consume: RECEIVED", msg)
	}
}
