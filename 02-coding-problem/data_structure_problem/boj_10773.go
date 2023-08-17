package data_structure_problem

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func RunBoj10773() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	K, _ := strconv.Atoi(scanner.Text())

	var sum int
	arr := []int{}

	for i := 0; i < K; i++ {
		scanner.Scan()
		N, _ := strconv.Atoi(scanner.Text())

		if N == 0 {
			sum -= arr[len(arr)-1]
			arr = arr[:len(arr)-1]
			continue
		}

		arr = append(arr, N)
		sum += N
	}

	fmt.Println(sum)
}
