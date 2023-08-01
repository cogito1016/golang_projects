package for_state

import "fmt"

func Run(){
	// addNumbers01(1,2,3,4,5);
	// addNumbers02(1,2,3,4,5);
	// forNumbers03(1,2,3,4,5);
	fmt.Println(addNumbers(1,2,3,4,5));
}

func forNumbers01(numbers ...int) int{
	for number := range numbers{
		fmt.Println(number);
	}

	return 0;
}

func forNumbers02(numbers ...int) int{
	for idx,value := range numbers{
		fmt.Println(idx,value);
	}
	return 0;
}

func forNumbers03(numbers ...int) int{
	for i := 0; i< len(numbers); i++{
		fmt.Println(i,numbers[i]);
	}

	return 0;
}

func addNumbers(numbers ...int) int{
	total := 0;

	for _,number := range numbers{
		total += number;
	}
	
	return total;
}