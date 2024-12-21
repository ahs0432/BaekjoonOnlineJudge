package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	ans := 0
	var s string

	for {
		s, _ = reader.ReadString('\n')
		s = strings.TrimSuffix(s, "\n")

		if s == "고무오리 디버깅 시작" {
			continue
		} else if s == "고무오리 디버깅 끝" {
			break
		} else if s == "문제" {
			ans++
		} else if s == "고무오리" {
			if ans == 0 {
				ans += 2
			} else {
				ans--
			}
		}
	}

	if ans != 0 {
		fmt.Println("힝구")
	} else {
		fmt.Println("고무오리야 사랑해")
	}
}
