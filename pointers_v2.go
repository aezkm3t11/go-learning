package main

import "fmt"

func main() {
	x := 10
	fmt.Println("Значение X", x)

	increment(&x)

	fmt.Println("Значение X после вызова функции increment", x)
}

func increment(p *int) {
	*p += 1
}
