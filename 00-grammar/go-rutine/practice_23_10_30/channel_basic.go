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

	wg.Add(1)
	go func(c chan<- int) {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second)
			c <- rand.Intn(10)
		}
		c = nil //의도적으로 블록시킨다, close를 할 시 select input := <-c에서 제로값을 무수히많이 받아버림
	}(ch)

	go func(c chan int) {
		for {
			select {
			case input := <-c:
				fmt.Println(input)
			case <-time.After(10 * time.Second):
				fmt.Println("Time over ")
				wg.Done()
				return
			}
		}
	}(ch)

	wg.Wait()
}
