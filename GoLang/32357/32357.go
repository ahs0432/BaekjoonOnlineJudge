package main

import (
	"bufio"
	"fmt"
	"os"
)

func isPalindrome(s string) bool {
	n := len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-1-i] {
			return false
		}
	}
	return true
}

func solution(strs []string) int {
	count := 0
	for _, s := range strs {
		if isPalindrome(s) {
			count++
		}
	}
	return count * (count - 1)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	strs := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &strs[i])
	}
	fmt.Println(solution(strs))
}
