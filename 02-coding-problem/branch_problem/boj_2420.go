package branch_problem

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func RunBoj2420(){
	scanner := bufio.NewScanner(os.Stdin);
	scanner.Scan();
	input := strings.Split(scanner.Text(), " ");
	N,_ := strconv.Atoi(input[0]);
	M,_ := strconv.Atoi(input[1]);
	fmt.Println(int(math.Abs(float64(N-M))));
}