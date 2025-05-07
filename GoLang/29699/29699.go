package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	s := "WelcomeToSMUPC"
	var n int
	fmt.Fscanln(reader, &n)
	fmt.Println(string(s[(n-1)%len(s)]))
}
