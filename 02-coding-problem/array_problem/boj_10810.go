package array_problem

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func RunBoj10810(){
	scanner := bufio.NewScanner(os.Stdin);
	writer := bufio.NewWriter(os.Stdout);

	scanner.Scan();
	input := strings.Split(scanner.Text(), " ");
	N,_ := strconv.Atoi(input[0]);
	M,_ := strconv.Atoi(input[1]);

	arr := make([]int,N+1);

	for i := 0; i < M; i++ {
		scanner.Scan();
		input := strings.Split(scanner.Text()," ");
		a,_ := strconv.Atoi(input[0]);
		b,_ := strconv.Atoi(input[1]);
		ball,_ := strconv.Atoi(input[2]);

		for a<=b {
			arr[a] = ball;
			arr[b] = ball;
			a++;
			b--;
		}
	}

	for i := 1; i <= N; i++ {
		writer.WriteString(strconv.Itoa(arr[i])+" ");
	}

	writer.Flush();
}