package practice_23_10_24

import (
	"fmt"
	"sync"
)

func Run고루틴을끝날때까지기다리기After() {

	var waitGroup sync.WaitGroup
	fmt.Println(waitGroup)

	for i := 0; i < 15; i++ {
		waitGroup.Add(1)
		go func(a int) {
			defer waitGroup.Done()
			fmt.Printf("%d ", a)
		}(i)
	}

	waitGroup.Wait()
	fmt.Println("Exsiting...")
}
