package math_problem

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func RunBoj4153Refactor(){
	writer := bufio.NewWriter(os.Stdout);
	scanner := bufio.NewScanner(os.Stdin);

	for {
		scanner.Scan();
		str := strings.Fields(scanner.Text());

		if str[0]=="0"{
			break;
		}

		a,_ := strconv.Atoi(str[0]);
		b,_ := strconv.Atoi(str[1]);
		c,_ := strconv.Atoi(str[2]);

		if(a>b){
			temp:=a;
			a=b;
			b=temp;
		}

		if(c>b){
			temp := c;
			c = b;
			b = temp;
		}

		result := "wrong";
		if b*b == (a*a+c*c){
			result = "right"
		}
		
		writer.WriteString(result+"\n");
	}

	writer.Flush();

}