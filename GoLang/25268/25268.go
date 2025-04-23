package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

var checkBytes []byte = []byte{'a', 'e', 'i', 'o', 'u'}

func checkIn(checkStr byte) int {
	check := 0
	for _, checkByte := range checkBytes {
		if checkByte == checkStr {
			check = 1
			break
		}
	}
	return check
}

func nameGen() string {
	name := ""

	var now byte
	var nowCheck int
	checks := []int{-1, -1}
	for i := 0; i < 8; i++ {
		now = byte(rand.Intn(26) + 'a')
		nowCheck = checkIn(now)

		if i > 1 {
			if checks[0] == checks[1] {
				for nowCheck == checks[1] {
					now = byte(rand.Intn(26) + 'a')
					nowCheck = checkIn(now)
				}
			}
		}

		name += string(now)
		checks[0] = checks[1]
		checks[1] = nowCheck
	}
	return name
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(reader, &n)

	for i := 0; i < n; i++ {
		fmt.Fprintln(writer, nameGen())
	}
	writer.Flush()
}
