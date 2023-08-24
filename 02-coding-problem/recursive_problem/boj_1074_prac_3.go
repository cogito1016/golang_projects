package recursive_problem

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Z
func RunBoj1074Prac03() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := strings.Fields(scanner.Text())
	N, _ := strconv.Atoi(input[0])
	r, _ := strconv.Atoi(input[1])
	c, _ := strconv.Atoi(input[2])

	result := zOrderTraversalPrac03(N, r, c)
	fmt.Println(result)
}

func zOrderTraversalPrac03(N int, r int, c int) int {
	result := 0

	for N > 0 {
		size := 1 << uint(N-1)

		if r < size && c < size {

		} else if r < size && c >= size {
			c -= size
			result = result + size*size
		} else if r >= size && c < size {
			r -= size
			result = result + 2*size*size
		} else if r >= size && c >= size {
			c -= size
			r -= size
			result = result + 3*size*size
		}
		N--
	}

	return result
}
