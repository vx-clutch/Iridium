package main

import (
	"fmt"
	"iridium/pkg/shell"
	"os"
)

func main() {
	var arg0 string
	var arg1 string
	if len(os.Args) > 2 {
		arg0 = os.Args[1]
		arg1 = os.Args[2]
	}
	if len(os.Args) > 3 {
		fmt.Println("error: Too many args")
	}
	fmt.Println(arg0, arg1)
	shell.Parse(arg0, &arg1)
}
