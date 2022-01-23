package main

import (
	"fmt"
)

func main() {

	fmt.Println("Определение количества слов, начинающихся с большой буквы")

	a := "Go is an Open source programming Language that makes it Easy to build simple, reliable, and efficient Software."

	var b int

	for len(a) > 0 {

		switch a[:1] {
		case "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z":
			b++
		}
		a = a[1:]
	}

	fmt.Println("Строка содержит", b, "слов с большой буквы")
}
