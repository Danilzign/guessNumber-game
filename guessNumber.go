package main

import (
	"fmt"
)

func Guess() {
	min := 0
	max := 1000

	var comparison string

	for i := 0; i <= 10; i++ {
		hiddenNumber := (max + min) / 2
		fmt.Println("Твое число", hiddenNumber, "?")
		fmt.Scan(&comparison)

		switch comparison {
		case "<":
			max = hiddenNumber
		case ">":
			min = hiddenNumber
		case "=":
			fmt.Println("Твое число:", hiddenNumber)
			i = 10
		default:
			fmt.Println("Неправильный ввод")
			i = 10
		}

		if hiddenNumber > 1000 || hiddenNumber < 0 {
			break
		}

	}
}
