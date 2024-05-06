package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	t := make([][]string, 10)

	for i := 0; i < 10; i++ {
		t[i] = make([]string, 10)

		for j := 0; j < 10; j++ {
			fmt.Fscan(reader, &t[i][j])
		}
	}

	for i := 0; i < 10; i++ {
		able := true
		for j := 1; j < 10; j++ {
			if t[i][j] != t[i][0] {
				able = false
			}
		}

		if able {
			fmt.Println(1)
			return
		}
	}

	for i := 0; i < 10; i++ {
		able := true
		for j := 1; j < 10; j++ {
			if t[j][i] != t[0][i] {
				able = false
			}
		}

		if able {
			fmt.Println(1)
			return
		}
	}

	fmt.Println(0)
}
