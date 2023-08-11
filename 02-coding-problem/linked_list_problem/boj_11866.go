package linked_list_problem

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func RunBoj11866() {
	writer := bufio.NewWriter(os.Stdout)

	var N, K int
	fmt.Scan(&N, &K)

	// Create a circular linked list
	linked := make([]int, N)
	for i := 0; i < N; i++ {
		linked[i] = i + 1
	}

	writer.WriteString("<")

	idx := 0

	for len(linked) > 0 {
		idx = (idx + K - 1) % len(linked)
		writer.WriteString(strconv.Itoa(linked[idx]))

		linked = append(linked[:idx], linked[idx+1:]...)

		if len(linked) > 0 {
			writer.WriteString(", ")
		}
	}

	writer.WriteString(">")
	writer.Flush()
}





