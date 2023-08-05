package print_problem

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func RunBoj10869(){
	scanner := bufio.NewScanner(os.Stdin);
	writer := bufio.NewWriter(os.Stdout);
	scanner.Scan();
	input := strings.Split(scanner.Text(), " ");
	A,_ := strconv.Atoi(input[0]);
	B,_ := strconv.Atoi(input[1]);

	writer.WriteString(strconv.Itoa(A+B)+"\n");
	writer.WriteString(strconv.Itoa(A-B)+"\n");
	writer.WriteString(strconv.Itoa(A*B)+"\n");
	writer.WriteString(strconv.Itoa(A/B)+"\n");
	writer.WriteString(strconv.Itoa(A%B));

	writer.Flush();
}