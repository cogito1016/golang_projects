package recursive_problem

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// #피보나치 #재귀
func RunBoj2747() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())

	fmt.Println(fibo_for(N))
}

func fibo_for(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	first := 0
	second := 1
	result := 0

	for i := 2; i <= n; i++ {
		result = first + second
		first = second
		second = result
	}

	return result
}

// 시간초과이므로 쓰지않습니다.
func fibo_recursive(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return fibo_recursive(n-1) + fibo_recursive(n-2)
}
