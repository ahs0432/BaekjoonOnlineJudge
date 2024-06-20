package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	s, _ := reader.ReadString('\n')
	s = strings.Trim(s, "\n")

	total := 0
	for _, c := range s {
		total += int(c)
	}

	fmt.Println(string(rune(total / len(s))))
}
