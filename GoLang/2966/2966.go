package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	a := []rune{'A', 'B', 'C'}
	b := []rune{'B', 'A', 'B', 'C'}
	c := []rune{'C', 'C', 'A', 'A', 'B', 'B'}

	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	var s string
	fmt.Fscanln(reader, &s)

	max := 0
	sa, sb, sc := 0, 0, 0

	for i, r := range s {
		if a[i%len(a)] == r {
			sa++
			if max < sa {
				max = sa
			}
		}

		if b[i%len(b)] == r {
			sb++
			if max < sb {
				max = sb
			}
		}

		if c[i%len(c)] == r {
			sc++
			if max < sc {
				max = sc
			}
		}
	}

	fmt.Println(max)
	if sa == max {
		fmt.Println("Adrian")
	}
	if sb == max {
		fmt.Println("Bruno")
	}
	if sc == max {
		fmt.Println("Goran")
	}
}
