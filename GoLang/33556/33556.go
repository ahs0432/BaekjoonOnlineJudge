package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func equals(a, b string) string {
	if a == "null" {
		return "NullPointerException"
	} else {
		if b == "null" {
			return "false"
		}

		if a == b {
			return "true"
		} else {
			return "false"
		}
	}
}

func equalsIgnoreCase(a, b string) string {
	if a == "null" {
		return "NullPointerException"
	} else {
		if b == "null" {
			return "false"
		}
		a = strings.ToLower(a)
		b = strings.ToLower(b)

		if a == b {
			return "true"
		} else {
			return "false"
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var a, b string
	fmt.Fscanln(reader, &a)
	fmt.Fscanln(reader, &b)

	fmt.Println(equals(a, b))
	fmt.Println(equalsIgnoreCase(a, b))
}
