package shell

import (
	"fmt"
	"iridium/pkg/iridium"
	"strings"
)

func Parse(file *string) {
	if *file == "" {
		fmt.Println("Incomplete Feature: Help Text")
		return
	}
	if !strings.HasSuffix(*file, ".ir") {
		fmt.Println(fmt.Sprintf("error: Incorrect file extention:\n\texpected: main.ir found: %v", *file))
		return
	}
	iridium.Run(*file)
}
