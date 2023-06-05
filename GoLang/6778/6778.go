package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscanln(reader, &n)
	fmt.Fscanln(reader, &m)

	if n >= 3 && m <= 4 {
		fmt.Println("TroyMartian")
	}

	if n <= 6 && m >= 2 {
		fmt.Println("VladSaturnian")
	}

	if n <= 2 && m <= 3 {
		fmt.Println("GraemeMercurian")
	}
}
