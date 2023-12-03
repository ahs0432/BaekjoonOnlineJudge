package main

import (
	"fmt"
)

func main() {
	for a := 6; a < 101; a++ {
		for b := 2; b < a; b++ {
			for c := b + 1; c < a; c++ {
				for d := c + 1; d < a; d++ {
					if a*a*a == (b*b*b + c*c*c + d*d*d) {
						fmt.Printf("Cube = %d, Triple = (%d,%d,%d)\n", a, b, c, d)
					} else if a*a*a < (b*b*b + c*c*c + d*d*d) {
						break
					}
				}
			}
		}
	}
}
