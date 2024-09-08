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

	var t1, t2 string
	fmt.Fscanln(reader, &t1)
	fmt.Fscanln(reader, &t2)

	t1s := strings.Split(t1, ":")
	t2s := strings.Split(t2, ":")

	h1, _ := strconv.Atoi(t1s[0])
	m1, _ := strconv.Atoi(t1s[1])
	s1, _ := strconv.Atoi(t1s[2])

	h2, _ := strconv.Atoi(t2s[0])
	m2, _ := strconv.Atoi(t2s[1])
	s2, _ := strconv.Atoi(t2s[2])

	t1i := h1*3600 + m1*60 + s1
	t2i := h2*3600 + m2*60 + s2
	ans := t2i - t1i

	if ans < 0 {
		ans += 60 * 60 * 24
	}

	fmt.Printf("%02d:%02d:%02d\n", (ans / 3600), (ans%3600)/60, (ans % 60))
}
