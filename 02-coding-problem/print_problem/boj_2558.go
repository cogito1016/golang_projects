package print_problem

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func RunBoj2558() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	A, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	B, _ := strconv.Atoi(scanner.Text())

	fmt.Println(A + B)
}
