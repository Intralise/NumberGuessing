package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/fatih/color"
)

type GameConfig struct {
	minNumber, maxNumber, secretNumber, maxAttempts int
	difficulty                                      string
}

type GameResult struct {
	Date     string `json:"date"`
	Result   string `json:"result"`
	Attempts int    `json:"attempts"`
}

var config GameConfig
var guessedNumbers []int16

func attempt(counter int) bool {

	var number int16
	inputCheck(&number, counter, int16(config.minNumber), int16(config.maxNumber))
	guessedNumbers = append(guessedNumbers, number)
	return checkGuess(number)
}

func playGame() {
	var result string
	for i := 1; i <= 10; i++ {
		res := attempt(i)
		if res {
			color.Green("Вы угадали! \n Игра окончена")
			result = "Победа"
			i = 10
			break
		}
		fmt.Printf("Осталось попыток %d \n", 10-i)
		if i == 10 {
			color.Red("Вы проиграли! \n Секретное число было %d", config.secretNumber)
			result = "Проигрыш"
		}
	}
	fmt.Println("Ваши ответы:")
	fmt.Println(guessedNumbers)
	fmt.Println(result)
	if result == "Победа" {
		saveGameResult(createResult(result))
	}

}

func createResult(result string) []byte {
	rankingsJson, err := json.Marshal(GameResult{
		Date:     time.Now().Format(time.RFC3339),
		Result:   result,
		Attempts: len(guessedNumbers),
	})
	if err != nil {
		panic(err)
	}
	return rankingsJson
}

func setupGame() {

	fmt.Println("Для начала выбери сложность, есть три варианта: Hard, Normal, Easy 👺")
	for {
		res, err := readDifficulty()
		if err == nil {
			config = difficultyChange(res)
			break
		}
		fmt.Println("Неправильное значение. Допустимые значения: Easy, Normal, Hard")

	}

	fmt.Print("Игра ")
	color.New(color.FgGreen).Printf("'Угадай число'")
	fmt.Printf(" - от %d до - %d началась! \n", config.minNumber, config.maxNumber)
	fmt.Printf("Угадайте число за %d попыток!\n", config.maxAttempts)
	guessedNumbers = nil
}

func checkGuess(guess int16) bool {
	gap := int16(config.secretNumber) - guess
	if gap == 0 {
		return true
	}
	direction := "Секретное число больше 👆"
	if gap < 0 {
		direction = "Секретное число меньше 👇"
	}

	if gap < 0 {
		gap = -gap
	}

	switch {
	case gap <= 5:
		color.Yellow("Горячо 🔥")
	case gap <= 10:
		color.Yellow("Тепло ⭐")
	default:
		color.Yellow("Холодно ❄️")
	}

	fmt.Println(direction)
	return false
}

func difficultyChange(difficult string) GameConfig {
	switch difficult {
	case "Easy":
		return GameConfig{0, 50, rand.Intn(50), 15, "Easy"}
	case "Normal":
		return GameConfig{0, 100, rand.Intn(100), 10, "Normal"}
	case "Hard":
		return GameConfig{0, 200, rand.Intn(200), 5, "Hard"}
	default:
		return GameConfig{0, 50, rand.Intn(50), 15, "Easy"}
	}

}

func readDifficulty() (string, error) {
	var difficult string
	fmt.Fscan(os.Stdin, &difficult)
	switch difficult {
	case "Easy", "Normal", "Hard":
		return difficult, nil
	default:
		return "", fmt.Errorf("Ошибочное значение")
	}

}
