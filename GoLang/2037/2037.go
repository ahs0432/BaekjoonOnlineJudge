package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	keyboard := [][]byte{{}, {'A', 'B', 'C'}, {'D', 'E', 'F'}, {'G', 'H', 'I'}, {'J', 'K', 'L'}, {'M', 'N', 'O'},
		{'P', 'Q', 'R', 'S'}, {'T', 'U', 'V'}, {'W', 'X', 'Y', 'Z'}}

	var p, w int
	fmt.Fscanln(reader, &p, &w)
	s, _ := reader.ReadString('\n')
	s = strings.TrimSuffix(s, "\n")

	total := 0
	past := 0

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			total += p
			past = 0
		}

		for j := 0; j < len(keyboard); j++ {
			for k := 0; k < len(keyboard[j]); k++ {
				if s[i] == keyboard[j][k] {
					if j == past {
						total += w
					}
					total += (1 + k) * p
					past = j
				}
			}
		}
	}
	fmt.Println(total)
}
