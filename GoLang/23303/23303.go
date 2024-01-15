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
	if strings.Contains(s, "D2") || strings.Contains(s, "d2") {
		fmt.Println("D2")
	} else {
		fmt.Println("unrated")
	}
}
