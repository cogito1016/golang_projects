package array_problem

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RunBoj2738(){
	scanner := bufio.NewScanner(os.Stdin);
	writer := bufio.NewWriter(os.Stdout);

	scanner.Scan();
	input := strings.Split(scanner.Text(), " ");
	N,_ := strconv.Atoi(input[0]);
	M,_ := strconv.Atoi(input[1]);

	arr := make([][]int,N);

	for i := 0; i < N; i++ {
		arr[i] = make([]int,M);
		scanner.Scan();
		input = strings.Split(scanner.Text(), " ");
		for j := 0; j < M; j++ {
			val,_ := strconv.Atoi(input[j]);
			arr[i][j] = val;
		}
	}

	for i := 0; i < N; i++ {
		scanner.Scan();
		input = strings.Split(scanner.Text(), " ");
		for j := 0; j < M; j++ {
			val,_ := strconv.Atoi(input[j]);
			arr[i][j] += val;
		}
	}

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			val := strconv.FormatInt(int64(arr[i][j]),10);
			fmt.Fprintf(writer,val+" ");
		}
		fmt.Fprintf(writer,"\n");
	}

	writer.Flush();
}