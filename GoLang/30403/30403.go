package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	var s string
	fmt.Fscanln(reader, &n)
	fmt.Fscanln(reader, &s)

	upperCase := []byte{'R', 'O', 'Y', 'G', 'B', 'I', 'V'}
	lowerCase := []byte{'r', 'o', 'y', 'g', 'b', 'i', 'v'}

	for i := 0; i < len(s); i++ {
		for j, upper := range upperCase {
			if upper == s[i] {
				upperCase = append(upperCase[0:j], upperCase[j+1:]...)
				break
			}
		}

		for j, lower := range lowerCase {
			if lower == s[i] {
				lowerCase = append(lowerCase[0:j], lowerCase[j+1:]...)
				break
			}
		}
	}

	if len(upperCase) == 0 && len(lowerCase) == 0 {
		fmt.Println("YeS")
	} else if len(upperCase) == 0 {
		fmt.Println("YES")
	} else if len(lowerCase) == 0 {
		fmt.Println("yes")
	} else {
		fmt.Println("NO!")
	}
}
