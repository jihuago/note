package main

import (
	"fmt"
	"time"
)

func main() {
	t, _ := time.Parse("2006-01-02 15:04:05", "2021-10-26 17:07:00")
	fmt.Println(t)
}
