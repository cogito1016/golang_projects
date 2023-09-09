package type_assertion_switch

import "fmt"

func Run() {

}

type Secret struct {
	SecretValue string
}

type Entry struct {
	F1 int
	F2 string
	F3 Secret
}

func Teststruct(x interface{}) {
	switch T := x.(type) {
	case Secret:
		fmt.Println("Secret type")
	case Entry:
		fmt.Println("Entry type")
	default:
		fmt.Println("Not supported type : %T\n", T)
	}
}
