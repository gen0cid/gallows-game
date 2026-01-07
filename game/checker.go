package game

func check(l rune, word []rune) bool {
	for _, v := range word {
		if l == v {
			return true
		}
	}
	return false
}
