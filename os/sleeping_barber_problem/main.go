package main

import (
	"fmt"
	"time"
)

func main() {
	b := NewBarber()
	b.name = "Rockey"
	WaitingRoom := make(chan *Customer, 5) // 5 Chairs
	Wakers := make(chan *Customer, 1)
	go b.barber(b, WaitingRoom, Wakers)

	time.Sleep(time.Millisecond * 100)
	wg = NewWaitGroup()
	n := 10
	wg.Add(n)

	// Spawn customers
	for i := 0; i < n; i++ {
		time.Sleep(time.Millisecond * 50)
		c := new(Customer)
		go c.customer(c, b, WaitingRoom, Wakers)
	}

	wg.Wait()
	fmt.Println("No more customers for the day")
}
