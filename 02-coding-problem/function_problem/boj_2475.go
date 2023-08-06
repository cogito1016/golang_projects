package function_problem

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func RunBoj2475(){
	scanner := bufio.NewScanner(os.Stdin);
	scanner.Scan()
	input := strings.Split(scanner.Text(), " ");

	result := 0;
	for _,val := range input {
		num,_ := strconv.Atoi(val);
		result += int(math.Pow( float64(num) ,2));
	}

	fmt.Println(result%10);
}