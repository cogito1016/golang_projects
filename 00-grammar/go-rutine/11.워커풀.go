package go_rutine

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type Client struct {
	id      int
	integer int
}

type Result struct {
	job    Client
	square int
}

var size = runtime.GOMAXPROCS(0)
var clients = make(chan Client, size)
var data = make(chan Result, size)

func Run워커풀기본() {
	nJobs := 10   //작업수
	nWorkers := 2 // 워커수

	go create(nJobs)
	finished := make(chan interface{})

	go func() {
		for d := range data {
			fmt.Printf("Client ID : %d\tint: ", d.job.id)
			fmt.Printf("%d\tsquare: %d\n", d.job.integer, d.square)
		}
		finished <- true
	}()

	var wg sync.WaitGroup
	for i := 0; i < nWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(data)
}

func worker(wg *sync.WaitGroup) {
	for c := range clients {
		square := c.integer * c.integer
		output := Result{c, square}
		data <- output
		time.Sleep(time.Second)
	}
	wg.Done()
}

func create(n int) {
	for i := 0; i < n; i++ {
		c := Client{i, i}
		clients <- c
	}
	close(clients)
}
