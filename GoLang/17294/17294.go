package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var k string
	fmt.Fscanln(reader, &k)

	if len(k) == 1 {
		fmt.Println("◝(⑅•ᴗ•⑅)◜..°♡ 뀌요미!!")
	} else {
		check := true
		n := int(k[0]-'0') - int(k[1]-'0')

		for i := 0; i < len(k)-1; i++ {
			if int(int(k[i]-'0')-int(k[i+1]-'0')) != n {
				fmt.Println("흥칫뿡!! <(￣ ﹌ ￣)>")
				check = false
				break
			}
		}

		if check {
			fmt.Println("◝(⑅•ᴗ•⑅)◜..°♡ 뀌요미!!")
		}
	}
}
