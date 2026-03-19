package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(reader, &n)

	var r, a int
	var s string
	var ans float64
	for i := 0; i < n; i++ {
		r, a = 0, 0
		for {
			s, _ = reader.ReadString('\n')
			s = strings.TrimSuffix(s, "\n")
			if s == "" {
				break
			}

			for _, c := range s {
				if c == '#' {
					a++
				} else {
					a++
					r++
				}
			}
		}

		ans = math.Round((float64(r)/float64(a)*100)*10) / 10
		if ans == math.Trunc(ans) {
			fmt.Fprintf(writer, "Efficiency ratio is %d%%.\n", int(ans))
		} else {
			fmt.Fprintf(writer, "Efficiency ratio is %.1f%%.\n", ans)
		}
	}
	writer.Flush()
}
