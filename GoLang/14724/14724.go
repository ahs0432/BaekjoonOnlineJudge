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

	club := []string{"PROBRAIN", "GROW", "ARGOS", "ADMIN", "ANT", "MOTION", "SPG", "COMON", "ALMIGHTY"}

	var tmp, ans, max int
	for i := 0; i < 9; i++ {
		for j := 0; j < n; j++ {
			fmt.Fscan(reader, &tmp)
			if max < tmp {
				max = tmp
				ans = i
			}
		}
	}

	fmt.Println(club[ans])
}
