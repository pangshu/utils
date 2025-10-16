package main

import (
	"fmt"
	DuduLog "utils/Log"
)

var (
	Log DuduLog.Log
)

func main() {
	Log.New(nil)

	fmt.Println("hello world")
}
