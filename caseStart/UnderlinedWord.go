package casestart

import "math/rand"

func underlinedWord() ([]string, []rune) {
	words := []string{
		"мост", "кран", "окно", "леска", "кусты",
		"кость", "туча", "муха", "гроза", "ворот",
		"шляпа", "птица", "башня", "лодка", "ягода",
		"ветка", "печка", "крыша", "гвоздь", "ключи",
		"сумка", "ручка", "лампа", "кружка", "якорь",
		"тигр", "лиса", "щука", "котик", "пчела",
	}
	index := rand.Intn(len(words))

	RuneWord := []rune(words[index])

	return make([]string, len(RuneWord)), RuneWord
}
