package slice

import "fmt"

func Run(){
	names := []string{"kim","lee"};
	// names[2] = "park"; 오류발생
	names = append(names,"park"); //append는 원본을변경하지않고 새로운배열을 반환한다.

	for idx,name := range names{
		fmt.Println(idx,name);
	}
}