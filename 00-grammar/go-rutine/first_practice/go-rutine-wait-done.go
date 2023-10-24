package first_practice

import (
	"fmt"
	"sync"
)

func RunWaitGoRutineDone() { // 고 루틴이 끝날때까지 프로그램을 대기시키는법에 대해

	var waitGroup sync.WaitGroup
	fmt.Println(waitGroup)

	for i := 0; i < 15; i++ {
		waitGroup.Add(1)
		go func(x int) {
			defer waitGroup.Done()
			fmt.Printf("%d ", x)
		}(i)
	}

	fmt.Printf("%#v\n", waitGroup)
	waitGroup.Wait()
	fmt.Printf("Exiting...")
}
