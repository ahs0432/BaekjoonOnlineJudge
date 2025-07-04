package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var s string
	s, _ = reader.ReadString('\n')

	if strings.Contains(s, "social") || strings.Contains(s, "history") || strings.Contains(s, "literacy") || strings.Contains(s, "language") {
		fmt.Println("digital humanities")
	} else if strings.Contains(s, "bigdata") || strings.Contains(s, "public") || strings.Contains(s, "society") {
		fmt.Println("public bigdata")
	}
}
