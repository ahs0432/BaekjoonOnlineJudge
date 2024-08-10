package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	colors := map[string]int{
		"black":  0,
		"brown":  1,
		"red":    2,
		"orange": 3,
		"yellow": 4,
		"green":  5,
		"blue":   6,
		"violet": 7,
		"grey":   8,
		"white":  9,
	}

	color := make([]string, 3)

	for i := 0; i < 3; i++ {
		fmt.Fscanln(reader, &color[i])
	}

	fmt.Println((colors[color[0]]*10 + colors[color[1]]) * int(math.Pow(10.0, float64(colors[color[2]]))))
}
