package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	var tmp []rune
	if n < 2023 {
		fmt.Println(0)
	} else {
		ans := 0

		for i := 2023; i < n+1; i++ {
			s := fmt.Sprintf("%d", i)

			tmp = make([]rune, 0)
			for _, c := range s {
				if c == '2' && len(tmp) == 0 {
					tmp = append(tmp, c)
				} else if c == '0' && len(tmp) == 1 {
					tmp = append(tmp, c)
				} else if c == '2' && len(tmp) == 2 {
					tmp = append(tmp, c)
				} else if c == '3' && len(tmp) == 3 {
					ans++
					break
				}
			}
		}
		fmt.Println(ans)
	}
}
