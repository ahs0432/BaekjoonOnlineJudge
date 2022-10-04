package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s, err := reader.ReadString('\n')

	if err != nil {
		return
	}

	answer := 0
	first := true

	for _, c := range s {
		if string(c) == " " {
			if !first {
				first = true
			}
		} else if first {
			if c == 10 {
				continue
			}
			first = false
			answer++
		}
	}

	fmt.Println(answer)
}
