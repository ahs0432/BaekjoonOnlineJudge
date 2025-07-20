package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var birth string
	fmt.Fscanln(reader, &birth)

	var n int
	fmt.Fscanln(reader, &n)

	var tmp string
	var a, b, c, key, value int
	ans, check := -1, -1
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &tmp)

		a, b, c = 0, 0, 0

		for j := 0; j < 4; j++ {
			bDayDigit, _ := strconv.Atoi(string(birth[j]))
			codingDigit, _ := strconv.Atoi(string(tmp[j]))
			diff := bDayDigit - codingDigit
			a += diff * diff
		}

		for j := 4; j < 6; j++ {
			bDayDigit, _ := strconv.Atoi(string(birth[j]))
			codingDigit, _ := strconv.Atoi(string(tmp[j]))
			diff := bDayDigit - codingDigit
			b += diff * diff
		}

		for j := 6; j < 8; j++ {
			bDayDigit, _ := strconv.Atoi(string(birth[j]))
			codingDigit, _ := strconv.Atoi(string(tmp[j]))
			diff := bDayDigit - codingDigit
			c += diff * diff
		}

		key, _ = strconv.Atoi(tmp)
		value = a * b * c

		if check != -1 {
			if value > check {
				check = value
				ans = key
			} else if value == check && key < ans {
				ans = key
			}
		} else {
			check = value
			ans = key
		}
	}
	fmt.Println(ans)
}
