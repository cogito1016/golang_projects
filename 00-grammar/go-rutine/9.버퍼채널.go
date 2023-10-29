package go_rutine

import (
	"fmt"
	"time"
)

func Run버퍼채널기본() {
	numbers := make(chan int, 5) //크기가 5인 버퍼채널, 5개보다 많은 정수를 저장할 수 없다.

	counter := 10

	for i := 0; i < counter; i++ {
		time.Sleep(time.Second)
		select {
		//처리
		case numbers <- i * i:
			fmt.Println("About to process", i)
		default:
			fmt.Print("No space for ", i, " ")
		}
	}

	fmt.Println()

	for {
		select {
		case num := <-numbers:
			fmt.Print("*", num, " ")
		default:
			fmt.Println("Nothing left to read!")
			return
		}
	}
}
