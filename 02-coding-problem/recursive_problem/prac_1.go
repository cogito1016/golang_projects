package recursive_problem

import "fmt"

func RunPrac01() {
	fmt.Println(multiple_recursive(5))
	fmt.Println(multiple_for(5))
}

func multiple_recursive(n int) int {
	if n == 1 {
		return 1
	}

	return n * multiple_recursive(n-1)
}

func multiple_for(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}

	return result
}
