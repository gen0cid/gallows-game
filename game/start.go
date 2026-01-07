package game

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Start() {

	underlinedWord, runeWord := makeUnderlinedWord()

	var input string
	loses := 0

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Загаданное слово:", underlinedWord)
		for {
			fmt.Print("Введите букву:")
			input, _ = reader.ReadString('\n')
			input = strings.TrimSpace(input)

			runes := []rune(input)

			if len(runes) == 1 {
				break
			} else {
				fmt.Println("Вы ввели не одну букву")
				fmt.Println("Повторите попытку!")
			}
		}

		// Обновляем маску
		for i := 0; i < len(runeWord); i++ {

			if []rune(input)[0] == runeWord[i] {
				underlinedWord[i] = input
			}
		}

		// проверка есть ли буква игрока в загаданном слове. если нету то счетчик ошибок увеличивается
		if check([]rune(input)[0], runeWord) == false {
			loses++
		}

		// висилицы для разного количества ошибок
		printGallows(loses, runeWord)

		isOpen := true
		// проверяю если слово разгадано
		for i, _ := range underlinedWord {
			if underlinedWord[i] == "_" {
				isOpen = false
				break
			}
		}
		if isOpen {
			fmt.Println("🎉 Поздравляем! Слово угадано!")
			fmt.Println(underlinedWord)
			return
		}

	}
}
