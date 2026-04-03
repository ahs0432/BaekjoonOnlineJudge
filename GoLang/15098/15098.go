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
	s = strings.TrimSuffix(s, "\n")

	words := strings.Split(s, " ")

	check := true
	checkTable := map[string]bool{}
	for _, word := range words {
		if checkTable[word] {
			check = false
		} else {
			checkTable[word] = true
		}
	}

	if check {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
