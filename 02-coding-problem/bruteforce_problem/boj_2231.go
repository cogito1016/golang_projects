package bruteforce_problem

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func RunBoj2231() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())

	result := 0

	for i := 1; i <= N; i++ {
		sum := i

		str := strconv.Itoa(i)

		for j := 0; j < len(str); j++ {
			num, _ := strconv.Atoi(string(str[j]))
			sum += num
		}

		if sum == N {
			result = i
			break
		}
	}

	fmt.Println(result)
}
