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
	fmt.Println(add(22,23));
	fmt.Println(add01(22,23));
	fmt.Println(printReverse("Hello","World"));
	varTest();
	varTest01();
	varTest02();
	typeTransform();
	constant();
} 

func add(x int,y int) int{
	return x+y;
}

func add01(x,y int) int{
	return x+y;
}

func printReverse(x,y string) (string, string){
	return y,x;
}

func varTest(){
	var a bool;
	var b int;
	var c float32;
	var d float64;
	var e string;
	fmt.Println(a,b,c,d,e); //false 0 0 0 
}

func varTest01(){
	var a = "Hello?"
	var b = 1
	var c = 0.23
	//or var a,b,c = "Hello?",1,0.23;
	fmt.Printf("%T, %T, %T\n",a,b,c); //string, int, float64
}

func varTest02(){
	a := 1
	fmt.Printf("%T, %d\n",a,a); //int, 1
}

func typeTransform(){
	var i int = 42;
	fmt.Printf("%T, %d\n",i,i); //int, 42
	var f float32 = float32(i);
	fmt.Printf("%T, %f\n",f,f); //float32, 42.000000
	var u uint = uint(f);
	fmt.Printf("%T, %d\n",u,u); //uint, 42
}

func constant(){
	const PI = 3.14;
	fmt.Printf("%T, %f\n",PI,PI);//float64, 3.140000
	const (
		Big = 1 << 100
		Small = Big >> 99
	)
}