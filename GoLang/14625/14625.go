package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var h1, h2, m1, m2 int
	var n int
	fmt.Fscanln(reader, &h1, &m1)
	fmt.Fscanln(reader, &h2, &m2)
	fmt.Fscanln(reader, &n)

	ans := 0
	str := ""

	for h := h1; h <= h2; h++ {
		for m := m1; m <= 59; m++ {
			str = fmt.Sprintf("%02d%02d", h, m)
			if strings.Contains(str, string(n+48)) {
				ans++
			}
			if h == h2 && m == m2 {
				break
			}
		}
		m1 = 0
	}
	fmt.Println(ans)
}
