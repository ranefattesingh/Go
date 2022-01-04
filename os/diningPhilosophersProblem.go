package main

import (
	"hash/fnv"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"
)

// Philosophers
var ph = []string{"Mark", "Russell", "Rocky", "Haris", "Root"}

const (
	HUNGER = 3                 // Number of times philosoper eats
	THINK  = time.Second / 100 // Mean time to think
	EAT    = time.Second / 100 // Mean time to eat
)

var dining sync.WaitGroup

var fmt = log.New(os.Stdout, "", 0)

func main() {
	fmt.Println("Table empty")
	dining.Add(5) // length of ph array
	fork0 := &sync.Mutex{}
	forkLeft := fork0

	for i := 1; i < len(ph); i++ {
		forkRight := &sync.Mutex{}
		go diningProblem(ph[i], forkLeft, forkRight)
		forkLeft = forkRight
	}
	go diningProblem(ph[0], fork0, forkLeft)
	dining.Wait()
	fmt.Println("Table empty")
}

func diningProblem(name string, dominantHand, otherHand *sync.Mutex) {
	fmt.Println("Seated:", name)

	h := fnv.New64a()
	h.Write([]byte(name))
	rg := rand.New(rand.NewSource(int64(h.Sum64()))) // Random Generator
	rSleep := func(t time.Duration) {
		time.Sleep(t/2 + time.Duration(rg.Int63n(int64(t))))
	}

	for h := HUNGER; h > 0; h-- {
		fmt.Println(name, "Hungry")
		dominantHand.Lock()
		otherHand.Lock()
		fmt.Println("Eating...")
		rSleep(EAT)
		dominantHand.Unlock()
		otherHand.Unlock()
		fmt.Println("Thinking...")
		rSleep(THINK)
	}

	fmt.Println(name, "Satisfied")
	dining.Done()
	fmt.Println(name, "Left the table")
}
