package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	s, _ := reader.ReadString('\n')
	s = strings.TrimSuffix(s, "\n")

	lefts := "qwertyasdfgzxcvb"
	rights := "uiophjklnm"
	cLeft, cRight, cOther := 0, 0, 0

	check := false
	for _, c := range s {
		if c >= 'A' && c <= 'Z' {
			cOther++
			c = c - 'A' + 'a'
		}

		check = false
		for _, right := range rights {
			if c == right {
				cRight++
				check = true
				break
			}
		}

		if !check {
			for _, left := range lefts {
				if c == left {
					cLeft++
					check = true
					break
				}
			}
		}

		if !check {
			cOther++
		}

	}

	for cOther != 0 {
		if cLeft <= cRight {
			cLeft++
		} else {
			cRight++
		}
		cOther--
	}

	fmt.Println(cLeft, cRight)
}
