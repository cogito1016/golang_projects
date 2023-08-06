package string_problem

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func RunBoj2744(){
	scanner := bufio.NewScanner(os.Stdin);
	scanner.Scan()
	str := scanner.Text();

	for _,char := range str{
		if unicode.IsUpper(char){
			fmt.Printf("%c", unicode.ToLower(char));
			continue;
		}

		fmt.Printf("%c",unicode.ToUpper(char));
	}
}