package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	var s string
	fmt.Fscanln(reader, &n)
	fmt.Fscanln(reader, &s)

	alpha := make([]int, 26)
	for i := 0; i < len(s); i++ {
		alpha[s[i]-'A']++
	}
	alpha['R'-'A'] /= 2
	alpha['E'-'A'] /= 2

	fmt.Println(min(alpha['B'-'A'], alpha['R'-'A'], alpha['O'-'A'], alpha['N'-'A'], alpha['Z'-'A'], alpha['E'-'A'], alpha['S'-'A'], alpha['I'-'A'], alpha['L'-'A'], alpha['V'-'A']))
}
