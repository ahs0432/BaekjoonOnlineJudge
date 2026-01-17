package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var k string
	fmt.Fscanln(reader, &k)

	k = k[2:]

	var n, d int
	if len(k) == 1 || k[0] != k[len(k)-1] {
		n, _ = strconv.Atoi(k)
		d = int(math.Pow(10, float64(len(k))))
	} else {
		n, _ = strconv.Atoi(k)
		d = 0
		for i := 0; i < len(k); i++ {
			d = (d * 10) + 9
		}
	}

	cd := gcd(n, d)
	n /= cd
	d /= cd

	fmt.Println("YES")
	fmt.Println(n, d)
}
