package shell

import (
	"filler/pkg/interpreter"
	"fmt"
	"strings"
)

func Parse(file *string) {
	if *file == "" {
		fmt.Println("Incomplete Feature: Help Text")
		return
	}
	if !strings.HasSuffix(*file, ".fl") {
		fmt.Println(fmt.Sprintf("error: Incorrect file extention:\n\texpected: main.fl found: %v", *file))
		return
	}
	interpreter.Run(*file)
}
