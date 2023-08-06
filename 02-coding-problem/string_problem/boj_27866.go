package string_problem

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func RunBoj27866(){
	scanner := bufio.NewScanner(os.Stdin);
	scanner.Scan();
	str := scanner.Text();
	scanner.Scan();
	idx,_ := strconv.Atoi(scanner.Text());
	idx--;

	fmt.Printf("%c",str[idx]);
}