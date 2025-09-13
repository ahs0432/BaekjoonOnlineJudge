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

	var n, k int
	fmt.Fscanln(reader, &n, &k)

	var s string
	fmt.Fscanln(reader, &s)

	sTable := strings.Split(s, ",")
	iTable := make([]int, len(sTable))
	for i := 0; i < len(sTable); i++ {
		iTable[i], _ = strconv.Atoi(sTable[i])
	}

	for i := 0; i < k; i++ {
		for j := 0; j < len(iTable)-1; j++ {
			iTable = append(iTable, iTable[1]-iTable[0])
			iTable = iTable[1:]
		}
		iTable = iTable[1:]
	}

	fmt.Fprint(writer, iTable[0])
	for i := 1; i < len(iTable); i++ {
		fmt.Fprintf(writer, ",%d", iTable[i])
	}
	fmt.Fprintln(writer)
	writer.Flush()
}
