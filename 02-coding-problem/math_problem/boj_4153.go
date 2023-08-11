package math_problem

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func RunBoj4153(){
	writer := bufio.NewWriter(os.Stdout);

	for {
		arr := make([]float64,3);
		fmt.Scan(&arr[0],&arr[1],&arr[2]);

		if arr[0] == 0{
			break;
		}

		sort.Float64s(arr);

		result := "wrong"
		if math.Pow(arr[2],2)==(math.Pow(arr[0],2)+math.Pow(arr[1],2)){
			result = "right"
		}
		
		writer.WriteString(result+"\n");
	}

	writer.Flush();

}