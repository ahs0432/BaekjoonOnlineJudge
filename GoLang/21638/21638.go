package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var t1, t2, v1, v2 int
	fmt.Fscanln(reader, &t1, &v1)
	fmt.Fscanln(reader, &t2, &v2)

	if t2 < 0 && v2 >= 10 {
		fmt.Println("A storm warning for tomorrow! Be careful and stay home if possible!")
	} else if t1 > t2 {
		fmt.Println("MCHS warns! Low temperature is expected tomorrow.")
	} else if v1 < v2 {
		fmt.Println("MCHS warns! Strong wind is expected tomorrow.")
	} else {
		fmt.Println("No message")
	}
}
