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

	a := []rune(n)
	sum := 0
	for i := 0; i < len(a); i++ {
		if a[i] == 'a' || a[i] == 'e' || a[i] == 'i' || a[i] == 'o' || a[i] == 'u' {
			sum++
		}
	}
	fmt.Println(sum)
}
