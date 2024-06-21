package main

import (
	"fmt"
	"runtime"
	"time"
)

// logic of myworker incapsulated to separated func
func myworker(num int, job <-chan string) {

	for p := range job {
		fmt.Printf("worked %d started. got signal: %s\n", num, p)
		time.Sleep(1 * time.Second)
	}
	fmt.Printf("worker %d done. ", num)

}

func pooling() {
	//func main() {

	// 1. create channel
	// 2. define pool (goroutines number)
	// 3. iterate through this pool
	// 4. create separate goroutines according the pool

	ch := make(chan string)

	// size of our pool
	g := runtime.NumCPU()

	// iterate through our pool
	for e := 0; e < g; e++ {

		// start separeate goroutines:
		/*
			go func(num int) {
				// iterate through the channel:
				for p := range ch {
					fmt.Printf("worker %d started... signal received: %s\n", num, p)

					//time.Sleep(1 * time.Second)

				}
				fmt.Printf("worker %d ended...", num)
			}(e)

		*/
		go myworker(e, ch)

	}

	const work = 5

	for w := 0; w < work; w++ {
		ch <- "some value"
		fmt.Println("value sent to the channel: ", w)
	}

	close(ch)

	fmt.Println("sent signal for termination")

}
