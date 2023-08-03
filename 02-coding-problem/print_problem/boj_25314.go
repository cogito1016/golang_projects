package print_problem

import "fmt"

func RunBoj25314(){
	var N int;
	var result string;

	fmt.Scan(&N);

	share := N/4;

	for i := 0; i < share; i++ {
		result += "long ";
	}

	if N%4 != 0 {
		result = result + "long ";
	}

	result = result+"int";

	fmt.Println(result);
}