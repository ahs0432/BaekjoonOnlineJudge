package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var s, ts string
	fmt.Fscanln(reader, &s)
	fmt.Fscanln(reader, &ts)

	cs := ""
	for i := 0; i < len(s); i++ {
		if !(s[i] >= '0' && s[i] <= '9') {
			cs += string(s[i])
		}
	}

	if strings.Contains(cs, ts) {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}
