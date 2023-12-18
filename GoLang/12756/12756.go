package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var aa, ah, ba, bh int
	fmt.Fscanln(reader, &aa, &ah)
	fmt.Fscanln(reader, &ba, &bh)

	as, bs := ah/ba, bh/aa
	at, bt := ah%ba, bh%aa

	if at != 0 {
		as++
	}
	if bt != 0 {
		bs++
	}

	if as > bs {
		fmt.Println("PLAYER A")
	} else if as < bs {
		fmt.Println("PLAYER B")
	} else {
		fmt.Println("DRAW")
	}
}
