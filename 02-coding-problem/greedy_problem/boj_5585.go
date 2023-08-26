package greedy_problem

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func RunBoj5585() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())

	charge := 1000 - N
	billCount := 0

	count, newCharge := chargeSolution(charge, 500)
	billCount += count
	count, newCharge = chargeSolution(newCharge, 100)
	billCount += count
	count, newCharge = chargeSolution(newCharge, 50)
	billCount += count
	count, newCharge = chargeSolution(newCharge, 10)
	billCount += count
	count, newCharge = chargeSolution(newCharge, 5)
	billCount += count
	count, newCharge = chargeSolution(newCharge, 1)
	billCount += count

	fmt.Println(billCount)
}

func chargeSolution(charge, unit int) (int, int) {
	billCount := charge / unit
	newCharge := charge - billCount*unit
	return billCount, newCharge
}
