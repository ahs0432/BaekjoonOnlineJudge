package main

import (
	"bufio"
	"fmt"
	"os"
)

type qType struct {
	a   int
	b   int
	now string
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(reader, &n)

	for i := 0; i < n; i++ {
		var a, b int
		fmt.Fscanln(reader, &a, &b)

		table := make([]bool, 10000)

		queue := make([]qType, 0)
		queue = append(queue, qType{a, b, ""})
		for len(queue) != 0 {
			now := queue[0]
			queue = queue[1:]

			if now.a == now.b {
				fmt.Fprintln(writer, now.now)
				break
			}

			tmp := (now.a * 2) % 10000
			if !table[tmp] {
				queue = append(queue, qType{tmp, now.b, now.now + "D"})
				table[tmp] = true
			}

			if now.a == 0 {
				tmp = 9999
			} else {
				tmp = now.a - 1
			}

			if !table[tmp] {
				queue = append(queue, qType{tmp, now.b, now.now + "S"})
				table[tmp] = true
			}

			tmp = (now.a*10)%10000 + (now.a*10)/10000
			if !table[tmp] {
				queue = append(queue, qType{tmp, now.b, now.now + "L"})
				table[tmp] = true
			}

			tmp = now.a/10 + (now.a%10)*1000
			if !table[tmp] {
				queue = append(queue, qType{tmp, now.b, now.now + "R"})
				table[tmp] = true
			}
		}
	}

	writer.Flush()
}
