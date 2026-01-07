package game

import "fmt"

func printGallows(loses int, RuneWord []rune) {

	fmt.Println("Ты не угадал букву!")
	fmt.Println("Количество ошибок:", loses)

	switch {
	case loses == 1:
		{

			fmt.Println("  _______")
			fmt.Println(" |       |")
			fmt.Println(" |       O")
			fmt.Println(" |       ")
			fmt.Println(" |       ")
			fmt.Println(" |       ")
			fmt.Println("_|_")
		}

	case loses == 2:
		{

			fmt.Println("  _______")
			fmt.Println(" |       |")
			fmt.Println(" |       O")
			fmt.Println(" |       |")
			fmt.Println(" |       ")
			fmt.Println(" |       ")
			fmt.Println("_|_")
		}
	case loses == 3:
		{

			fmt.Println("  _______")
			fmt.Println(" |       |")
			fmt.Println(" |       O")
			fmt.Println(" |      /|")
			fmt.Println(" |       ")
			fmt.Println(" |       ")
			fmt.Println("_|_")
		}
	case loses == 4:
		{

			fmt.Println("  _______")
			fmt.Println(" |       |")
			fmt.Println(" |       O")
			fmt.Println(" |      /|\\")
			fmt.Println(" |       ")
			fmt.Println(" |       ")
			fmt.Println("_|_")
		}
	case loses == 5:
		{

			fmt.Println("  _______")
			fmt.Println(" |       |")
			fmt.Println(" |       O")
			fmt.Println(" |      /|\\")
			fmt.Println(" |      / ")
			fmt.Println(" |       ")
			fmt.Println("_|_")

		}
	case loses == 6:
		{

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
