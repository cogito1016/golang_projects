package array_problem

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RunBoj2566() {
	var x, y, max int

	scanner := bufio.NewScanner(os.Stdin)
	for i := 1; i <= 9; i++ {
		scanner.Scan()
		line := strings.Fields(scanner.Text())

		for j := 1; j <= 9; j++ {
			N, _ := strconv.Atoi(line[j-1])
			if N >= max {
				max = N
				y = i
				x = j
			}
		}
	}

	fmt.Println(max)
	fmt.Printf("%d %d", y, x)
}
