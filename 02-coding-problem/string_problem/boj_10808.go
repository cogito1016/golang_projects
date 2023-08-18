package string_problem

import (
	"bufio"
	"os"
	"strconv"
)

func RunBoj10808() {
	writer := bufio.NewWriter(os.Stdout)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()
	//a 97 z는 122
	arr := make([]int, 26)

	for i := 0; i < len(str); i++ {
		idx := str[i] - 97
		arr[idx]++
	}

	for i := 0; i < 26; i++ {
		writer.WriteString(strconv.Itoa(arr[i]) + " ")
	}

	writer.Flush()
}
