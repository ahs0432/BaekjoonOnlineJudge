package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	var s string
	fmt.Fscanln(reader, &s)

	var i64 int64
	var sum int64
	var tmp string
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			tmp += string(s[i])
		} else if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') {
			if len(tmp) != 0 {
				i64, _ = strconv.ParseInt(tmp, 10, 64)
				sum += i64
				tmp = ""
			}
		} else {
			continue
		}
	}

	if len(tmp) != 0 {
		i64, _ = strconv.ParseInt(tmp, 10, 64)
		sum += i64
	}
	fmt.Println(sum)
}
