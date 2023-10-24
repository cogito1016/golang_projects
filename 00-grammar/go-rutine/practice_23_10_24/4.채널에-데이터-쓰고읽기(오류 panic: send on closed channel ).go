package practice_23_10_24

import (
	"fmt"
	"sync"
)

func Run채널에데이터쓰고읽기() {
	/**
	1. 버퍼채널은 버퍼가 채워지면 채널을 닫는다.
	2. 닫으면 다음 동작을 수행한다.
	3. 버퍼채널이 아닌경우에는, 누가 데이터를 가져갈때까지 실행을 멈춘다.
	*/
	c := make(chan int, 1)

	var waitGroup sync.WaitGroup

	waitGroup.Add(1)
	go func(c chan int) {
		defer waitGroup.Done()
		writeToChannel(c, 10)
		fmt.Println("Exit.")
	}(c)

	fmt.Println("Read : ", <-c)

	_, ok := <-c
	if ok {
		fmt.Println("Channel is open")
	} else {
		fmt.Println("Channel is closed")
	}

	waitGroup.Wait()

	ch := make(chan bool)
	for i := 0; i < 5; i++ {
		go printer(ch)
	}
	n := 0
	for i := range ch {
		fmt.Println(i)
		if i == true {
			n++
		}
		if n > 2 {
			fmt.Println("n : ", n)
			close(ch)
			break
		}
	}

	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}
}

func writeToChannel(c chan<- int, x int) {
	c <- x
	close(c)
}

func printer(ch chan<- bool) {
	ch <- true
}
