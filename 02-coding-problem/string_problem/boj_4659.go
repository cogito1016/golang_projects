package string_problem

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func RunBoj4659() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		str := scanner.Text()
		if str == "end" {
			break
		}

		necessaryChecker := false
		doubleChecker := false
		tripleChecker := false
		tripleMoChecker := 0
		tripleJaChecker := 0
		var continueStr string

		for i := 0; i < len(str); i++ {
			char := string(str[i])

			if !necessaryChecker {
				necessaryChecker = isVowels(char)
			}

			if isVowels(char) {
				if tripleJaChecker != 0 {
					tripleJaChecker = 0
				}
				tripleMoChecker++
			} else {
				if tripleMoChecker != 0 {
					tripleMoChecker = 0
				}
				tripleJaChecker++
			}

			if tripleMoChecker == 3 || tripleJaChecker == 3 {
				tripleChecker = true
				break
			}

			if len(continueStr) == 0 {
				continueStr += char
			} else if len(continueStr) == 1 {
				if (continueStr == "e" && char == "e") || (continueStr == "o" && char == "o") {
					continueStr = ""
				} else if continueStr == char {
					doubleChecker = true
					break
				} else {
					continueStr = char
				}
			}
		}

		if !necessaryChecker || tripleChecker || doubleChecker {
			fmt.Println("<" + str + ">" + " is not acceptable.")
			continue
		}

		fmt.Println("<" + str + ">" + " is acceptable.")
	}

}

func isVowels(char string) bool { //모음체크
	necessaryAlphabet := "aeiou"
	if strings.Index(necessaryAlphabet, char) != -1 { // *5
		return true
	}
	return false
}
