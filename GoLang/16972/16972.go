package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)

	var a, b int
	for i := 0; i < 814; i++ {
		a = rand.Intn(16281) - 8140
		b = rand.Intn(16281) - 8140
		fmt.Fprintln(writer, a, b)
	}
	writer.Flush()
}
