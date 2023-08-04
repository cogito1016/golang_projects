package basicarray

import "fmt"

func Run(){
	arr := [5]int{1,2,3}; //basic
	arr[3] = 4;
	arr[4] = 5;

	for idx,value := range arr{
		fmt.Println(idx,value);
	}
}