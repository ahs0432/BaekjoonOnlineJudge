package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	all := ""
loop:
	for {
		s, err := reader.ReadString('\n')
		s = strings.TrimSuffix(s, "\n")
		switch err {
		case nil:
		case io.EOF:
			break loop
		}
		all += s
	}

	allS := strings.Split(all, ",")
	var sum, tmp int

	for _, v := range allS {
		tmp, _ = strconv.Atoi(v)
		sum += tmp
	}
	fmt.Println(sum)
}
