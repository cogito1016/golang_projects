package math_problem

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func RunBoj1978(){
	scanner := bufio.NewScanner(os.Stdin);
	scanner.Scan()
	N,_ := strconv.Atoi(scanner.Text());

	if N == 0 {
		return ;
	}

	scanner.Scan()
	input := strings.Split(scanner.Text(), " ");

	result := 0;
	for i := 0; i < N; i++ {
		target,_ := strconv.Atoi(input[i]); 

		if(isPrime(target)){
			result++;
		}
	}

	fmt.Println(result);
}

func isPrime(num int) bool{
	if num == 1 {
		return false;
	}

	max := int(math.Sqrt(float64(num)));

	for i := 2; i <= max; i++ {
		if(num%i==0){
			return false;
		}
	}

	return true;
}