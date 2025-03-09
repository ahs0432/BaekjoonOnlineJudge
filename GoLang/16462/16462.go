package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	ans := 0
	var tmp, now int
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &tmp)

		if tmp >= 100 {
			ans += 100
		} else {
			now = 1
			for tmp != 0 {
				if tmp%10 == 0 || tmp%10 == 6 {
					ans += 9 * now
				} else {
					ans += (tmp % 10) * now
				}
				tmp /= 10
				now *= 10
			}
		}
	}

	fmt.Println(math.Round(float64(ans) / float64(n)))
}
