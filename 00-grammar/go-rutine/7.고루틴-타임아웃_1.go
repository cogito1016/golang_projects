package go_rutine

import (
	"fmt"
	"time"
)

/*
*
Result
c2 ok
Time out
*/
func Run고루틴타임아웃_1() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "c2 ok"
	}()

	go func() {
		time.Sleep(8 * time.Second)
		c1 <- "c1 ok"
	}()

	for {
		select {
		case res := <-c2:
			fmt.Println(res)
		case res := <-c1:
			fmt.Println(res)
		case <-time.After(4 * time.Second):
			fmt.Println("Time out")
			return
		}
	}
}
