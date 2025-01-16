package main

import (
	"bufio"
	"fmt"
	"os"
)

func rspCalc(ms, tk string) (int, int) {
	if ms == "R" && tk == "S" {
		return 1, 0
	} else if ms == "S" && tk == "P" {
		return 1, 0
	} else if ms == "P" && tk == "R" {
		return 1, 0
	} else if ms == "S" && tk == "R" {
		return 0, 1
	} else if ms == "P" && tk == "S" {
		return 0, 1
	} else if ms == "R" && tk == "P" {
		return 0, 1
	} else {
		return 0, 0
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	ms := make([]string, 2)
	tk := make([]string, 2)
	fmt.Fscanln(reader, &ms[0], &ms[1], &tk[0], &tk[1])

	score := make([]int, 2)
	var tmpMS, tmpTK int
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			tmpMS, tmpTK = rspCalc(ms[i], tk[j])
			score[0] += tmpMS
			score[1] += tmpTK
		}
	}

	if score[0] == 4 {
		fmt.Println("MS")
	} else if score[1] == 4 {
		fmt.Println("TK")
	} else {
		if ms[0] == ms[1] && score[1] >= 2 {
			fmt.Println("TK")
		} else if tk[0] == tk[1] && score[0] >= 2 {
			fmt.Println("MS")
		} else {
			fmt.Println("?")
		}
	}
}
