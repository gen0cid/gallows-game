package game

import (
	"fmt"
)

func Start() {

	try := try{
		loses: 0,
		imageGallow: [][]string{
			{" ", " ", " ", " ", "+", "-", "-", "-", "-", "-", "+", " "},
			{" ", " ", " ", " ", "|", " ", " ", " ", " ", " ", "|", " "},
			{" ", " ", " ", " ", "|", " ", " ", " ", " ", " ", "|", " "},
			{" ", " ", " ", " ", "|", " ", " ", " ", " ", " ", "|", " "},
			{" ", " ", " ", " ", "|", " ", " ", " ", " ", " ", " ", " "},
			{" ", " ", " ", " ", "|", " ", " ", " ", " ", " ", " ", " "},
			{" ", " ", " ", " ", "|", " ", " ", " ", " ", " ", " ", " "},
			{" ", " ", " ", " ", "|", " ", " ", " ", " ", " ", " ", " "},
			{" ", " ", " ", " ", "|", " ", " ", " ", " ", " ", " ", " "},
			{" ", " ", " ", " ", "|", " ", " ", " ", " ", " ", " ", " "},
			{" ", " ", " ", " ", "|", " ", " ", " ", " ", " ", " ", " "},
			{"-", "-", "-", "-", "+", "-", "-", "-", "-", "-", " ", " "},
		},
	}

	var input string
	var runeAskWord []rune
	letters := make(map[string]bool)

	// записываю в поле underlinedWord структуры try массив из нижних подчеркиваний длинной загаданного слова
	// записываю в runeArrWord массив из рун заггаданного слова
	try.underlinedWord, runeAskWord = makeUnderlinedWord()

	for {

		for {
			// делаю валидацию слова пользователя
			for {

				//читаю слово пользователя и записываю его в переменную input
				fmt.Print("Введите букву:")
				input = reader()

				if validation([]rune(input)) {
					break
				} else {
					fmt.Println("Вы ввели не одну букву")
					fmt.Println("Повторите попытку!")
				}
			}
			_, ok := letters[input]

			if !ok {

				// добавляю в мапу и слайс букву пользователя
				letters[input] = true
				try.usedLetters = append(try.usedLetters, input)

				// Обновляем маску
				for i := 0; i < len(runeAskWord); i++ {

					if []rune(input)[0] == runeAskWord[i] {
						try.underlinedWord[i] = input
					}

				}
			} else {
				fmt.Println("Вы уже вводили эту букву, попробуйте другую")
				continue
			}
			break
		}

		// проверка есть ли буква игрока в загаданном слове.
		// если нету то счетчик ошибок увеличивается
		if check([]rune(input)[0], runeAskWord) == false {
			try.loses++
		}

		// меняю изображение висилицы для разного количества поражений
		try.printGallows(try.loses)

		//если loses = 6 - игра окончена
		if try.loses == 6 {
			fmt.Println("Игра окончена!")
			break
		}

		// проверяю если слово разгадано
		isUnderlined := true
		for i, _ := range try.underlinedWord {
			if try.underlinedWord[i] == "_" {
				isUnderlined = false
			}
		}
		if isUnderlined {
			fmt.Println("🎉 Поздравляем! Слово угадано!")
			fmt.Println(try.underlinedWord)
			break
		}
		try.print()
	}

}
