package go_rutine

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func RunNIL채널기본() {
	c := make(chan int)
	rand.Seed(time.Now().Unix())
	wg.Add(1)
	go add(c)
	go send(c)
	wg.Wait()
}

func add(c chan int) {
	sum := 0
	t := time.NewTimer(time.Second)

	for {
		select {
		case input := <-c:
			sum = sum + input
		case <-t.C:
			c = nil
			fmt.Println(sum)
			wg.Done()
		}
	}
}

func send(c chan int) {
	for {
		c <- rand.Intn(10)
	}
}
