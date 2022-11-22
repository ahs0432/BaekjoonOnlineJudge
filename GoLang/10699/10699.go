package main

import (
	"fmt"
	"time"
)

func main() {
	loc, _ := time.LoadLocation("Asia/Seoul")
	fmt.Println(time.Now().In(loc).Format("2006-01-02"))
}
