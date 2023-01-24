package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)

	fmt.Fprintln(writer, "       _.-;;-._\n'-..-'|   ||   |\n'-..-'|_.-;;-._|\n'-..-'|   ||   |\n'-..-'|_.-''-._|")
	writer.Flush()
}
