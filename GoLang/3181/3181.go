package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	checkWord := []string{"i", "pa", "te", "ni", "niti", "a", "ali", "nego", "no", "ili"}

	var s string
	s, _ = reader.ReadString('\n')
	s = strings.TrimSuffix(s, "\n")

	w := strings.Split(s, " ")

	ans := ""
	check := false
	for i := 0; i < len(w); i++ {
		if i != 0 {
			check = false
			for _, cw := range checkWord {
				if w[i] == cw {
					check = true
					break
				}
			}

			if check {
				continue
			}
		}

		ans += string(w[i][0] - 32)
	}
	fmt.Println(ans)
}
