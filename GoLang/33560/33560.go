package main

import (
	"bufio"
	"fmt"
	"os"
)

func finish(scores []int, ans []int) {
	if scores[2] >= 35 && scores[2] < 65 {
		ans[0]++
	} else if scores[2] >= 65 && scores[2] < 95 {
		ans[1]++
	} else if scores[2] >= 95 && scores[2] < 125 {
		ans[2]++
	} else if scores[2] >= 125 {
		ans[3]++
	}

	scores[0] = 1
	scores[1] = 4
	scores[2] = 0
	scores[3] = 0
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	ans := []int{0, 0, 0, 0}
	scores := []int{1, 4, 0, 0}

	var tmp int
	for i := 0; i < n; i++ {
		if scores[3] > 240 {
			finish(scores, ans)
			i--
			continue
		}

		fmt.Fscan(reader, &tmp)
		if tmp == 1 {
			finish(scores, ans)
			continue
		} else if tmp == 2 {
			if scores[0] > 1 {
				scores[0] /= 2
			} else {
				scores[1] += 2
			}
		} else if tmp == 4 {
			scores[3] += 56
		} else if tmp == 5 {
			if scores[1] > 1 {
				scores[1]--
			}
		} else if tmp == 6 {
			if scores[0] < 32 {
				scores[0] *= 2
			}
		}

		scores[2] += scores[0]
		scores[3] += scores[1]
	}

	for i := 0; i < 4; i++ {
		fmt.Println(ans[i])
	}
}
