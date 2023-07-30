package main

import (
	"fmt"

	"github.com/cogito1016/golang/01-module/greetings"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Gladys");
    fmt.Println(message)
}