package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	var s string
	fmt.Fscanln(reader, &n, &s)

	count := 0
	uniqueChars := ""
	seen := make(map[rune]bool)

	for _, c := range s {
		if !seen[c] {
			seen[c] = true
			uniqueChars += string(c)
		} else {
			count++
		}
	}

	merge := strconv.Itoa(n+1906) + uniqueChars + strconv.Itoa(count+4)

	mergeRunes := []rune(merge)
	for i, j := 0, len(mergeRunes)-1; i < j; i, j = i+1, j-1 {
		mergeRunes[i], mergeRunes[j] = mergeRunes[j], mergeRunes[i]
	}
	fmt.Println("smupc_" + string(mergeRunes))
}
