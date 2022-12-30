package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	first := true
	s := ""
	for {
		tmp, _ := reader.ReadString('.')

		if !first {
			tmp = tmp[1:]
		} else {
			first = false
		}

		s += tmp

		if string(tmp[len(tmp)-1]) == "." {
			if len(s) == 1 {
				break
			}

			queue := make([]string, 0)
			check := true
			for _, r := range s {
				if string(r) == "[" {
					queue = append(queue, string(r))
				} else if string(r) == "(" {
					queue = append(queue, string(r))
				} else if string(r) == "]" {
					if len(queue) < 1 || queue[len(queue)-1] != "[" {
						fmt.Println("no")
						check = false
						break
					}

					queue = queue[:len(queue)-1]
				} else if string(r) == ")" {
					if len(queue) < 1 || queue[len(queue)-1] != "(" {
						fmt.Println("no")
						check = false
						break
					}

					queue = queue[:len(queue)-1]
				}
			}

			if check {
				if len(queue) != 0 {
					fmt.Println("no")
				} else {
					fmt.Println("yes")
				}
			}

			s = ""
		}
	}
}
