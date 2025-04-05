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
	fmt.Fscanln(reader, &s)

	lines := strings.Split(s, "|")

	minors := []byte{'A', 'D', 'E'}
	majors := []byte{'C', 'F', 'G'}

	minorCount := 0
	majorCount := 0

	for _, line := range lines {
		for _, minor := range minors {
			if line[0] == minor {
				minorCount++
				break
			}
		}

		for _, major := range majors {
			if line[0] == major {
				majorCount++
				break
			}
		}
	}

	if minorCount == majorCount {
		for _, minor := range minors {
			if lines[len(lines)-1][len(lines[len(lines)-1])-1] == minor {
				minorCount++
				break
			}
		}

		for _, major := range majors {
			if lines[len(lines)-1][len(lines[len(lines)-1])-1] == major {
				majorCount++
				break
			}
		}
	}

	if minorCount < majorCount {
		fmt.Println("C-major")
	} else {
		fmt.Println("A-minor")
	}
}
