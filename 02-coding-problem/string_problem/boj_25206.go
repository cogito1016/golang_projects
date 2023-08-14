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
	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < 20; i++ {
		scanner.Scan()
		input := strings.Fields(scanner.Text())
		point, _ := strconv.ParseFloat(input[1], 64)
		grade := input[2]
		if grade == "P" {
			continue
		}

		pointSum += point

		pointAndGradeSum += point * getPointByGrade(grade)
	}

	fmt.Printf("%.6f", pointAndGradeSum/pointSum)
}

func getPointByGrade(grade string) float64 {
	var point float64

	if grade == "F" {
		return 0.0
	} else if grade == "A0" {
		point = 4.0
	} else if grade == "B0" {
		point = 3.0
	} else if grade == "C0" {
		point = 2.0
	} else if grade == "D0" {
		point = 1.0
	} else if grade == "A+" {
		point = 4.5
	} else if grade == "B+" {
		point = 3.5
	} else if grade == "C+" {
		point = 2.5
	} else if grade == "D+" {
		point = 1.5
	}

	return point
}
