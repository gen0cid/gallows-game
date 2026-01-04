package casestart

import "fmt"

func printGallows(Loses int, RuneWord []rune) {
	switch {
	case Loses == 1:
		{
			fmt.Println("Ты не угадал букву!")
			fmt.Println("Количество ошибок:", Loses)

			fmt.Println("  _______")
			fmt.Println(" |       |")
			fmt.Println(" |       O")
			fmt.Println(" |       ")
			fmt.Println(" |       ")
			fmt.Println(" |       ")
			fmt.Println("_|_")
		}

	case Loses == 2:
		{
			fmt.Println("Ты не угадал букву!")
			fmt.Println("Количество ошибок:", Loses)

			fmt.Println("  _______")
			fmt.Println(" |       |")
			fmt.Println(" |       O")
			fmt.Println(" |       |")
			fmt.Println(" |       ")
			fmt.Println(" |       ")
			fmt.Println("_|_")
		}
	case Loses == 3:
		{
			fmt.Println("Ты не угадал букву!")
			fmt.Println("Количество ошибок:", Loses)

			fmt.Println("  _______")
			fmt.Println(" |       |")
			fmt.Println(" |       O")
			fmt.Println(" |      /|")
			fmt.Println(" |       ")
			fmt.Println(" |       ")
			fmt.Println("_|_")
		}
	case Loses == 4:
		{
			fmt.Println("Ты не угадал букву!")
			fmt.Println("Количество ошибок:", Loses)

			fmt.Println("  _______")
			fmt.Println(" |       |")
			fmt.Println(" |       O")
			fmt.Println(" |      /|\\")
			fmt.Println(" |       ")
			fmt.Println(" |       ")
			fmt.Println("_|_")
		}
	case Loses == 5:
		{
			fmt.Println("Ты не угадал букву!")
			fmt.Println("Количество ошибок:", Loses)

			fmt.Println("  _______")
			fmt.Println(" |       |")
			fmt.Println(" |       O")
			fmt.Println(" |      /|\\")
			fmt.Println(" |      / ")
			fmt.Println(" |       ")
			fmt.Println("_|_")

		}
	case Loses == 6:
		{
			fmt.Println("Ты не угадал букву!")
			fmt.Println("Количество ошибок:", Loses)

			fmt.Println("  _______")
			fmt.Println(" |       |")
			fmt.Println(" |       O")
			fmt.Println(" |      /|\\")
			fmt.Println(" |      / \\")
			fmt.Println(" |       ")
			fmt.Println("_|_")

			fmt.Println("Игра окончена!")
			fmt.Println("Загаданное слово было:", string(RuneWord))
			return
		}
	}
}
