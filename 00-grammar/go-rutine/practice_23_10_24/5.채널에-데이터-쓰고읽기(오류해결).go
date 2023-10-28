package practice_23_10_24

import (
	"fmt"
)

/*
*
해결이된 코드이므로 go run --race run.go 에도 아무런 오류코드가 발생되지 않는다.
*/
func Run채널에데이터쓰고읽기해결() {
	ch := make(chan bool)

	go printer_fixed(ch, 5)

	for i := range ch { // 채널이 printer_fixed 때 닫혔기 때문에 알아서 빠져나온다.
		fmt.Println(i, " ")
	}

	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}
}

func printer_fixed(ch chan<- bool, times int) {
	for i := 0; i < times; i++ {
		ch <- true
	}
	close(ch)
}
