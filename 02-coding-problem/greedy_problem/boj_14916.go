package greedy_problem

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func RunBoj14916() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text()) //거스름돈액수
	fiveCoinCount := 0
	twoCoinCount := 0

	//2원과 5원짜리로만 + 거스름돈 개수 최소화
	fiveCoinCount += N / 5
	N = N % 5
	twoCoinCount += N / 2
	N = N % 2

	if N != 0 && fiveCoinCount > 0 {
		fiveCoinCount--
		N += 5
		twoCoinCount += N / 2
		N = N % 2
	}

	if N != 0 {
		fmt.Println(-1)
	} else {
		fmt.Println(fiveCoinCount + twoCoinCount)
	}

}
