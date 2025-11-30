package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var s string
	s, _ = reader.ReadString('\n')
	s = strings.TrimSuffix(s, "\n")

	findList := []rune{'U', 'C', 'P', 'C'}
	ans := true

	reg := regexp.MustCompile("[a-z ]")

	newString := reg.ReplaceAllString(s, "")

	newRunes := []rune(newString)
	currentIndex := 0

	for _, f := range findList {
		foundIndex := -1

		for i := currentIndex; i < len(newRunes); i++ {
			if newRunes[i] == f {
				foundIndex = i
				break
			}
		}

		if foundIndex != -1 {
			currentIndex = foundIndex + 1
		} else {
			ans = false
			break
		}
	}

	if ans {
		fmt.Println("I love UCPC")
	} else {
		fmt.Println("I hate UCPC")
	}
}
