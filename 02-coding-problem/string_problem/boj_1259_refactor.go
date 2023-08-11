package string_problem

import (
	"bufio"
	"os"
)

func RunBoj1259Refactor(){
	scanner := bufio.NewScanner(os.Stdin);
	writer := bufio.NewWriter(os.Stdout);

	for {
		scanner.Scan();
		input := []byte (scanner.Text());
		if input[0]=='0' {
			break;
		}

		strLength := len(input)-1;
		result := "yes";
		for i := 0; i < len(input)/2; i++ {
			if(input[i] != input[strLength-i]){
				result = "no";
				break
			}
		}
		writer.WriteString(result+"\n")
	}
	writer.Flush();
}