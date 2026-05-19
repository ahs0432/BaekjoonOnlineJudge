package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var check bool
	var str1, str2 string
	var count1, count2 []int
	for i := 1; ; i++ {
		fmt.Fscanln(reader, &str1)
		fmt.Fscanln(reader, &str2)

		if str1 == "END" && str2 == "END" {
			break
		}

		count1 = make([]int, 26)
		count2 = make([]int, 26)

		for j := 0; j < len(str1); j++ {
			count1[str1[j]-'a']++
		}

		for j := 0; j < len(str2); j++ {
			count2[str2[j]-'a']++
		}

		check = true
		for j := 0; j < 26; j++ {
			if count1[j] != count2[j] {
				check = false
				break
			}
		}

		if check {
			fmt.Fprintf(writer, "Case %d: same\n", i)
		} else {
			fmt.Fprintf(writer, "Case %d: different\n", i)
		}
	}
	writer.Flush()
}
