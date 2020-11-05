package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	for {
		fmt.Println("hi", i)
		i++
		time.Sleep(time.Second * 5)
	}
}
