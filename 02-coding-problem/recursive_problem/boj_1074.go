package recursive_problem

import (
	"bufio"
	"fmt"
	"math"
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

	N = int(math.Pow(float64(2), float64(N)))
	z_recursive(N, 0, 0)
}

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
}
