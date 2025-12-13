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

	ans := 0
	var h, m, d, time int
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%2d:%2d %d\n", &h, &m, &d)
		time = h*60 + m

		for j := 0; j < d; j++ {
			if time >= 420 && time < 1140 {
				ans += 10
			} else {
				ans += 5
			}
			time++

			if time == 1440 {
				time = 0
			}
		}
	}
	fmt.Println(ans)
}
