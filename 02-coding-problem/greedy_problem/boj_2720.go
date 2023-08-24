package greedy_problem

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func RunBoj2720() {
	writer := bufio.NewWriter(os.Stdout)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	T, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < T; i++ {
		scanner.Scan()
		C, _ := strconv.Atoi(scanner.Text())

		quarter := C / 25
		C = C - quarter*25
		fmt.Fprintf(writer, strconv.Itoa(quarter)+" ")

		dime := C / 10
		C = C - dime*10
		fmt.Fprintf(writer, strconv.Itoa(dime)+" ")

		nickel := C / 5
		C = C - nickel*5
		fmt.Fprintf(writer, strconv.Itoa(nickel)+" ")

		penny := C
		C = C - penny
		fmt.Fprintf(writer, strconv.Itoa(penny)+" ")
		fmt.Fprintln(writer)
	}
	writer.Flush()
}
