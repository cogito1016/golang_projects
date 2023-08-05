package struct_state

import "fmt"

type person struct {//구조체
	//생성자는없음
	name string
	age int
	favFood []string
}

func Run(){
	fmt.Println("Struct Test");
	favFood := []string{"apple"};
	jayden := person{name:"jayden",age:18,favFood:favFood}
	//jayden := person{"jayden",18,favFood} 이렇게 사용해도 됨
	fmt.Println(jayden);//{jayden 18 [apple]}
	fmt.Println(jayden.favFood);//[apple]
}