package greedy_problem

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func RunBoj10162() {
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	T, _ := strconv.Atoi(scanner.Text())

	if T%10 != 0 {
		fmt.Fprintf(writer, "-1")
		return
	}

	fmt.Fprintf(writer, strconv.Itoa(T/300)+" ")
	T = T % 300
	fmt.Fprintf(writer, strconv.Itoa(T/60)+" ")
	T = T % 60
	fmt.Fprintf(writer, strconv.Itoa(T/10)+" ")
	T = T % 10
}
