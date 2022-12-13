package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)

	fmt.Fprintln(writer, "고려대학교")
	writer.Flush()
}
