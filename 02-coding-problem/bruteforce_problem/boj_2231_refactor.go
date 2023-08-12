package bruteforce_problem

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func RunBoj2231Refactor() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	N, _ := strconv.Atoi(input)

	result := 0
	minumum := 9 * len(input)

	for i := N - minumum; i <= N; i++ {
		sum := i
		s := strconv.Itoa(i)

		for j := 0; j < len(s); j++ {
			one, _ := strconv.Atoi(string(s[j]))
			sum += one
		}

		if sum == N {
			result = i
			break
		}
	}

	fmt.Println(result)
}
