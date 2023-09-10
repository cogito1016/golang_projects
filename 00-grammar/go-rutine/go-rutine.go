package go_rutine

import (
	"fmt"
	"runtime"
	"time"
)

func RunBasic() {
	//MyInfo()
	//GoRutinesInfo()

	//고루틴의 순서는 무작위
	//go TestGoRutine(10)
	//go Printme(15)
	//LazyDone()

	for i := 0; i < 10; i++ {
		go func(x int) { fmt.Printf("%d ", x) }(i)
	} //0 9 4 7 8 5 3 1 2 6 Exiting...
	lazyDone()
}

func myInfo() {
	fmt.Println("You are using ", runtime.Compiler)    //gc
	fmt.Println("on a ", runtime.GOARCH, "machine")    // arm64
	fmt.Println("Using Go version", runtime.Version()) //1.20.6
}

func goRutinesInfo() {
	fmt.Printf("GOMAXPROCS : %d\n", runtime.GOMAXPROCS(0)) //인자가 0이면 현재값, 다른수면 변경
}

func testGoRutine(x int) {
	fmt.Printf("%d ", x)
}

func printme(x int) {
	fmt.Printf("%d", x)
}

func lazyDone() { //Go프로그램은 Go루틴이 끝날때까지 기다리지않기때문에 출력문을 확인하기위해 일부러 지연시키는 것
	time.Sleep(time.Second)
	fmt.Println("Exiting...")
}
