package service

import (
	"fmt"
)

func Hello(greeting string) string {
	say(greeting)
	return "hello world"
}

func say(p string) {
	fmt.Println(p)
}
