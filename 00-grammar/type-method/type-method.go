package type_method_state

import "fmt"

type ar2x2 [2][2]int

func Run() {
	a := ar2x2{{1, 2}, {3, 4}}
	b := ar2x2{{5, 6}, {7, 8}}
	fmt.Println(a)
	fmt.Println(b)
	a.typeMethodVersion(b)
	fmt.Println(a)
}

func (a *ar2x2) typeMethodVersion(b ar2x2) {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			a[i][j] += b[i][j]
		}
	}
}
