package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var w, p int
	fmt.Fscanln(reader, &w, &p)

	r := make([]bool, w+1)
	l := make([]int, 0, p+2)
	l = append(l, 0)
	l = append(l, w)

	var tmp, check int
	for i := 0; i < p; i++ {
		fmt.Fscan(reader, &tmp)

		for _, j := range l {
			check = int(math.Abs(float64(tmp - j)))
			r[check] = true
		}
		l = append(l, tmp)
	}

	for i := 1; i <= w; i++ {
		if r[i] {
			fmt.Fprint(writer, i, " ")
		}
	}
	fmt.Fprintln(writer, w)
	writer.Flush()
}
