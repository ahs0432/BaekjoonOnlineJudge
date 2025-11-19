package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	cjr := make([]float64, 6)
	ek := make([]float64, 6)

	for i := 0; i < 6; i++ {
		fmt.Fscan(reader, &cjr[i])
	}

	for i := 0; i < 6; i++ {
		fmt.Fscan(reader, &ek[i])
	}

	cjrTotal := cjr[0]*13 + cjr[1]*7 + cjr[2]*5 + cjr[3]*3 + cjr[4]*3 + cjr[5]*2
	ekTotal := ek[0]*13 + ek[1]*7 + ek[2]*5 + ek[3]*3 + ek[4]*3 + ek[5]*2 + 1.5

	if cjrTotal < ekTotal {
		fmt.Println("ekwoo")
	} else {
		fmt.Println("cocjr0208")
	}

}
