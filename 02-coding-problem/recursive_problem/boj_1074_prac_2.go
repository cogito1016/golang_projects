package recursive_problem

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Z
func RunBoj1074Prac02() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := strings.Fields(scanner.Text())
	N, _ := strconv.Atoi(input[0])
	r, _ := strconv.Atoi(input[1])
	c, _ := strconv.Atoi(input[2])

	result := zOrderTraversalPrac02(N, r, c)
	fmt.Println(result)
}

func zOrderTraversalPrac02(n int, r int, c int) int {
	result := 0

	for n > 0 {
		size := 1 << uint(n-1)

		if r < size && c < size {

		} else if r < size && c >= size {
			//2
			c = c - size
			result = result + size*size
		} else if r >= size && c < size {
			r = r - size
			result = result + size*size*2
		} else if r >= size && c >= size {
			r = r - size
			c = c - size
			result = result + size*size*3
		}

		n--
	}

	return result
}
