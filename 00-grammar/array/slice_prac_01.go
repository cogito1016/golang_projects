package array_statement

import "fmt"

func SlicePrac01() {
	test_one := []int{1, 2, 3, 4}
	test_two := []int{5, 6, 7, 8}
	fmt.Println(attached(test_two, test_one))
}

// 두개배열연결해 새로운배열반환
func attached(slice_one, slice_two []int) []int {
	return append(slice_one, slice_two...)
}
