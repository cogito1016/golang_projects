package go_rutine

import (
	"fmt"
	"runtime"
)

func Run() {
	MyInfo()
}

func MyInfo() {
	fmt.Println("You are using ", runtime.Compiler)    //gc
	fmt.Println("on a ", runtime.GOARCH, "machine")    // arm64
	fmt.Println("Using Go version", runtime.Version()) //1.20.6
}
