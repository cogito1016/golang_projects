package print_problem

import "fmt"

func RunBoj1002(){
	//3개의 정수를 10 11 12 처럼 한줄로 입력받는 방법
	var a, b, c int;
	fmt.Scan(&a, &b, &c);

	var result int;
	result = a + b + c;

	fmt.Println(result);
}