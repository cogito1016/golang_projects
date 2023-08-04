package array_statement

import "fmt"
 
func AppendExam(){
	var arr []int;

	for i := 0; i < 10; i++ {
		arr = append(arr,i+1);
	}

	fmt.Println(arr);//[1 2 3 4 5 6 7 8 9 10]
	fmt.Println(arr[0:5]);//[1 2 3 4 5]
	fmt.Println(arr[:3]);//[1 2 3]
	fmt.Println(arr[2:3]);//[3]
	fmt.Println(arr[3:]);// [4 5 6 7 8 9 10]

	arr[0] = arr[9];
	fmt.Println(arr);
}