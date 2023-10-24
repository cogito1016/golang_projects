package practice_23_10_24

import (
	"fmt"
	"time"
)

func Run고루틴을끝날때까지기다리기Before() {

	for i := 0; i < 15; i++ {
		go func(a int) {
			fmt.Printf("%d ", a)
		}(i)
	}

	time.Sleep(time.Second) // 슬립을걸어 임시조치는했으나, 프로그램이 고루틴을 기다린것은 아님
	fmt.Println("Exsiting...")
}
