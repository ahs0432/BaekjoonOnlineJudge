package main

import (
	"bufio"
	"fmt"
	"os"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n, now int
	fmt.Fscanln(reader, &n)
	fmt.Fscanln(reader, &now)

	var tmp, sum int
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &tmp)
		sum += min(abs(now-tmp), 360-now+tmp, now+360-tmp)
		now = tmp
	}
	fmt.Println(sum)
}
