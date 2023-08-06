package function_problem

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RunBoj15964(){
	scanner := bufio.NewScanner(os.Stdin);
	scanner.Scan();
	input := strings.Split(scanner.Text(), " ");
	A,_ := strconv.Atoi(input[0]);
	B,_ := strconv.Atoi(input[1]);

	fmt.Println(calc(A,B));
}

func calc(a,b int)int{
	return (a+b)*(a-b);
}