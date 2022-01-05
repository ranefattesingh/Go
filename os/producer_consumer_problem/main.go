package main

import (
	"flag"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

func main() {

	// profile flags
	cpuprofile := flag.String("cpuprofile", "", "write cpu profile to `file`")
	memprofile := flag.String("memprofile", "", "write memory profile to `file`")

	// get maximum number of messages from flag
	max := flag.Int("n", 5, "define the number of messages")
	flag.Parse()

	// utilize max number of cores available
	runtime.GOMAXPROCS(runtime.NumCPU())

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("Could not create cpu profile", err)
		}

		if err = pprof.StartCPUProfile(f); err != nil {
			log.Fatal("Could not start cpu profile", err)
		}

		defer pprof.StopCPUProfile()
	}

	msgs := make(chan int)  // Channel to send messages
	done := make(chan bool) // Channel to control production of messages

	// Start a goroutine for Produce.produce
	go NewProducer(&msgs, &done).produce(*max)

	// Start a goroutine for Consumer.consume
	go NewConsumer(&msgs).consume()

	// Finish the program when the production is done
	<-done

	// Memory profile
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("Could not create memory profile", err)
		}
		runtime.GC() // get up to date statistics

		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("Could not write memory profile", err)
		}
		f.Close()
	}
}
