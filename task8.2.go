package main

import (
	"fmt"
)

func main() {

	fmt.Println("Дни недели.")

	fmt.Println("Введите будний день недели: пн, вт, ср, чт, пт:")
	var day string
	_, _ = fmt.Scan(&day)

	switch day {
	case "пн":
		fmt.Println("Понедельник")
		fallthrough
	case "вт":
		fmt.Println("Вторник")
		fallthrough
	case "ср":
		fmt.Println("Среда")
		fallthrough
	case "чт":
		fmt.Println("Четверг")
		fallthrough
	case "пт":
		fmt.Println("Пятница")
	default:
		fmt.Println("Неверно указан день недели.")
	}
}
