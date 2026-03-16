package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	days := []string{"Mon", "Tue", "Wed", "Thu", "Fri"}

	dayMap := make(map[string]int)
	for i, day := range days {
		dayMap[day] = i
	}

	var t, n int
	fmt.Fscanln(reader, &t, &n)

	var d1, d2 string
	var t1, t2 int
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &d1, &t1, &d2, &t2)

		if d1 == d2 {
			t -= (t2 - t1)
		} else {
			t -= 24*(dayMap[d2]-dayMap[d1]) + (t2 - t1)
		}
	}

	if t > 48 {
		fmt.Println(-1)
	} else if t <= 0 {
		fmt.Println(0)
	} else {
		fmt.Println(t)
	}
}
