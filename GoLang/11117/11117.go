package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var t int
	fmt.Fscanln(reader, &t)

	var alpha, alphaCopy []int
	var cookieNum int
	var cookieWord, nowWord string
	var check bool
	for i := 0; i < t; i++ {
		alpha = make([]int, 26)
		fmt.Fscanln(reader, &cookieWord)
		for i := 0; i < len(cookieWord); i++ {
			alpha[cookieWord[i]-'A']++
		}

		fmt.Fscanln(reader, &cookieNum)
		for j := 0; j < cookieNum; j++ {
			alphaCopy = make([]int, 26)
			copy(alphaCopy, alpha)

			fmt.Fscanln(reader, &nowWord)

			check = true
			for k := 0; k < len(nowWord); k++ {
				if alphaCopy[nowWord[k]-'A'] == 0 {
					check = false
					break
				} else {
					alphaCopy[nowWord[k]-'A']--
				}
			}

			if check {
				fmt.Fprintln(writer, "YES")
			} else {
				fmt.Fprintln(writer, "NO")
			}
		}
	}
	writer.Flush()
}
