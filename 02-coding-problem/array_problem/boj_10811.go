package array_problem

import "fmt"

func Run(){
	var N int;
	var M int;
	var arr []int;

	fmt.Scan(&N, &M);

	for i := 0; i < N; i++ {
		arr = append(arr,i+1);
	}

	for i := 0; i < M; i++ {
		var a int;
		var b int;
		fmt.Scan(&a, &b);

		for (a!=b)&&(a<b){
			temp := arr[a-1];
			arr[a-1] = arr[b-1];
			arr[b-1] = temp;
			a++;
			b--;
		}
	}
	
	for i := 0; i < N; i++ {
		fmt.Print(arr[i]," ");
	}
}