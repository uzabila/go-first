// Делаем калькулятор индекса массы тела
package main

import (
	"fmt"
	"math"
)

const imtPow float64 = 2.0

var userHeight float64 // 0.0 по умолчанию
var userWeight float64

func main() {
	fmt.Println("______ Калькулятор индекса веса _______")
	userWeight, userHeight := getUserInput()
	bodyIndex := calculateIndex(userWeight, userHeight)

	outputIndex(bodyIndex)

	switch {
	case bodyIndex < 16:
		fmt.Println("Сильный дефицит массы тела")
	case bodyIndex < 18.5:
		fmt.Println("Дефицит массы тела")
	case bodyIndex < 25:
		fmt.Println("Норма")
	case bodyIndex < 30:
		fmt.Println("Избыточная масса")
	case bodyIndex < 35:
		fmt.Println("1-я степень ожирения")
	case bodyIndex < 40:
		fmt.Println("2-я степень ожирения")
	default:
		fmt.Println("3-я степень ожирения")
	}

	/* IF ELSE */
	/*
		if bodyIndex < 16 {
			fmt.Println("Сильный дефицит массы тела")
		} else if bodyIndex >= 16 && bodyIndex < 18.5 {
			fmt.Println("Дефицит массы тела")
		} else if bodyIndex >= 18.5 && bodyIndex < 25 {
			fmt.Println("Норма")
		} else if bodyIndex >= 25 && bodyIndex < 30 {
			fmt.Println("Избыточная масса")
		} else if bodyIndex >= 30 && bodyIndex < 35 {
			fmt.Println("1-я степень ожирения")
		} else if bodyIndex >= 35 && bodyIndex < 40 {
			fmt.Println("2-я степень ожирения")
		} else if bodyIndex > 40 {
			fmt.Println("3-я степень ожирения")
		}
	*/

}

func outputIndex(bodyIndex float64) {
	fmt.Printf("Ваш индекс массы тела: %.1f\n", bodyIndex)
}

func calculateIndex(userWeight, userHeight float64) float64 {
	var bodyIndex = userWeight / math.Pow(userHeight/100, imtPow)
	return bodyIndex
}

func getUserInput() (float64, float64) {

	fmt.Print("Какой ваш рост (см)? ")
	fmt.Scan(&userHeight)
	fmt.Print("Какой ваш вес (кг)? ")
	fmt.Scan(&userWeight)

	return userWeight, userHeight
}
