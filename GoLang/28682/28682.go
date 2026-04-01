package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(reader, &n)

	lectures := []string{"swimming", "bowling", "soccer"}

	for i := 0; i < n; i++ {
		fmt.Fprint(writer, lectures[0], " ")
	}
	fmt.Fprintln(writer)
	writer.Flush()

	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	failLectures := strings.Split(line, " ")

	var next int
	for i := 0; i < n; i++ {
		fail := failLectures[i]
		if fail == "bowling" {
			next = 2
		} else {
			next = 1
		}
		fmt.Fprint(writer, lectures[next], " ")
	}
	fmt.Fprintln(writer)
	writer.Flush()
}
