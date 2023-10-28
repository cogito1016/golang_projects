package go_rutine

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
		//--race결과 3.Previous read at 0x00c000186010 by goroutine 9:
		//--race결과 4.Goroutine 9 (running) created at:
	}
	n := 0
	for i := range ch { //채널을 사용한 range문은 break나 채널이 닫힐경우에만 빠져나온다
		fmt.Println(i)
		if i == true {
			n++
		}
		if n > 2 {
			fmt.Println("n : ", n)
			close(ch) //--race결과 1.Write at 0x00c000186010 by main goroutine:
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
	ch <- true //--race결과 2.Previous read at 0x00c000186010 by goroutine 9:
}

/**
--race결과
채널을 닫는것과 채널에 쓰는 코드가 공존한다.
채널에 쓰는것이 고 루틴에 의해 실행된다.
고 루틴의 실행순서는 보장될 수 없으므로,
채널닫기, 채널쓰기의 순서도 보장할 수 없다.
따라서 경쟁상태이다.
이를 해결하기 위해서는 printer()함수의 구현을 변경해야한다.
*/
