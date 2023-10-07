package greedy_problem

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func RunBoj1026() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())

	A := make([]int, N)
	scanner.Scan()
	aStrs := strings.Fields(scanner.Text())
	for i := 0; i < N; i++ {
		A[i], _ = strconv.Atoi(aStrs[i])
	}

	B := make([]int, N)
	scanner.Scan()
	bStrs := strings.Fields(scanner.Text())
	for i := 0; i < N; i++ {
		B[i], _ = strconv.Atoi(bStrs[i])
	}

	sort.Ints(A)
	sort.Ints(B)

	result := 0
	for i := 0; i < N; i++ {
		result += A[i] * B[N-1-i]
	}

	fmt.Println(result)
}
