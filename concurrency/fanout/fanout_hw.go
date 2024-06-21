package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

/*
	1. Collect results in main (implement fan in)
	2. As work make multiplication * 2
	3. Implement Context

	Q&A
	1. time.Sleep(1 * time.Second) in "func worker" affect which workers involved in work
	2. in fan out what the point to use buffered channels?
*/

type resultSet struct {
	In  int
	Out int
}

func worker(workerNumber int, jobsCh <-chan int, resultCh chan<- int) {
	// will block untill we have something in the channel

	for job := range jobsCh {
		fmt.Printf("Worker %d starting...", workerNumber)
		// **** some work imitation
		time.Sleep(1 * time.Second)
		res := job * 2
		resultCh <- res

		fmt.Printf("worker %d is done, %d * 2 = %d\n", workerNumber, job, res)
	}
}

func main() {
	//func fanout_hw() {

	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()

	// ***** we have some "work/step" that we can broke down to several goroutines*****
	jobs := make([]int, 0)
	jobsResult := make([]int, 0)

	for i := 1; i < 10; i++ {
		jobs = append(jobs, i)
	}

	fmt.Println(jobs)

	// **** here we define how many workers will be working simulteniously (goroutines on our pool)
	numWorkers := runtime.NumCPU()

	fmt.Printf("Number of workers: %d\n", numWorkers)

	// Channels for jobs and results
	jobsCh := make(chan int)
	resultCh := make(chan int, len(jobs)) // should be length of input jobs

	// **** start 2 workers(numWorkers) for doing our work (pooling)
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobsCh, resultCh)
	}

	// now we have N workers and they are idling
	// send jobs:
	for _, job := range jobs {
		jobsCh <- job // bloocking if no workers
	}
	// close channel to prevent leaking
	close(jobsCh)

	//****************************begin-fan-in****************************
	// fan-in collent results:
	var wg sync.WaitGroup
	wg.Add(len(jobs)) // set counter to the num of jobs

	// launch goroutine to wait for all output from jobs
	go func() {
		wg.Wait()       // wait for all jobs to be done
		close(resultCh) //close the results channel after work is done
	}()

	// here we are process results:
	for result := range resultCh {
		//fmt.Println("results: ", result)

		jobsResult = append(jobsResult, result)
		wg.Done() // decrease the waitgroup counter when result processed
	}
	//****************************end-fan-in****************************

	fmt.Printf("input slice: %d\n", jobs)
	fmt.Printf("output slice: %d\n", jobsResult)

}
