package main

import "fmt"

const pi = 3.1415

func main() {
	printCircleArea(2)
	printCircleArea(-2)
}

func printCircleArea(radius int) {
	if radius <= 0 {
		fmt.Println("Радиус круга не может быть отрицательным!")
		return
	}

	fmt.Printf("Радиус круга: %d см.\n", radius)
	fmt.Println("Формула для расчета площади круга: A=πr²")

	circleArea := calculateCircleArea(radius)
	fmt.Printf("Площадь круга: %.2f см²\n\n", circleArea)
}

func calculateCircleArea(radius int) float32 {
	return float32(radius) * float32(radius) * pi
}
