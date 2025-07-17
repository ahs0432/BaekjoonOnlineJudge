package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var r, c, ac, ad, sr, sc int
	fmt.Fscanln(reader, &r, &c)
	fmt.Fscanln(reader, &ac, &ad)
	fmt.Fscanln(reader, &sr, &sc)

	if sr == r && (ad+sr)&1 != 0 {
		fmt.Println("YES!")
	} else {
		fmt.Println("NO...")
	}
}
