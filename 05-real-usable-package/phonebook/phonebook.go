package main

import (
	"fmt"
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
		entry := Entry{name: "kim", surname: "sur" + strconv.Itoa(i), tel: "010-0000-1111"}
		data = append(data, entry)
	}

	fmt.Println(*search("sur9"))
	list()
}
