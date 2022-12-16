package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscanln(reader, &n)
	fmt.Fscanln(reader, &m)

	button := make([]bool, 10)
	for i := 0; i < m; i++ {
		var tmp int
		if i == m-1 {
			fmt.Fscanln(reader, &tmp)
		} else {
			fmt.Fscan(reader, &tmp)
		}

		button[tmp] = true
	}

	min := 0

	if n < 100 {
		min = 100 - n
	} else {
		min = n - 100
	}

	if min != 0 {
		for i := 0; i <= n*2+9; i++ {
			str := strconv.Itoa(i)
			now := len(str)

			if n < i {
				now += i - n
			} else {
				now += n - i
			}

			if now < min {
				check := true
				for _, c := range str {
					if button[int(c)-48] {
						check = false
						break
					}
				}

				if check {
					min = now
				}
			}
		}
	}

	fmt.Println(min)
}
