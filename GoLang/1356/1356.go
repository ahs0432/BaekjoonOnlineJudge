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

	check := false
	var tmpL, tmpR int = 1, 1
	for i := 0; i < len(n)-1; i++ {
		tmpR = 1
		tmpL *= int(n[i] - '0')

		for k := i + 1; k < len(n); k++ {
			tmpR *= int(n[k] - '0')
		}

		if tmpL == tmpR {
			check = true
			break
		}
	}

	if check {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
