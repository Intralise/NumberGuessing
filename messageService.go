package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func init() {
	fmt.Print("Привет! Ты попал в игру '")
	color.New(color.FgGreen).Printf("%s", "Угадай число")
	fmt.Println("' приветствуем нашего нового игрока! ✨")
}

func inputCheck(number *int16, counter int, minNumber int16, maxNumber int16) {
	fmt.Printf("Попытка #%d - ", counter)
	for {
		var s string

		fmt.Print("Введите целое число: ")
		fmt.Scan(&s)

		var sign int16 = 1
		start := 0

		if s[0] == '-' {
			sign = -1
			start = 1
		}

		var n int16 = 0
		ok := true

		for i := start; i < len(s); i++ {
			if s[i] < '0' || s[i] > '9' {
				ok = false
				break
			}
			n = n*10 + int16(s[i]-'0')
		}

		if !ok {
			fmt.Println("Неверный ввод. Нужно целое число.")
			continue
		}
		if n < minNumber && n > maxNumber {
			fmt.Printf("Неверный ввод. Нужно число в диапазоне от %d до %d.\n", minNumber, maxNumber)
			continue
		}

		*number = n * sign
		return
	}
}

func askToPlayAgain() bool {
	fmt.Println("Хотите сыграть ещё раз? (Да / Нет)")
	var answer string
	fmt.Fscan(os.Stdin, &answer)
	return answer == "Да"
}
