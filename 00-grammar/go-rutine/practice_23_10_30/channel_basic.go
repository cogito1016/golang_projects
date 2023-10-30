package practice_23_10_30

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func RunChannelBasic() {
	ch := make(chan int)

	go func(c chan<- int) {
		time.Sleep(time.Second)
		wg.Add(1)
		c <- rand.Intn(10)
	}(ch)

	go func(c chan int) {
		for {
			select {
			case <-c:
				fmt.Println(<-c)
				wg.Done()
			case <-time.After(5 * time.Second):
				fmt.Println("Time over ")
			}
		}
	}(ch)

	wg.Wait()
}
