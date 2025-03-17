package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(reader, &n)

	var s string
	s, _ = reader.ReadString('\n')
	s = strings.TrimSuffix(s, "\n")
	t := strings.Split(s, " ")

	for i := 0; i < n; i++ {
		if t[i] == ".-" {
			fmt.Fprint(writer, "A")
		} else if t[i] == "-..." {
			fmt.Fprint(writer, "B")
		} else if t[i] == "-.-." {
			fmt.Fprint(writer, "C")
		} else if t[i] == "-.." {
			fmt.Fprint(writer, "D")
		} else if t[i] == "." {
			fmt.Fprint(writer, "E")
		} else if t[i] == "..-." {
			fmt.Fprint(writer, "F")
		} else if t[i] == "--." {
			fmt.Fprint(writer, "G")
		} else if t[i] == "...." {
			fmt.Fprint(writer, "H")
		} else if t[i] == ".." {
			fmt.Fprint(writer, "I")
		} else if t[i] == ".---" {
			fmt.Fprint(writer, "J")
		} else if t[i] == "-.-" {
			fmt.Fprint(writer, "K")
		} else if t[i] == ".-.." {
			fmt.Fprint(writer, "L")
		} else if t[i] == "--" {
			fmt.Fprint(writer, "M")
		} else if t[i] == "-." {
			fmt.Fprint(writer, "N")
		} else if t[i] == "---" {
			fmt.Fprint(writer, "O")
		} else if t[i] == ".--." {
			fmt.Fprint(writer, "P")
		} else if t[i] == "--.-" {
			fmt.Fprint(writer, "Q")
		} else if t[i] == ".-." {
			fmt.Fprint(writer, "R")
		} else if t[i] == "..." {
			fmt.Fprint(writer, "S")
		} else if t[i] == "-" {
			fmt.Fprint(writer, "T")
		} else if t[i] == "..-" {
			fmt.Fprint(writer, "U")
		} else if t[i] == "...-" {
			fmt.Fprint(writer, "V")
		} else if t[i] == ".--" {
			fmt.Fprint(writer, "W")
		} else if t[i] == "-..-" {
			fmt.Fprint(writer, "X")
		} else if t[i] == "-.--" {
			fmt.Fprint(writer, "Y")
		} else if t[i] == "--.." {
			fmt.Fprint(writer, "Z")
		} else if t[i] == ".----" {
			fmt.Fprint(writer, "1")
		} else if t[i] == "..---" {
			fmt.Fprint(writer, "2")
		} else if t[i] == "...--" {
			fmt.Fprint(writer, "3")
		} else if t[i] == "....-" {
			fmt.Fprint(writer, "4")
		} else if t[i] == "....." {
			fmt.Fprint(writer, "5")
		} else if t[i] == "-...." {
			fmt.Fprint(writer, "6")
		} else if t[i] == "--..." {
			fmt.Fprint(writer, "7")
		} else if t[i] == "---.." {
			fmt.Fprint(writer, "8")
		} else if t[i] == "----." {
			fmt.Fprint(writer, "9")
		} else if t[i] == "-----" {
			fmt.Fprint(writer, "0")
		} else if t[i] == "--..--" {
			fmt.Fprint(writer, ",")
		} else if t[i] == ".-.-.-" {
			fmt.Fprint(writer, ".")
		} else if t[i] == "..--.." {
			fmt.Fprint(writer, "?")
		} else if t[i] == "---..." {
			fmt.Fprint(writer, ":")
		} else if t[i] == "-....-" {
			fmt.Fprint(writer, "-")
		} else if t[i] == ".--.-." {
			fmt.Fprint(writer, "@")
		}
	}
	fmt.Fprintln(writer)
	writer.Flush()
}
