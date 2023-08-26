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
		entry := Entry{name: setRandomName(3, 3), surname: setRandomName(4, 8), tel: "010-" + setRandomNumber(4) + "-" + setRandomNumber(4)}
		data = append(data, entry)
	}

	fmt.Println(search("sur9"))
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

func setRandomName(minLen, maxLen int) string { //소문자이름 랜덤생성
	var nameLen int

	if minLen == maxLen {
		nameLen = minLen
	} else {
		nameLen = rand.Intn(maxLen-minLen) + minLen
	}

	name := ""
	//아스키코드 a~z 97~122
	for i := 0; i < nameLen; i++ {
		name += string(rand.Intn(122-97) + 97)
	}

	return name
}
