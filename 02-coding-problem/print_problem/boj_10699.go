package print_problem

import (
	"fmt"
	"time"
)

func RunBoj10699(){
	loc,_ := time.LoadLocation("Asia/Seoul");

	t := time.Now();
	t = t.In(loc);
	fmt.Println(t.Format("2006-01-02"));
}