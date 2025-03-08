package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n string
	fmt.Fscanln(reader, &n)

	var now int64
	var ans int64
	for i := 0; i < len(n); i++ {
		now = 1
		for j := 0; j < len(n); j++ {
			ans = ans + (int64(n[i]-'0') * now)
			now *= 10
		}
	}
	fmt.Println(ans)
}
