package data_structure_problem

import (
	"bufio"
	"os"
	"strconv"
)

func RunBoj2161() {
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())

	arr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = i + 1
	}

	for {
		if len(arr) == 1 {
			str := strconv.Itoa(arr[0])
			writer.WriteString(str)
			break
		}
		str := strconv.Itoa(arr[0])
		writer.WriteString(str + " ")
		arr = arr[1:]
		arr = append(arr, arr[0:1]...)
		arr = arr[1:]

	}

	//if N%2 == 0 {
	//	start := 1
	//	weight := 2
	//	count := 0
	//
	//	for count < N {
	//		for i := start; i <= N; i = i + weight {
	//			str := strconv.Itoa(i)
	//			writer.WriteString(str + " ")
	//			count++
	//			if count >= N {
	//				break
	//			}
	//		}
	//		start = start * 2
	//		weight = weight * 2
	//	}
	//} else {
	//}

}
