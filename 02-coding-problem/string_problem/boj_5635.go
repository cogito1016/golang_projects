package string_problem

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RunBoj5635() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	firstPersonRaw := strings.Fields(scanner.Text())

	firstPerson := getPerson(firstPersonRaw)
	oldPerson := firstPerson
	youngPerson := firstPerson

	for i := 1; i < N; i++ {
		scanner.Scan()
		targetPerson := getPerson(strings.Fields(scanner.Text()))

		oldPerson = getOldPerson(oldPerson, targetPerson)
		youngPerson = getYoungPerson(youngPerson, targetPerson)
	}

	fmt.Println(youngPerson.name)
	fmt.Println(oldPerson.name)
}

func getOldPerson(oldPerson person, targetPerson person) person {
	if oldPerson.year >= targetPerson.year {
		if oldPerson.year == targetPerson.year {
			if oldPerson.month >= targetPerson.month {
				if oldPerson.month == targetPerson.month {
					if oldPerson.day >= targetPerson.day {
						if oldPerson.day > targetPerson.day {
							return targetPerson
						} else {
							return oldPerson
						}
					}
				} else {
					return targetPerson
				}
			}
		} else {
			return targetPerson
		}
	}

	return oldPerson
}

func getYoungPerson(youngPerson person, targetPerson person) person {
	if youngPerson.year <= targetPerson.year {
		if youngPerson.year == targetPerson.year {
			if youngPerson.month <= targetPerson.month {
				if youngPerson.month == targetPerson.month {
					if youngPerson.day <= targetPerson.day {
						if youngPerson.day < targetPerson.day {
							return targetPerson
						} else {
							return youngPerson
						}
					}
				} else {
					return targetPerson
				}
			}
		} else {
			return targetPerson
		}
	}

	return youngPerson
}

func getPerson(perRaw []string) person {
	day, _ := strconv.Atoi(perRaw[1])
	month, _ := strconv.Atoi(perRaw[2])
	year, _ := strconv.Atoi(perRaw[3])
	return person{name: perRaw[0], day: day, month: month, year: year}
}

type person struct {
	name  string
	day   int
	month int
	year  int
}
