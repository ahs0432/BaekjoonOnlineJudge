package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscanln(reader, &n)

	str := []string{"SciComLove", "ciComLoveS", "iComLoveSc", "ComLoveSci", "omLoveSciC", "mLoveSciCo", "LoveSciCom", "oveSciComL", "veSciComLo", "eSciComLov"}
	fmt.Println(str[n%10])
}
