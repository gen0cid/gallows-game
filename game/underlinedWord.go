package game

import "math/rand"

func makeUnderlinedWord() ([]string, []rune) {
	var index int

	words := []string{
		"мост", "кран", "окно", "леска", "кусты",
		"кость", "туча", "муха", "гроза", "ворот",
		"шляпа", "птица", "башня", "лодка", "ягода",
		"ветка", "печка", "крыша", "гвоздь", "ключи",
		"сумка", "ручка", "лампа", "кружка", "якорь",
		"тигр", "лиса", "щука", "котик", "пчела",
	}

	index = rand.Intn(len(words))
	underlinedWord := []string{}

	for i := 0; i < len([]rune(words[index])); i++ {
		underlinedWord = append(underlinedWord, "_")
	}

	runeAskWord := []rune(words[index])

	return underlinedWord, runeAskWord
}
