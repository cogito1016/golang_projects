package math_problem

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RunBoj2609(){
	scanner := bufio.NewScanner(os.Stdin);
	scanner.Scan()
	inputArr := strings.Split(scanner.Text(), " ");
	A,_ := strconv.Atoi(inputArr[0]);
	B,_ := strconv.Atoi(inputArr[1]);

	maximumCommitment := gcd(A,B); 
	minumumAirDrain := (A*B)/maximumCommitment;

	fmt.Println(maximumCommitment)
	fmt.Println(minumumAirDrain)
}

func gcd(a int, b int) int{
	if a<b {
		c := a
		a = b;
		b = c;
	}

	if b==0{
		return a;
	}

	result := a%b;
	return gcd(b,result);
}