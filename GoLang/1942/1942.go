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

	var str1, str2 string
	var tmp1, tmp2 []string
	var h1, m1, s1, h2, m2, s2 int
	var count, sum int
	for i := 0; i < 3; i++ {
		fmt.Fscanln(reader, &str1, &str2)
		tmp1 = strings.Split(str1, ":")
		tmp2 = strings.Split(str2, ":")

		h1, _ = strconv.Atoi(tmp1[0])
		m1, _ = strconv.Atoi(tmp1[1])
		s1, _ = strconv.Atoi(tmp1[2])

		h2, _ = strconv.Atoi(tmp2[0])
		m2, _ = strconv.Atoi(tmp2[1])
		s2, _ = strconv.Atoi(tmp2[2])

		count = 0

		for {
			sum = h1/10 + h1%10 + m1/10 + m1%10 + s1/10 + s1%10
			if sum%3 == 0 {
				count++
			}

			if h1 == h2 && m1 == m2 && s1 == s2 {
				break
			}

			s1++

			if s1 >= 60 {
				s1 = 0
				m1++
			}

			if m1 >= 60 {
				m1 = 0
				h1++
			}

			if h1 >= 24 {
				h1 = 0
			}
		}
		fmt.Fprintln(writer, count)
	}
	writer.Flush()
}
