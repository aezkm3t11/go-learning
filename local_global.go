package main

import "fmt"

var globalVar ="Я глобальная переменнная!"
func main () {
  localVar := "а я локальная переменнная, объявлена внутри функции main()"

  fmt.Println (globalVar)
  fmt.Println (localVar)
  
}
//локальная и глобальная переменные
