package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	n += (2400 - 2013)
	fmt.Printf("%c%c\n", (n+5)%12+'A', (n+9)%10+'0')
}
