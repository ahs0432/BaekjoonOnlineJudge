package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	ans := 0
	k := 0

	var tmp string
	for {
		fmt.Fscanln(reader, &tmp)

		if tmp == "=" {
			fmt.Println(ans)
			break
		} else if tmp == "+" {
			k = 0
		} else if tmp == "-" {
			k = 1
		} else if tmp == "*" {
			k = 2
		} else if tmp == "/" {
			k = 3
		} else {
			if k == 0 {
				a, _ := strconv.Atoi(tmp)
				ans += a
			} else if k == 1 {
				a, _ := strconv.Atoi(tmp)
				ans -= a
			} else if k == 2 {
				a, _ := strconv.Atoi(tmp)
				ans *= a
			} else if k == 3 {
				a, _ := strconv.Atoi(tmp)
				ans /= a
			}
		}
	}
}
