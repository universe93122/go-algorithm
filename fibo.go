package main

import (
	"fmt"
	"time"

	"golang.org/x/tour/tree"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(time.Second) // simulate work
		results <- job * 2
	}
}

func Fibo(n int) int {
	if n <= 2 {
		return n
	}
	return Fibo(n-2) + Fibo(n-1)
}

func ping(pingCh chan string, pongChn chan string) {
	for {
		msg := <-pingCh
		fmt.Println("Ping Received", msg)
		time.Sleep(500 * time.Millisecond)
		pongChn <- "Pong"
	}
}

func pong(pingCh chan string, pongChn chan string) {
	for {
		msg := <-pongChn
		fmt.Println("Ping Received", msg)
		time.Sleep(500 * time.Millisecond)
		pingCh <- "Ping"
	}
}

func Work() {
	const numJobs = 5
	const numWorkers = 3

	jobs := make(chan int, numJobs)
	results := make(chan int, numWorkers)

	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		fmt.Println("Result:", a, <-results)
	}

}
func main() {

	Reder()

	fmt.Println("companre trees:", Same(tree.New(1), tree.New(1)))

}
