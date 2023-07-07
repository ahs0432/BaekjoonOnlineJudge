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

	check := make([]bool, 5)
	for i := 0; i < len(n); i++ {
		if n[i] == 'M' {
			check[0] = true
		} else if n[i] == 'O' {
			check[1] = true
		} else if n[i] == 'B' {
			check[2] = true
		} else if n[i] == 'I' {
			check[3] = true
		} else if n[i] == 'S' {
			check[4] = true
		}
	}

	if check[0] && check[1] && check[2] && check[3] && check[4] {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
