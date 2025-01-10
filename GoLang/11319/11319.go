package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(reader, &n)

	var s string
	var check bool
	var vowels, consonants int
	checkVowels := []rune{'A', 'E', 'I', 'O', 'U', 'a', 'e', 'i', 'o', 'u'}

	for i := 0; i < n; i++ {
		s, _ = reader.ReadString('\n')
		s = strings.TrimSuffix(s, "\n")

		vowels = 0
		consonants = 0
		for _, c := range s {
			if c == ' ' {
				continue
			}

			check = false
			for _, checkVowel := range checkVowels {
				if c == checkVowel {
					check = true
					vowels++
					break
				}
			}

			if !check {
				consonants++
			}
		}

		fmt.Fprintln(writer, consonants, vowels)
	}
	writer.Flush()
}
