package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var c int
	fmt.Fscanln(reader, &c)

	max, now := 0, 0
	var tmp string
	for i := 0; i < c; i++ {
		fmt.Fscanln(reader, &tmp)

		now = 0
		for j := 0; j < len(tmp); j++ {
			if len(tmp)-j-1 >= 4 && tmp[j] == 'w' && tmp[j+1] == 'h' && tmp[j+2] == 'i' && tmp[j+3] == 'l' && tmp[j+4] == 'e' {
				now++
			}

			if len(tmp)-j-1 >= 2 && tmp[j] == 'f' && tmp[j+1] == 'o' && tmp[j+2] == 'r' {
				now++
			}
		}

		if now > max {
			max = now
		}
	}
	fmt.Println(max)
}
