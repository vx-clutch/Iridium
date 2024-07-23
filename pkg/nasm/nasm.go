package nasm

import (
	"fmt"
	"runtime"
)

func Nasm(asm string) {
	os := runtime.GOOS
	var nasmFlags string
	switch os {
	case "linux":
		nasmFlags = "-f elf64"
	case "windows":
		nasmFlags = "-f win64"
	case "darwin":
		nasmFlags = "-f macho64"
	default:
		fmt.Printf("Unsupported OS: %s\n", os)
		return
	}
	cmd := fmt.Sprintf("nasm %v -o a.o %v; ld -o a.out a.o", nasmFlags, asm)
}
