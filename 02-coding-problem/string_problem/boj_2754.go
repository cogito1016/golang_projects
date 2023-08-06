package string_problem

import (
	"bufio"
	"fmt"
	"os"
)

func RunBoj2754(){
	scanner := bufio.NewScanner(os.Stdin);
	scanner.Scan()
	grade := scanner.Text();

	eng := grade[0];
	var point float32

	if eng=='F'{
		fmt.Println("0.0");
		return;
	}else if eng=='A'{
		point = 4.0;
	}else if eng=='B'{
		point = 3.0;
	}else if eng=='C'{
		point = 2.0;
	}else{
		point = 1.0;
	}

	weight := grade[1];

	if weight=='+'{
		point += 0.3;
	}else if weight=='-'{
		point -= 0.3;
	}

	fmt.Printf("%.1f",point);
}