package game

import "math/rand"

func makeUnderlinedWord() ([]string, []rune) {
	words := []string{
		"мост", "кран", "окно", "леска", "кусты",
		"кость", "туча", "муха", "гроза", "ворот",
		"шляпа", "птица", "башня", "лодка", "ягода",
		"ветка", "печка", "крыша", "гвоздь", "ключи",
		"сумка", "ручка", "лампа", "кружка", "якорь",
		"тигр", "лиса", "щука", "котик", "пчела",
	}
	index := rand.Intn(len(words))

	runeWord := []rune(words[index])

	underlinedWord := []string

	for i, _ := range underlinedWord {
		underlinedWord[i] = "_"
	}

	return underlinedWord, runeWord
}
