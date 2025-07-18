package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var s string
	s, _ = reader.ReadString('\n')
	s = strings.TrimSuffix(s, "\n")

	targetStrings := strings.Split(s, " ")

	var checkString string
	fmt.Fscanln(reader, &checkString)
	count := 0

	for _, targetString := range targetStrings {
		if len(targetString) <= len(checkString) {
			continue
		}

		if targetString[0:len(checkString)] == checkString {
			count++
		}
	}
	fmt.Println(count)
}
