package go_rutine

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
*
Select는 Switch와 비슷하게 작동하여 Case와 Default를 설정할 수 있다.
CASE로 여러 채널을 다룬다.
*/
func RunSELECT키워드() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(1)
	go func() {
		gen(0, 2*n, createNumber, end)
		waitGroup.Done()
	}()
}

func gen(min, max int, createNumber chan int, end chan bool) {
	time.Sleep(time.Second)
	for {
		select {
		case createNumber <- rand.Intn(max-min) + min:
		case <-end:
			fmt.Println("Eended ! ")
		//return
		case <-time.After(4 * time.Second):
			fmt.Println("time.After()!")
			return
		}
	}
}
