package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Entry struct {
	name    string
	surname string
	tel     string
}

var data []Entry

func search(surname string) *Entry {
	for key, value := range data {
		if value.surname == surname {
			return &data[key]
		}
	}
	return nil
}

func list() {
	for _, v := range data {
		fmt.Println(v)
	}
}

func main() {
	for i := 0; i < 10; i++ { //dummy
		entry := Entry{name: "kim", surname: "sur" + strconv.Itoa(i), tel: "010-" + setRandomNumber(4) + "-" + setRandomNumber(4)}
		data = append(data, entry)
	}

	fmt.Println(*search("sur9"))
	list()
}

func setDummyData() {

}

func setRandomNumber(len int) string { //자릿수를받아 랜덤수를 반환
	var str string
	for i := 0; i < len; i++ {
		str += strconv.Itoa(rand.Intn(9))
	}

	return str
}

//func setRandomName() string {
//
//}
