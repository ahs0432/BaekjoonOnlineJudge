package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	p := 0
	n := 0
	for i := 0; i < 4; i++ {
		var tmp int
		fmt.Fscanln(reader, &tmp)

		if i == 0 {
			p = tmp
			continue
		}

		if n == 0 {
			if p < tmp {
				n = 1
			} else if p == tmp {
				n = 2
			} else {
				n = 3
			}
		} else {
			if n == 4 {
				continue
			} else if (n == 1 && !(p < tmp)) || (n == 2 && !(p == tmp)) || (n == 3 && !(p > tmp)) {
				n = 4
			}
		}
		p = tmp
	}

	if n == 1 {
		fmt.Println("Fish Rising")
	} else if n == 2 {
		fmt.Println("Fish At Constant Depth")
	} else if n == 3 {
		fmt.Println("Fish Diving")
	} else {
		fmt.Println("No Fish")
	}
}
