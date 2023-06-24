package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n float64
	fmt.Fscanln(reader, &n)

	pi := 3.1415926535897932384
	fmt.Printf("%.7f\n", math.Sqrt(n/pi)*2*pi)
}
