package go_rutine

import (
	"fmt"
	"sync"
)

func RunChannelBasic() {
	var waitGroup sync.WaitGroup

	//버퍼를 사용하는 채널
	c := make(chan int, 1)
	waitGroup.Add(1)
	go func(c chan int) {
		defer waitGroup.Done()
		writeToChannel(c, 10)
		fmt.Println("Exit.")
	}(c)

	fmt.Println("Read:,", <-c)

	waitGroup.Wait()

	//버퍼를사용하지않는 채널
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
			fmt.Println("n: ", n)
			close(ch) //채널을닫는다
			break
		}
	}

	for i := 0; i < 5; i++ {
		fmt.Println(<-ch) //falese 가 5번 출력된다, 이유? 닫힌채널에서 데이터를 가져오려면, 채널타입의 제로값을 가져온다.
	}
}

func writeToChannel(c chan int, x int) {
	c <- x
	close(c)
}

func printer(ch chan bool) {
	ch <- true
}

func printerRefactored(ch chan<- bool) { //매개변수로받는 채널을 쓰기전용으로 만든다.
	ch <- true
}

func writeToChannelRefactored(c chan<- int, x int) { //매개변수로받는 채널을 쓰기전용으로 만든다.
	fmt.Println("1", x)
	c <- x
	fmt.Println("2", x)
}

func f2(out <-chan int, in chan<- int) {
	x := <-out
	fmt.Println("Read (f2):", x)
	in <- x
	return
}
