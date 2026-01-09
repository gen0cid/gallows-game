package game

import (
	"fmt"
	"strings"
)

func (t *try) print() {
	fmt.Println("Загаданное слово:", t.underlinedWord)
	fmt.Println("Количество поражений:", t.loses)
	fmt.Println("Использованные буквы:", t.usedLetters)

	for _, v := range t.imageGallow {
		line := strings.Join(v, "")
		fmt.Println(line)
	}

}
