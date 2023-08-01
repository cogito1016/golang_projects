package pointer

import (
	"fmt"
)

func Run(){
	a := 10;
	b := a;
	fmt.Println(a,b);

	a = 10;
	c := &a; // b = &a 할 시, cannot use &a (value of type *int) as int value in assignment
	fmt.Println(a,c); //10 0x140000140b8
	*c = 20;
	fmt.Println(a,c); //20 0x140000140b8
	fmt.Println(&a, &b); //0x140000140b8 0x140000140c0
	fmt.Println(a, b); //20 10
	fmt.Println(a, b); //20 10
	
	//a의 type을 알고 싶을때
	fmt.Printf("%T\n",a); //int
	fmt.Printf("%T\n",&a); //*int


	// &b = 0x140000140b8
}