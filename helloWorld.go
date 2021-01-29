package main

import (
	"fmt"
	"time"
)

func main() {
	res := "hello world at " + time.Now().String()
	fmt.Println(res)
}
