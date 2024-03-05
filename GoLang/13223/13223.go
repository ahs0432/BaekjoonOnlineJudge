package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var curtime, salttime string
	fmt.Fscanln(reader, &curtime)
	fmt.Fscanln(reader, &salttime)

	cts := strings.Split(curtime, ":")
	sts := strings.Split(salttime, ":")

	ctsh, _ := strconv.Atoi(cts[0])
	ctsm, _ := strconv.Atoi(cts[1])
	ctss, _ := strconv.Atoi(cts[2])

	stsh, _ := strconv.Atoi(sts[0])
	stsm, _ := strconv.Atoi(sts[1])
	stss, _ := strconv.Atoi(sts[2])

	curSec := ctsh*3600 + ctsm*60 + ctss
	saltSec := stsh*3600 + stsm*60 + stss

	diffSec := saltSec - curSec
	if diffSec <= 0 {
		diffSec += 86400
	}

	fmt.Printf("%02d:%02d:%02d", diffSec/3600, (diffSec%3600)/60, (diffSec%3600)%60)
}
