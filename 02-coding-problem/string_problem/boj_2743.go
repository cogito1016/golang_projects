package string_problem

import (
	"bufio"
	"fmt"
	"os"
)

func RunBoj2743(){
	scanner := bufio.NewScanner(os.Stdin);
	scanner.Scan()
	str := scanner.Text();
	fmt.Println(len(str));
}