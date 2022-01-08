package main

import (
	"fmt"
)

func main() {

	fmt.Println("Расчёт сдачи.")

	bills := []int{5, 5, 5, 10, 20} // Вводим необходимые значения

	fmt.Println(lemonadeChange(bills))

}

func lemonadeChange(bills []int) bool {

	lemonadeChange := true

	var (
		coinFive   int
		coinTen    int
		coinTwenty int
	)

	for _, value := range bills {
		switch value {
		case 5:
			coinFive += 1
		case 10:
			coinTen += 1

			switch {
			case coinFive > 0:
				coinFive -= 1
			default:
				lemonadeChange = false
			}

		case 20:
			coinTwenty += 1

			switch {
			case coinFive > 0 && coinTwenty > 0:
				coinFive -= 1
				coinTen -= 1
			case coinFive >= 3:
				coinFive -= 3
			default:
				lemonadeChange = false
			}
		default:
			fmt.Println("У нас принимают монеты номиналом 10, 15 и 20")
		}
	}
	return lemonadeChange
}
