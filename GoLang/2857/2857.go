package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	ans := ""
	for i := 0; i < 5; i++ {
		var n string
		fmt.Fscanln(reader, &n)

		for j := 0; j < len(n)-2; j++ {
			if n[j:j+3] == "FBI" {
				ans += string(byte((i+1)+48)) + " "
				break
			}
		}
	}

	if ans == "" {
		fmt.Println("HE GOT AWAY!")
	} else {
		fmt.Println(ans)
	}
}
