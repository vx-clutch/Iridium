package main

import (
	"filler/pkg/shell"
	"fmt"
	"os"
)

func main() {
	var arg0 string
	if len(os.Args) > 1 {
		arg0 = os.Args[1]
	}
	if len(os.Args) > 2 {
		fmt.Println("error: Too many args")
	}
	shell.Parse(&arg0)
}
