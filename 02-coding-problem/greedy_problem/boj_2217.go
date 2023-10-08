package greedy_problem

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

/*
10 15 일 때, 10 10 = 20 이 최댓값
13 15 일 때, 13 13 = 26 이 최댓값
1 10 15 일 때, 10 10 = 20 이 최댓값

1 10 15 일 때,
1 1  1 = 3
0 10 10 = 20
0 0  15 = 15
이므로, 20이 최댓값
*/
func RunBoj2217() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())

	arr := make([]int, N)

	for i := 0; i < N; i++ {
		scanner.Scan()
		target, _ := strconv.Atoi(scanner.Text())

		arr[i] = target
	}
	sort.Ints(arr)

	max := 0

	for i := 0; i < N; i++ {
		maxTarget := (arr[i] * (N - i))
		if maxTarget > max {
			max = maxTarget
		}
	}

	fmt.Println(max)
}
