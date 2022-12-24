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
	fmt.Println(":fan::fan::fan:")
	fmt.Println(":fan::" + n + "::fan:")
	fmt.Println(":fan::fan::fan:")
}
