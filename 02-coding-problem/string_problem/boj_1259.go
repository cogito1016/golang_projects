package string_problem

import (
	"bufio"
	"os"
	"strconv"
)

func RunBoj1259(){
	scanner := bufio.NewScanner(os.Stdin);
	writer := bufio.NewWriter(os.Stdout);

	input := getInterger(scanner);

	for input != 0 {
		str := strconv.Itoa(input);
		if checkPalindromes(str){
			writer.WriteString("yes")
		}else{
			writer.WriteString("no")
		}
		writer.WriteString("\n")
		input = getInterger(scanner);
	}

	writer.Flush();
}

func getInterger(scanner *bufio.Scanner)int{
	scanner.Scan();
	str := scanner.Text()
	result,_ := strconv.Atoi(str);
	return result;
}


func checkPalindromes(str string) bool{
	left := 0;
	right := len(str)-1;
	for left<=right{
		if str[left] != str[right]{
			return false;
		}
		left++;
		right--;
	}
	return true;
}
