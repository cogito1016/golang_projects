package main

import (
	"fmt"
	"math"
)

//factored import statement

func main() {

	/*
	GO에서는 대문자로 시작하는 이름이 Export된다.
	즉 아래구문에서 Printf와 Sqrt는 각각 fmt와 math패키지에서 export된것이다.
	패키지를 참조할때는 export된것들만 참조가가능하다.
	즉, 대문자로 시작하는 이름만 참조할 수 있다.
	*/

	fmt.Printf("Hello WOrld %g ", math.Sqrt(7));
} 