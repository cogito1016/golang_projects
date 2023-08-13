package string_problem

import (
	"bufio"
	"fmt"
	"os"
)

func RunBoj4999() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	jaehwan := len(scanner.Text())
	scanner.Scan()
	doctor := len(scanner.Text())

	if jaehwan >= doctor {
		fmt.Println("go")
	} else {
		fmt.Println("no")
	}
}
