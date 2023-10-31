package examples

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
*
Producer-Consumer Problem:
- Implement a program with two goroutines - a producer and a consumer.
The producer generates random numbers and sends them to the consumer through a channel.
The consumer should calculate the square of the received number and print it.
Ensure that the producer keeps generating numbers and the consumer keeps consuming them concurrently.
*/

var wg2 sync.WaitGroup

func RunProducerConsumer() {
	ch := make(chan int)

	wg2.Add(1)

	go producer(ch)
	go consumer(ch)

	wg2.Wait()
}

func producer(c chan<- int) {
	defer close(c)

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		c <- rand.Intn(10) + 1
	}
}

func consumer(c <-chan int) {
	defer wg2.Done()
	for {
		select {
		case input, ok := <-c:
			if !ok {
				fmt.Println("Channel Closed")
				return
			}
			fmt.Println(input * input)
		case <-time.After(20 * time.Second):
			fmt.Println("Time Over")
			return
		}
	}
}
