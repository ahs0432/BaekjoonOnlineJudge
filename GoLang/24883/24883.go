package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n string
	fmt.Fscanln(reader, &n)

	if n == "n" || n == "N" {
		fmt.Println("Naver D2")
	} else {
		fmt.Print("Naver Whale")
	}
}
