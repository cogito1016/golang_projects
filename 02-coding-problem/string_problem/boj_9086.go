package string_problem

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func RunBoj9086(){
	scanner := bufio.NewScanner(os.Stdin);
	scanner.Scan()
	N,_ := strconv.Atoi(scanner.Text());

	for i := 0; i < N; i++ {
		scanner.Scan();
		str := scanner.Text();
		fmt.Printf("%c%c\n",str[0],str[len(str)-1]);
	}
}