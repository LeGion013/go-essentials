package main

/*

-------     -------     -------
|step1| --> |step2| --> |step3|
-------     -------     -------

if step2 can be broken down into multiple steps,
we can run them in parallel. This process of
breaking down the step is called FAN OUT:

            -------
            |step1|
            -------
               ↓
   -------------------------     ==> FAN OUT (step2 broke down to several steps)
   ↓           ↓           ↓
-------     -------     -------
|step2| --> |step2| --> |step2|
-------     -------     -------
   ↓           ↓           ↓
   -------------------------     ==> FAN IN
               ↓
            -------
            |step3|
            -------

// Explanation from Bill:
// FanOut: you are a manager and you hire one new employee for the exact amount
// of work you have to get done. Each new employee knows immediately what they
// are expected to do and starts their work. You sit waiting for all the results
// are received by you. No given employee needs an immediate guarantee that you
// received their result.

*/

import (
	"fmt"
	"math/rand"
	"time"
)

// func main() {
func fanout() {
	fmt.Println("app started...")

	// step1 will be broke down to n goroutines:
	gonums := 5

	// create channel which will collect all results from goroutines:
	ch := make(chan string, gonums)

	for g := 0; g <= gonums; g++ {
		go func(gonum int) {
			// simulation of work thourgh sleep
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
			// write/populate some data to channel / signal that work need to be done:
			ch <- "some data"
			fmt.Println("small task done and sent signal: ", gonum)
		}(g)
	}

	// iterate through results
	for gonums > 0 {
		// read the value from channel: blocking operation "waiting to work to come"
		val := <-ch

		gonums--

		fmt.Println(val)
		fmt.Println("channel got the message from routine: ", gonums)

	}

	time.Sleep(time.Second)
	fmt.Println("-----------------------------------------")

	fmt.Println("terminated...")
}
