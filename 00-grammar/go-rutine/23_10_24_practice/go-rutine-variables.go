package _3_10_24_practice

import (
	"fmt"
	"runtime"
)

func RunGoRutineVariables() {
	fmt.Println("You are using ", runtime.Compiler, " ")  //컴파일러, gc | gccgo
	fmt.Println("on a ", runtime.GOARCH, " machine")      //아키텍처, ARM64 | ...
	fmt.Println("Using Go Version ", runtime.Version())   //고버전 go1.20.6 | ...
	fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0)) //동시에 실행할 수 있는 프로세서 수 10 | ...

	//GOMAXPROCS를 변경하는 방법
	//1. runtime.GOMAXPROCS(0)에 0이아닌 1보다 같거나 큰 수를 할당한다.
}
