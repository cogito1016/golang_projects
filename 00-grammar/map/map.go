package map_state

import "fmt"

func Run(){
	fmt.Println("MapTest");

	jayden := map[string]string{"name":"jayden","age":"18"}
	fmt.Println(jayden);//map[age:18 name:jayden]

	for key,value := range jayden{
		fmt.Println(key, value);
		/**
		name jayden
		age 18
		**/
	}
}