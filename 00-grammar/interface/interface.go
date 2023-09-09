package _interface_state

import (
	"fmt"
	"sort"
)

func Run() {
	data := []S2{
		S2{1, "One", S1{1, "S1_1", 10}},
		S2{2, "Two", S1{2, "S1_2", 20}},
		S2{-1, "Three", S1{-1, "S1_3", -20}},
	}

	fmt.Println(data)

	sort.Sort(S2slice(data))

	fmt.Println(data)
}

type S1 struct {
	F1 int
	F2 string
	F3 int
}

type S2 struct {
	F1 int
	F2 string
	F3 S1
}

type S2slice []S2

func (a S2slice) Len() int {
	return len(a)
}

func (a S2slice) Less(i, j int) bool {
	return a[i].F3.F1 < a[j].F3.F1
}

func (a S2slice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
