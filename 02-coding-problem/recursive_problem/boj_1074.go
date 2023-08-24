package recursive_problem

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Z

var targetY int
var targetX int
var result int

func RunBoj1074() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := strings.Fields(scanner.Text())
	N, _ := strconv.Atoi(input[0])
	r, _ := strconv.Atoi(input[1])
	c, _ := strconv.Atoi(input[2])
	targetY = r
	targetX = c
	result = 0

	result := zOrderTraversal(N, r, c)
	fmt.Println(result)
}

func zOrderTraversal(N, r, c int) int {
	result := 0

	for N > 0 {
		size := 1 << uint(N-1)

		if r < size && c < size {

		} else if r < size && c >= size {
			result += size * size
			c -= size
		} else if r >= size && c < size {
			result += 2 * size * size
			r -= size
		} else {
			result += 3 * size * size
			r -= size
			c -= size
		}

		N--
	}

	return result
}

//정석적인방법, 그러나 최악의경우 시간초과가 발생한다.
func z_recursive(n int, y int, x int) {
	if n == 2 {
		if y == targetY && x == targetX {
			fmt.Println(result)
			return
		}
		result++
		if y == targetY && x+1 == targetX {
			fmt.Println(result)
			return
		}
		result++
		if y+1 == targetY && x == targetX {
			fmt.Println(result)
			return
		}
		result++
		if y+1 == targetY && x+1 == targetX {
			fmt.Println(result)
			return
		}
		result++
		return
	}

	z_recursive(n/2, y, x)
	z_recursive(n/2, y, x+n/2)
	z_recursive(n/2, y+n/2, x)
	z_recursive(n/2, y+n/2, x+n/2)
//}
