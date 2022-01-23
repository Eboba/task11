package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println("Вывод чисел, содержащихся в строке")

	a := "a10 10 20b 20 30c30 30 dd 33"
	a = a + " "
	n := 0
	m := 1

	var answer string

	for len(a)+1 != m {
		_, err := strconv.ParseInt(a[n:m-1], 10, 64)

		if a[m-1:m] == " " {
			if err == nil {
				answer += a[n:m]
			}
			n = m
		}
		m++

	}
	fmt.Println("В строке содержатся числа в десятичном формате:")
	fmt.Println(answer)

}
