package math_problem

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RunBoj11050(){
	scanner := bufio.NewScanner(os.Stdin);
	scanner.Scan();
	str := strings.Split(scanner.Text(), " ");
	N,_ := strconv.Atoi(str[0]);
	K,_ := strconv.Atoi(str[1]);

	cnk := factorial(N) / (factorial(K)*factorial(N-K));

	fmt.Println(cnk);
}

func factorial(num int) int{ //10팩은 약360만이므로
	if num==0 {
		return 1;
	}
	return num*factorial(num-1);
}