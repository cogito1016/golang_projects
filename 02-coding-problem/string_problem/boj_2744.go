package string_problem

import (
	"bufio"
	"fmt"
	"os"
)

func RunBoj2744(){
	scanner := bufio.NewScanner(os.Stdin);
	scanner.Scan()
	str := scanner.Text();

	for _,char := range str{
		if char>=97 {
			char -= 32;
		}else{
			char += 32;
		}

		fmt.Printf("%c",char);
	}
}