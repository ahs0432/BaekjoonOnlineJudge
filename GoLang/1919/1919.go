package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a, b string
	fmt.Fscanln(reader, &a)
	fmt.Fscanln(reader, &b)

	wordsA := make([]int, 26)
	wordsB := make([]int, 26)

	for _, c := range a {
		wordsA[c-97]++
	}

	for _, c := range b {
		wordsB[c-97]++
	}

	ans := 0
	for i := 0; i < 26; i++ {
		if wordsA[i] != wordsB[i] {
			if wordsA[i]-wordsB[i] < 0 {
				ans += (wordsB[i] - wordsA[i])
			} else {
				ans += (wordsA[i] - wordsB[i])
			}
		}
	}
	fmt.Println(ans)
}
