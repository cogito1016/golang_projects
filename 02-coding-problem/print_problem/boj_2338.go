package print_problem

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func RunBoj2338() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	aStr := scanner.Text()
	scanner.Scan()
	bStr := scanner.Text()

	a := new(big.Int)
	b := new(big.Int)

	a.SetString(aStr, 10)
	b.SetString(bStr, 10)

	// Calculate A + B
	sum := new(big.Int)
	sum.Add(a, b)
	fmt.Println(sum.String())

	// Calculate A - B
	diff := new(big.Int)
	diff.Sub(a, b)
	fmt.Println(diff.String())

	// Calculate A * B
	product := new(big.Int)
	product.Mul(a, b)
	fmt.Println(product.String())
}
