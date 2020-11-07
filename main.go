package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	for {
		fmt.Println("hi-test2", i)
		i++
		time.Sleep(time.Second * 5)
	}
}
