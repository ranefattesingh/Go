package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	nAssemblies = 3
	partList    = []string{"A", "B", "C"}
	wg          sync.WaitGroup
)

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= nAssemblies; i++ {
		fmt.Println("Begin assembly cycle", i)
		wg.Add(len(partList))
		for _, part := range partList {
			go worker(part)
		}
		wg.Wait()
		fmt.Println("End assembly cycle", i)
	}
}

func worker(part string) {
	fmt.Println(part, "work is started.")
	time.Sleep(time.Duration(rand.Int63n(1e6)))
	fmt.Println(part, "work is completed.")
	wg.Done()
}
