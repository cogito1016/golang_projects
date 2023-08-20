package implement_problem

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func RunBoj1475() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()

	checkBoard := make([]int, 10)
	set := 1

	for i := 0; i < len(str); i++ {
		num, _ := strconv.Atoi(string(str[i]))

		if checkBoard[num] < set {
			checkBoard[num]++
			continue
		}

		if checkBoard[num] >= set {

			if num == 9 && checkBoard[6] < set {
				checkBoard[6]++
				continue
			}

			if num == 6 && checkBoard[9] < set {
				checkBoard[9]++
				continue
			}
			checkBoard[num]++
			set++
		}
	}

	fmt.Println(set)
}
