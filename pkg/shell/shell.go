package shell

import (
	"fmt"
	"iridium/pkg/iridium"
	"iridium/pkg/onyx"
	"strings"
)

func Parse(command string, file *string) {
	if *file == "" || command == "" {
		fmt.Println("Incomplete Feature: Help Text")
		return
	}
	if command == "onyx" {
		if !strings.HasSuffix(*file, ".onyx") {
			fmt.Println(fmt.Sprintf("error: Incorrect file extention:\n\texpected: main.onyx found: %v", file))
			return
		}
		onyx.Compile(*file)
	}
	if command == "iridium" {
		if !strings.HasSuffix(*file, ".ir") {
			fmt.Println(fmt.Sprintf("error: Incorrect file extention:\n\texpected: main.ir found: %v", file))
			return
		}
		iridium.Compile(*file)
	}
}
