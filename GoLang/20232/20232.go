package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	winners := []string{"ITMO", "SPbSU", "SPbSU", "ITMO", "ITMO", "SPbSU", "ITMO", "ITMO", "ITMO", "ITMO", "ITMO", "PetrSU, ITMO", "SPbSU", "SPbSU", "ITMO", "ITMO", "ITMO", "ITMO", "SPbSU", "ITMO", "ITMO", "ITMO", "ITMO", "SPbSU", "ITMO"}
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)
	fmt.Println(winners[n-1995])
}
