package go_rutine

import (
	"fmt"
	"time"
)

func Run고루틴은순서보장이안된다() {
	go func(x int) {
		fmt.Println(x)
	}(10)

	go printMe(15)

	time.Sleep(time.Second)

	/**
	Result 1.
	 +  15
	10
	*/

	/**
	Result 2.
	10
	 +  15
	*/
}

func printMe(x int) {
	fmt.Println(" + ", x)
}
