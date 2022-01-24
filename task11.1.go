package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("Определение количества слов, начинающихся с большой буквы")

	a := "Go is an Open source programming Language that makes it Easy to build simple, reliable, and efficient Software."

	a = " " + a
	b := strings.Title(a)
	c := 0

	for len(a)-1 > 0 {
		if a[0:1] == " " && a[1:2] == b[1:2] {
			c++
		}
		b = b[1:]
		a = a[1:]
	}

	fmt.Println("Строка содержит", c, "слов с большой буквы")
}
