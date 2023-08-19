package dynamic_programming_problem

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func RunBoj24416() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())

	fmt.Println(strconv.Itoa(recursive(N)) + " " + strconv.Itoa(dynamic(N)))
}

func recursive(N int) int {
	if N == 1 || N == 2 {
		return 1
	}
	return recursive(N-1) + recursive(N-2)
}

func dynamic(N int) int {
	arr := make([]int, N)
	arr[0] = 1
	arr[1] = 1

	count := 0
	for i := 2; i < N; i++ {
		arr[i] = arr[i-2] + arr[i-1]
		count++
	}
	return count
}
