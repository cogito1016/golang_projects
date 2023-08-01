package if_state

import "fmt"

func Run(){
	fmt.Println(canDrink(18));
	fmt.Println(canDrink2(18));
}

func canDrink(age int) bool{
	if age > 18{
		return true;
	}
	
	return false;
}

func canDrink2(age int) bool{
	if koreanAge := age+2; koreanAge>18{
		return true;
	}

	return false;
}