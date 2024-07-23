package nasm

import (
	"fmt"
	"os/exec"
	"runtime"
)

func Nasm(asm string) {
	runos := runtime.GOOS
	var nasmFlags string
	switch runos {
	case "linux":
		nasmFlags = "-f elf64"
	case "windows":
		nasmFlags = "-f win64"
	case "darwin":
		nasmFlags = "-f macho64"
	default:
		fmt.Printf("Unsupported OS: %s\n", runos)
		return
	}
	cmd := exec.Command("nasm", nasmFlags, "-o", "main.o", asm)
	_, err := cmd.Output()
	if err != nil {
		panic(err)
	}
}
