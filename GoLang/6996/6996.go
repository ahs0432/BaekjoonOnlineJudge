package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(reader, &n)

	var check bool
	var str1, str2 string
	var strCount1, strCount2 []int
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &str1, &str2)

		if len(str1) != len(str2) {
			fmt.Fprintf(writer, "%s & %s are NOT anagrams.\n", str1, str2)
			continue
		}

		strCount1 = make([]int, 26)
		strCount2 = make([]int, 26)

		for _, c := range str1 {
			strCount1[c-'a']++
		}

		for _, c := range str2 {
			strCount2[c-'a']++
		}

		check = true
		for j := 0; j < 26; j++ {
			if strCount1[j] != strCount2[j] {
				check = false
				break
			}
		}

		if check {
			fmt.Fprintf(writer, "%s & %s are anagrams.\n", str1, str2)
		} else {
			fmt.Fprintf(writer, "%s & %s are NOT anagrams.\n", str1, str2)
		}
	}
	writer.Flush()
}
