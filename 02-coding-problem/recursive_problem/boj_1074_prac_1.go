package recursive_problem

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Z
func RunBoj1074Prac01() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := strings.Fields(scanner.Text())
	N, _ := strconv.Atoi(input[0])
	r, _ := strconv.Atoi(input[1])
	c, _ := strconv.Atoi(input[2])

	result := zOrderTraversalPrac01(N, r, c)
	fmt.Println(result)
}

func zOrderTraversalPrac01(n int, r int, c int) int {
	result := 0

	for n > 0 {
		size := 1 << uint(n-1)

		if r < size && c < size {

		} else if r < size && c >= size { //2사분면
			result += size * size //1사분면의 크기만큼 더해준다.
			c -= size             //2사분면으로 이동
		} else if r >= size && c < size { //3사분면
			result += 2 * size * size //1,2사분면의 크기만큼 더해준다.
			r -= size                 //3사분면으로 이동
		} else {
			result += 3 * size * size //1,2,3사분면의 크기만큼 더해준다.
			r -= size
			c -= size
		}

		n-- //n을 줄여준다.
	}

	return result
}
