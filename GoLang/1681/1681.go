package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, l int
	fmt.Fscanln(reader, &n, &l)

	ans := 0
	var tmp int
	var check bool
	for count := 0; count != n; {
		ans++
		tmp = ans
		check = false
		for tmp != 0 {
			if tmp%10 == l {
				check = true
				break
			}
			tmp /= 10
		}

		if check {
			continue
		}

		count++
	}
	fmt.Println(ans)
}
