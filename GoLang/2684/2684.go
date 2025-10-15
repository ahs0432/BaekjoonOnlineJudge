package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(reader, &n)

	var s string
	var coins map[string]int
	for i := 0; i < n; i++ {
		coins = map[string]int{"TTT": 0, "TTH": 0, "THT": 0, "THH": 0, "HTT": 0, "HTH": 0, "HHT": 0, "HHH": 0}
		fmt.Fscanln(reader, &s)

		for j := 0; j < len(s)-2; j++ {
			coins[s[j:j+3]]++
		}

		fmt.Fprintln(writer, coins["TTT"], coins["TTH"], coins["THT"], coins["THH"], coins["HTT"], coins["HTH"], coins["HHT"], coins["HHH"])
	}
	writer.Flush()
}
