package recursive_problem

import "fmt"

// 배열 요소의 합
func RunPrac02() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(forVersion(arr))
	fmt.Println(recursiveVersion(arr))
}

func forVersion(arr []int) int {
	result := 0
	for i := 0; i < len(arr); i++ {
		result += arr[i]
	}

	return result
}

func recursiveVersion(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}

	return arr[len(arr)-1] + recursiveVersion(arr[:len(arr)-1])
}
