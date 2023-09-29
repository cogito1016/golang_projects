package greedy_problem

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func RunBoj1439() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()

	counter := make([]int, 2)

	//first setting
	flag := 1
	firstNum, _ := strconv.Atoi(string(str[0]))
	if firstNum == 0 {
		flag = 0
	}
	counter[flag]++

	for i := 1; i < len(str); i++ {
		value := str[i]
		target, _ := strconv.Atoi(string(value))
		if flag != target {
			counter[target]++
			flag = target
		}
	}

	fmt.Println(math.Min(float64(counter[0]), float64(counter[1])))
}
