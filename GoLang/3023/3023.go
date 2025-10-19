package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var r, c int
	fmt.Fscanln(reader, &r, &c)

	cards := make([][]byte, r)
	for i := 0; i < r; i++ {
		fmt.Fscanln(reader, &cards[i])
	}

	var a, b int
	fmt.Fscanln(reader, &a, &b)

	for i := 0; i < r; i++ {
		for j := c - 1; j >= 0; j-- {
			cards[i] = append(cards[i], cards[i][j])
		}
	}

	for i := r - 1; i >= 0; i-- {
		cards = append(cards, make([]byte, len(cards[i])))
		copy(cards[len(cards)-1], cards[i])
	}

	if cards[a-1][b-1] == '#' {
		cards[a-1][b-1] = '.'
	} else {
		cards[a-1][b-1] = '#'
	}

	for _, card := range cards {
		fmt.Fprintln(writer, string(card))
	}
	writer.Flush()
}
