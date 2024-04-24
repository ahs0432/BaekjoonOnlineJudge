package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(reader, &n)

	res := 0.0
	var tmp float32
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &tmp)
		a := 1.0 - float64(res/100)
		b := 1.0 - float64(tmp/100)
		res = float64(1-(a*b)) * 100.0
		fmt.Fprintln(writer, res)
	}
	writer.Flush()
}
