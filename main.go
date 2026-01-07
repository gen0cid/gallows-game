package main

import (
	"bufio"
	"fmt"
	"gallows-game/game"
	"os"
)

func main() {

	for {
		fmt.Println("Привет, это висилица!")
		fmt.Println("Start - начать игру")
		fmt.Println("Exit - выйти из игры")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		text := scanner.Text()
		fmt.Println(text)

		switch {
		case text == "Start":
			game.Start()
		case text == "Exit":
			fmt.Println("До свидания!")
			return
		default:
			fmt.Println("Ты ввел некорректную команду")
		}

	}
}
