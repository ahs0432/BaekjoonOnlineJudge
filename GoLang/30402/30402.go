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
	for i := 0; i < 15; i++ {
		s, _ = reader.ReadString('\n')

		if strings.Contains(s, "w") {
			fmt.Println("chunbae")
			break
		} else if strings.Contains(s, "b") {
			fmt.Println("nabi")
			break
		} else if strings.Contains(s, "g") {
			fmt.Println("yeongcheol")
			break
		}
	}
}
