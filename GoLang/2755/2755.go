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

	total := 0.0
	nowTotal := 0.0
	for i := 0; i < n; i++ {
		var a, c string
		var b float64

		fmt.Fscanln(reader, &a, &b, &c)
		total += b

		if c == "A+" {
			nowTotal += (4.3 * b)
		} else if c == "A0" {
			nowTotal += (4.0 * b)
		} else if c == "A-" {
			nowTotal += (3.7 * b)
		} else if c == "B+" {
			nowTotal += (3.3 * b)
		} else if c == "B0" {
			nowTotal += (3.0 * b)
		} else if c == "B-" {
			nowTotal += (2.7 * b)
		} else if c == "C+" {
			nowTotal += (2.3 * b)
		} else if c == "C0" {
			nowTotal += (2.0 * b)
		} else if c == "C-" {
			nowTotal += (1.7 * b)
		} else if c == "D+" {
			nowTotal += (1.3 * b)
		} else if c == "D0" {
			nowTotal += (1.0 * b)
		} else if c == "D-" {
			nowTotal += (0.7 * b)
		}
	}

	if nowTotal == 0.0 {
		fmt.Println(0)
		return
	}
	fmt.Printf("%0.2f\n", math.Round((nowTotal/total)*100)/100)
}
