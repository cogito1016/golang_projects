package practice_23_10_30

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

/*
*
Goroutines and Channels Basics:
- Create a program that uses two goroutines.
The first one should generate numbers from 1 to 10 and send them to the second goroutine using a channel.
The second goroutine should square the received numbers and print the squared results.
*/
func RunChannelBasic() {
	ch := make(chan int)
	timeout := 20 * time.Second

	wg.Add(1)
	go func(c chan<- int) {
		defer close(c)
		for i := 1; i <= 10; i++ {
			c <- i
		}
	}(ch)

	go func(c chan int) {
		defer wg.Done()
		for {
			select {
			case input, ok := <-c: //닫힌거면 ok가 false
				if !ok {
					fmt.Println("Channel is closed")
					return
				}
				fmt.Println(input * input)
			case <-time.After(timeout): //첫 번째 고루틴이 충분히 끝나고도 남을 시간으로 타임아웃설정
				fmt.Println("Time over ")
				return
			}
		}
	}(ch)

	wg.Wait()
}
