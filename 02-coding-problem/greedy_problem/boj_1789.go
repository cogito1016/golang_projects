package greedy_problem

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func RunBoj1789() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	S, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	var result int64
	var count int64
	var i int64

	for S >= result {
		count++
		i++
		result += i
	}

	if result > S {
		count -= 1
	}

	fmt.Println(count)
}
