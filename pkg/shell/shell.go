package shell

import (
	"fmt"
	"iridium/pkg/irep"
	"iridium/pkg/iridium"
	"strings"
)

func Parse(command string, file string) {
	if file == "" || command == "" {
		fmt.Println("Incomplete Feature: Help Text")
		return
	}
	if command == "irep" {
		if !strings.HasSuffix(file, ".irep") {
			fmt.Println(fmt.Sprintf("error: Incorrect file extention:\n\texpected: main.irep found: %v", file))
			return
		}
		irep.Compile(file)
	}
	if command == "iridium" {
		if !strings.HasSuffix(file, ".ir") {
			fmt.Println(fmt.Sprintf("error: Incorrect file extention:\n\texpected: main.ir found: %v", file))
			return
		}
		iridium.Compile(file)
	}
}
