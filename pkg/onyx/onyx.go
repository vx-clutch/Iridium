package onyx

import (
	"fmt"
	"os"
)

func Compile(file string) {
	contents, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	m_compile(contents)
}

func m_compile(source []byte) {
	src := string(source)
	src += "\n"
	fmt.Println(src)
	tokens := m_lexer(src)
	fmt.Println(tokens)
}

type token struct {
	kind  string
	value string
}

func m_lexer(src string) []token {
	source := string(src[:])
	current := 0
	tokens := []token{}

	for current < len([]rune(source)) {
		char := string([]rune(source)[current])
		if char == "\t" {
			current++
			tokens = append(tokens, token{kind: "indent", value: "\t"})
		}
		if char == " " || char == "\n" || char == "" {
			current++
			continue
		}
		if isNumber(char) {
			value := ""
			for isNumber(char) {
				value += char
				current++
				char = string([]rune(source)[current])
			}
			tokens = append(tokens, token{
				kind:  "number",
				value: value,
			})
			continue
		}
		if isLetter(char) {
			value := ""
			for isLetter(char) {
				value += char
				current++
				char = string([]rune(source)[current])
			}
			tokens = append(tokens, token{
				kind:  "name",
				value: value,
			})
			continue
		}
		break
	}
	return tokens
}

func isNumber(char string) bool {
	if char == "" {
		return false
	}
	n := []rune(char)[0]
	if n >= '0' && n <= '9' {
		return true
	}
	return false
}

func isLetter(char string) bool {
	if char == "" {
		return false
	}
	n := []rune(char)[0]
	if n >= 'a' && n <= 'z' {
		return true
	}
	return false
}

type node[T any] struct {
	operation string
	args      []T
}

type ast struct {
	kind  string
	value []node
}

func m_ast(tokens []token) []ast {
	prog := ast{
		"prog",
		[]node,
	}
}
