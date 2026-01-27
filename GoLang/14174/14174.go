package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	var a, b string
	var tmp1, tmp2 []int
	ans := make([]int, 26)
	for i := 0; i < n; i++ {
		tmp1, tmp2 = make([]int, 26), make([]int, 26)
		fmt.Fscanln(reader, &a, &b)

		for j := 0; j < len(a); j++ {
			tmp1[a[j]-'a']++
		}
		for j := 0; j < len(b); j++ {
			tmp2[b[j]-'a']++
		}
		for j := 0; j < 26; j++ {
			ans[j] += max(tmp1[j], tmp2[j])
		}
	}

	for i := 0; i < 26; i++ {
		fmt.Println(ans[i])
	}
}
