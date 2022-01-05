package main

import "fmt"

type Producer struct {
	msgs *chan int
	done *chan bool
}

func NewProducer(msgs *chan int, done *chan bool) *Producer {
	return &Producer{msgs: msgs, done: done}
}

func (p *Producer) produce(max int) {
	fmt.Println("produce: STARTED")
	for i := 0; i < max; i++ {
		fmt.Println("produce: SENDING", i)
		*p.msgs <- i
	}
	*p.done <- true
	fmt.Println("produce: DONE")
}
