package main

import (
	"errors"
	"fmt"
)

const pi = 3.1415

func main() {
	printCircleArea(2)
	printCircleArea(-2)
}

func printCircleArea(radius int) {
	area, err := calculateCircleArea(radius)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Радиус круга: %d см.\n", radius)
	fmt.Println("Формула для расчета площади круга: A=πr²")
	fmt.Printf("Площадь круга: %.2f см²\n\n", area)
}

func calculateCircleArea(radius int) (float32, error) {
	if radius <= 0 {
		return 0, errors.New("Радиус круга не может быть отрицательным")
	}
	return float32(radius) * float32(radius) * pi, nil
}

