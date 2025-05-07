package main

import "fmt"

// Глобальная переменная
var globalVar = "Я глобальная переменная!"

func main() {
  // Локальная переменная
  localVar := "а я локальная переменная, объявлена внутри функции main()"

  fmt.Println(globalVar)
  fmt.Println(localVar)
}
