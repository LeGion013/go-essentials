package main

import (
	"fmt"
	"runtime"
	"time"
)

// reading from the channel
// 1 range
// 2 select
func worker(workerNumber int, jobsCh <-chan int) {
	// will block untill we have something in the channel
	for job := range jobsCh {
		fmt.Printf("Worker %d starting...", workerNumber)

		// **** some work imitation
		time.Sleep(1 * time.Second)
		fmt.Printf("worker %d is done, res: %d\n", workerNumber, job)
	}
}

// func main() {
func fanout_hw() {

	// ***** we have some "work/step" that we can broke down to several goroutines*****
	jobs := make([]int, 0)

	for i := 1; i < 100; i++ {
		jobs = append(jobs, i)
	}

	fmt.Println(jobs)

	// ver 0 - no control over how many workers we run
	// wg := sync.WaitGroup{}
	// wg.Add(len(jobs))
	// for _, job := range jobs {
	// 	_ = job
	// 	go worker(job)
	// }
	// wg.Wait()
	// // exit

	// numWorkers := 100

	// **** here we define how many workers will be working simulteniously (goroutines on our pool)
	numWorkers := runtime.NumCPU()

	fmt.Printf("Number of workers: %d\n", numWorkers)

	jobsCh := make(chan int)

	// **** start 2 workers(numWorkers) for doing our work
	// pooling:
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobsCh)
	}

	// now we have N workers and they are idling
	for _, job := range jobs {
		// go worker(job)
		jobsCh <- job // bloocking if no workers
	}

	// close channel to prevent leaking
	close(jobsCh)

}
