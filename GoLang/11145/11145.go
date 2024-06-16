package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var t int
	fmt.Fscanln(reader, &t)
	r, _ := regexp.Compile("^[\\s]*([\\d]+)[\\s]*$")

	for i := 0; i < t; i++ {
		tmp, _ := reader.ReadString('\n')
		tmp = strings.TrimSpace(tmp)

		if check := r.MatchString(tmp); check {
			a := false
			for i := 0; i < len(tmp); i++ {
				if a {
					if tmp[i] >= '0' && tmp[i] <= '9' {
						fmt.Fprint(writer, string(tmp[i]))
					}
				} else if tmp[i] != '0' && tmp[i] != ' ' {
					fmt.Fprint(writer, string(tmp[i]))
					a = true
				}
			}

			if !a {
				fmt.Fprint(writer, 0)
			}
			fmt.Fprintln(writer)
		} else {
			fmt.Fprintln(writer, "invalid input")
		}
	}
	writer.Flush()
}
