package game

import (
	"bufio"
	"os"
	"strings"
)

func reader() string {

	var i string
	reader := bufio.NewReader(os.Stdin)

	i, _ = reader.ReadString('\n')
	i = strings.TrimSpace(i)
	return i
}
