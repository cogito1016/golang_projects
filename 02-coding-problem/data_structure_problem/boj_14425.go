package data_structure_problem

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RunBoj14425() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := strings.Fields(scanner.Text())
	N, _ := strconv.Atoi(input[0])
	M, _ := strconv.Atoi(input[1])

	dictionary := make(map[string]bool)

	for i := 0; i < N; i++ {
		scanner.Scan()
		str := scanner.Text()
		dictionary[str] = true
	}

	result := 0

	for i := 0; i < M; i++ {
		scanner.Scan()
		str := scanner.Text()
		_, ok := dictionary[str]
		if ok {
			result++
		}
	}

	fmt.Println(result)
}
