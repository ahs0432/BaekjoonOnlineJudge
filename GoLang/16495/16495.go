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

	var ans int64
	for i := 0; i < len(s); i++ {
		ans *= 26
		ans += int64(s[i] - 'A' + 1)
	}
	fmt.Println(ans)
}
