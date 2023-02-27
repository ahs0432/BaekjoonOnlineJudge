package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().In(time.FixedZone("GMT", 0)).Format("2006\n01\n02"))
}
