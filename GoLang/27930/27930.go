package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var s string
	fmt.Fscanln(reader, &s)

	korea := "KOREA"
	yonsei := "YONSEI"

	koreaScore := 0
	yonseiScore := 0

	for i := 0; i < len(s); i++ {
		if s[i] == korea[koreaScore] {
			koreaScore++
		}
		if s[i] == yonsei[yonseiScore] {
			yonseiScore++
		}

		if koreaScore == len(korea) {
			fmt.Println(korea)
			break
		}
		if yonseiScore == len(yonsei) {
			fmt.Println(yonsei)
			break
		}
	}
}
