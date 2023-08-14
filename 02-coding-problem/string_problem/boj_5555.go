package string_problem

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func RunBoj5555() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	target := scanner.Text()
	targetLength := len(target)
	result := 0
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < N; i++ {
		scanner.Scan()
		ringString := scanner.Text()
		ringStringLength := len(ringString)
		targetCharIdx := 0
		targetChar := string(target[targetCharIdx])
		correctCounter := 0
		turnFlag := false

		for j := 0; j < ringStringLength; j++ {

			if targetChar == string(ringString[j]) {
				targetCharIdx++
				correctCounter++

				if correctCounter == targetLength {
					break
				}

				if j == ringStringLength-1 {
					turnFlag = true
					j = -1
				}
				targetChar = string(target[targetCharIdx])
			} else if turnFlag {
				break
			}
		}

		if correctCounter == targetLength {
			result++
		}
	}

	fmt.Println(result)
}
