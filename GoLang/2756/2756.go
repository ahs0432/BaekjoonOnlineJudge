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

	var n int
	fmt.Fscanln(reader, &n)

	var score1, score2 int
	var x, y, tmp float64
	for i := 0; i < n; i++ {
		score1, score2 = 0, 0
		for j := 0; j < 3; j++ {
			fmt.Fscan(reader, &x, &y)
			tmp = math.Hypot(x, y)

			if tmp <= 3 {
				score1 += 100
			} else if tmp <= 6 {
				score1 += 80
			} else if tmp <= 9 {
				score1 += 60
			} else if tmp <= 12 {
				score1 += 40
			} else if tmp <= 15 {
				score1 += 20
			}
		}

		for j := 0; j < 3; j++ {
			fmt.Fscan(reader, &x, &y)
			tmp = math.Hypot(x, y)

			if tmp <= 3 {
				score2 += 100
			} else if tmp <= 6 {
				score2 += 80
			} else if tmp <= 9 {
				score2 += 60
			} else if tmp <= 12 {
				score2 += 40
			} else if tmp <= 15 {
				score2 += 20
			}
		}

		if score1 > score2 {
			fmt.Fprintf(writer, "SCORE: %d to %d, PLAYER 1 WINS.\n", score1, score2)
		} else if score1 < score2 {
			fmt.Fprintf(writer, "SCORE: %d to %d, PLAYER 2 WINS.\n", score1, score2)
		} else {
			fmt.Fprintf(writer, "SCORE: %d to %d, TIE.\n", score1, score2)
		}
	}
	writer.Flush()
}
