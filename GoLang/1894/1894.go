package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Point struct {
	x, y float64
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var p1, p2, p3, p4 Point

	for {
		_, err := fmt.Fscan(reader, &p1.x, &p1.y, &p2.x, &p2.y, &p3.x, &p3.y, &p4.x, &p4.y)
		if err == io.EOF {
			break
		}

		if p2 == p3 {
			p1, p2 = p2, p1
		} else if p1 == p4 {
			p3, p4 = p4, p3
		} else if p2 == p4 {
			p1, p2 = p2, p1
			p3, p4 = p4, p3
		}

		resX := p1.x + (p2.x - p1.x) + (p4.x - p1.x)
		resY := p1.y + (p2.y - p1.y) + (p4.y - p1.y)

		fmt.Fprintf(writer, "%.3f %.3f\n", resX, resY)
	}
	writer.Flush()
}
