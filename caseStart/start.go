package casestart

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Start() {

	UnderlinedWord, RuneWord := underlinedWord()

	for i, _ := range UnderlinedWord {
		UnderlinedWord[i] = "_"
	}

	var input string
	Loses := 0

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Загаданное слово:", UnderlinedWord)
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
		for i := 0; i < len(RuneWord); i++ {

			if []rune(input)[0] == RuneWord[i] {
				UnderlinedWord[i] = input
			}
		}

		// проверка есть ли буква игрока в загаданном слове. если нету то счетчик ошибок увеличивается
		if check([]rune(input)[0], RuneWord) == false {
			Loses++
		}

		// висилицы для разного количества ошибок
		printGallows(Loses, RuneWord)

		isOpen := true
		// проверяю если слово разгадано
		for i, _ := range UnderlinedWord {
			if UnderlinedWord[i] == "_" {
				isOpen = false
				break
			}
		}
		if isOpen {
			fmt.Println("🎉 Поздравляем! Слово угадано!")
			fmt.Println(UnderlinedWord)
			return
		}

	}
}
