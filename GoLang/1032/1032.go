package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	var main, tmp string
	fmt.Fscanln(reader, &main)

	mainTmp := []byte(main)

	for i := 1; i < n; i++ {
		fmt.Fscanln(reader, &tmp)
		for j := 0; j < len(main); j++ {
			if mainTmp[j] != byte('?') && mainTmp[j] != byte(tmp[j]) {
				mainTmp[j] = '?'
			}
		}
	}

	fmt.Println(string(mainTmp))
}
