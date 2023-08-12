package deque_problem

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func RunBoj10866(){
	scanner := bufio.NewScanner(os.Stdin);
	writer := bufio.NewWriter(os.Stdout);

	defer writer.Flush()

	scanner.Scan()
	N,_ := strconv.Atoi(scanner.Text());

	var arr []int;
	var result []int;

	for N>0 {
		input := getInstructor();
		arr,result = actionDeque(arr,result,input);
		N--
	}

	for i := 0; i < len(result); i++ {
		writer.WriteString(strconv.Itoa(result[i])+"\n");
	}
}

func getInstructor() []string{
	scanner := bufio.NewScanner(os.Stdin);
	scanner.Scan();
	return strings.Fields(scanner.Text());
}

func actionDeque(arr []int, result []int,instructArr []string) ([]int,[]int) {
	instructor := instructArr[0];

	switch instructor{
	case "push_back":
		number,_ := strconv.Atoi(instructArr[1]);
		arr = pushBack(arr,number);
	case "push_front":
		number,_ := strconv.Atoi(instructArr[1]);
		arr = pushFront(arr,number);
	case "front":
		result = append(result,front(arr));
	case "back":
		result = append(result,back(arr));
	case "pop_front":
		value := 0
		value,arr = popFront(arr);
		result = append(result,value);
	case "pop_back":
		value := 0
		value,arr = popBack(arr);
		result = append(result,value);
	case "size":
		result = append(result,size(arr));
	case "empty":
		value := 0;
		if empty(arr){
			value = 1;
		}
		result = append(result, value);
	}

	return arr,result
}

func empty(arr []int)bool{
	return len(arr)==0;
}

func pushBack(arr []int, number int) []int{
	return append(arr,number);
}

func pushFront(arr []int, number int) []int{
	var newArr []int
	newArr = append(newArr, number);
	newArr = append(newArr, arr...); 
	return newArr
}

func front(arr []int)int{
	if(empty(arr)){
		return -1;
	}

	return arr[0];
}

func back(arr []int)int{
	if(empty(arr)){
		return -1;
	}

	return arr[len(arr)-1];
}

func size(arr []int)int{
	return len(arr);
}

func popFront(arr []int) (int,[]int){
	result := front(arr);
	if result == -1{
		return -1,arr;
	}
	var newArr []int;
	newArr = append(newArr, arr[1:]...);
	return result,newArr
}

func popBack(arr []int) (int,[]int){
	result := back(arr);
	if result == -1{
		return -1,arr;
	}
	var newArr []int;
	newArr = append(newArr, arr[0:len(arr)-1]...);
	return result,newArr
}