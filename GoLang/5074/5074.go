package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var s1, s2 string

	for {
		fmt.Fscanln(reader, &s1, &s2)

		h1, _ := strconv.Atoi(strings.Split(s1, ":")[0])
		m1, _ := strconv.Atoi(strings.Split(s1, ":")[1])

		h2, _ := strconv.Atoi(strings.Split(s2, ":")[0])
		m2, _ := strconv.Atoi(strings.Split(s2, ":")[1])

		if h1 == m1 && m1 == h2 && h2 == m2 {
			break
		}

		t := h1*60 + m1 + h2*60 + m2
		n := t / 60 / 24
		h := t / 60 % 24
		m := t % 60

		if n > 0 {
			fmt.Fprintf(writer, "%02d:%02d +%d\n", h, m, n)
		} else {
			fmt.Fprintf(writer, "%02d:%02d\n", h, m)
		}
	}
	writer.Flush()
}
