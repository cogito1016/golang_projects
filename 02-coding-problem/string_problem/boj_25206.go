package string_problem

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RunBoj25206() {

	pointSum := 0.0
	pointAndGradeSum := 0.0

	for i := 0; i < 20; i++ {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := strings.Fields(scanner.Text())
		point, _ := strconv.ParseFloat(input[1], 32)
		grade := input[2]
		if grade == "P" {
			continue
		}

		pointSum += point
		pointAndGradeSum += point * getPointByGrade(grade)
	}

	fmt.Printf("%f", pointAndGradeSum/pointSum)

}
func getPointByGrade(grade string) float64 {
	eng := grade[0]
	var point float64

	if eng == 'F' {
		return 0.0
	} else if eng == 'A' {
		point = 4.0
	} else if eng == 'B' {
		point = 3.0
	} else if eng == 'C' {
		point = 2.0
	} else {
		point = 1.0
	}

	if len(grade) == 2 {
		weight := grade[1]

		if weight == '+' {
			point += 0.5
		}
	}

	return point
}
