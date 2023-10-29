package go_rutine

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

var result = make(chan bool)

func Run고루틴타임아웃_2() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())
	duration := time.Duration(int32(N)) * time.Millisecond
	fmt.Printf("Timeout period is %s\n", duration)

	go timeout(duration)

	val := <-result
	if val {
		fmt.Println("Time out")
	} else {
		fmt.Println("OK") // <-temp구문이 없는데도 OK가 출력되는 이유는 temp채널이 닫혀 zerovalue인 false를 val값이 할당받았기 때문이다.
	}
}

func timeout(t time.Duration) {
	temp := make(chan int)
	go func() {
		time.Sleep(5 * time.Second)
		defer close(temp)
	}()

	select {
	case <-temp:
		result <- false
	case <-time.After(t):
		result <- true
	}
}
