package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("Определение количества слов, начинающихся с большой буквы")

	a := "Go is an Open source programming Language that makes it Easy to build simple, reliable, and efficient Software."

	b := strings.ToLower(a)
	c := 0

	for len(a) > 0 {
		if a[0:1] != b[0:1] {
			c++
		}
		b = b[1:]
		a = a[1:]
	}

	fmt.Println("Строка содержит", c, "слов с большой буквы")
}
