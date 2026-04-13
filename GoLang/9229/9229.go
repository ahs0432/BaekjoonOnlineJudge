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

	var a, ns string
	for {
		line, _ := reader.ReadString('\n')
		words := strings.Fields(line)
		if len(words) == 0 {
			continue
		}

		s := words[0]
		if s == "#" {
			break
		}

		isCorrect := true

		for {
			if len(words) > 1 {
				a = words[1]
				words = words[1:]
			} else {
				ns, _ = reader.ReadString('\n')
				words = strings.Fields(ns)
				if len(words) == 0 {
					continue
				}
				a = words[0]
			}

			if a == "#" {
				break
			}

			if isCorrect {
				if len(s) != len(a) {
					isCorrect = false
				} else {
					cnt := 0
					for i := 0; i < len(a); i++ {
						if a[i] != s[i] {
							cnt++
						}
					}
					if cnt != 1 {
						isCorrect = false
					}
				}
			}
			s = a
		}

		if isCorrect {
			fmt.Fprintln(writer, "Correct")
		} else {
			fmt.Fprintln(writer, "Incorrect")
		}
	}
	writer.Flush()
}
