package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var s string
	fmt.Fscanln(reader, &s)

	robot, box, flag := 0, 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '@' {
			robot = i
		} else if s[i] == '#' {
			box = i
		} else if s[i] == '!' {
			flag = i
		}
	}

	ans := 0
	if (robot < box && box < flag) || (robot > box && box > flag) {
		if box-robot < 0 {
			ans += robot - box - 1
		} else {
			ans += box - robot - 1
		}

		if flag-box < 0 {
			ans += box - flag
		} else {
			ans += flag - box
		}
		fmt.Println(ans)
	} else {
		fmt.Println(-1)
	}
}
