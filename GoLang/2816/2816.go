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

	var tmp string
	kbs1, kbs2 := 0, 0
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &tmp)
		if tmp == "KBS1" {
			kbs1 = i
		} else if tmp == "KBS2" {
			kbs2 = i
		}
	}

	if kbs1 > kbs2 {
		kbs2++
	}

	for i := 0; i < kbs1; i++ {
		fmt.Fprint(writer, "1")
	}

	for i := 0; i < kbs1; i++ {
		fmt.Fprint(writer, "4")
	}

	if kbs2 != 1 {
		for i := 0; i < kbs2; i++ {
			fmt.Fprint(writer, "1")
		}

		for i := 0; i < kbs2-1; i++ {
			fmt.Fprint(writer, "4")
		}
	}
	fmt.Fprintln(writer)
	writer.Flush()
}
