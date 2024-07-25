package iridium

import (
	"fmt"
	"log"
	"os"
)

func Compile(file string) {
	contents, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	m_compile(contents)
}

type token struct {
	kind  string
	value string
}

func m_compile(source []byte) {
	src := string(source)
	tokens := m_lexer(src)
	ast := m_ast(tokens)
	fmt.Println(src, tokens, ast)
}

func m_lexer(src string) []token {
	source := string(src[:])
	current := 0
	tokens := []token{}
	fmt.Println(len([]rune(source)))

	for current < len([]rune(source)) {
		char := string([]rune(source)[current])
		if char == "(" {
			tokens = append(tokens, token{kind: "paren", value: "("})
			goto next
		}
		if char == ")" {
			tokens = append(tokens, token{kind: "paren", value: ")"})
			goto next
		}
		if char == "[" {
			tokens = append(tokens, token{kind: "bracket", value: "["})
			goto next
		}
		if char == "]" {
			tokens = append(tokens, token{kind: "bracket", value: "]"})
			goto next
		}
		if char == "{" {
			tokens = append(tokens, token{kind: "brace", value: "{"})
			goto next
		}
		if char == "}" {
			tokens = append(tokens, token{kind: "brace", value: "}"})
			goto next
		}
		if char == ";" {
			tokens = append(tokens, token{kind: "semicolon", value: ";"})
			goto next
		}
		if char == ":" {
			tokens = append(tokens, token{kind: "colon", value: ":"})
			goto next
		}
		if char == "=" {
			tokens = append(tokens, token{kind: "assign", value: "="})
			goto next
		}
		if char == " " || char == "\n" || char == "\t" || char == "" {
			goto next
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
	next:
		current++
		continue
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

type node struct {
	kind       string
	value      string
	name       string
	callee     *node
	expression *node
	body       []node
	params     []node
	arguments  *[]node
	context    *[]node
}

type ast node

var pc int

var pt []token

func m_ast(tokens []token) ast {
	pc = 0
	pt = tokens

	ast := ast{
		kind: "prog",
		body: []node{},
	}
	for pc < len(pt) {
		ast.body = append(ast.body, walk())
	}
	return ast
}

func walk() node {
	token := pt[pc]
	if token.kind == "name" {
		pc++
		switch token.value {
		case "fn":
			return node{kind: "Function"}
		case "exit":
			return node{kind: "Exit"}
		case "void":
			return node{kind: "Void"}
		}
	}
	if token.kind == "semicolon" {
		pc++
		return node{
			kind:  "Semicolon",
			value: token.value,
		}
	}
	if token.kind == "number" {
		pc++
		return node{
			kind:  "NumberLiteral",
			value: token.value,
		}
	}
	if token.kind == "name" {
		pc++
		return node{
			kind:  "Ident",
			value: token.value,
		}
	}
	if token.kind == "assign" {
		pc++
		return node{
			kind:  "Assign",
			value: "=",
		}
	}
	if token.kind == "brace" {
		pc++
		token = pt[pc]
		return node{
			kind:   "ObjectLiteral",
			value:  token.value,
			params: []node{},
		}
	}
	if token.kind == "paren" && token.value == "(" {
		pc++
		token = pt[pc]
		n := node{
			kind:   "CallExpression",
			name:   token.value,
			params: []node{},
		}
		pc++
		token = pt[pc]
		for token.kind != "paren" || (token.kind == "paren" && token.value != ")") {
			n.params = append(n.params, walk())
			token = pt[pc]
		}
		pc++
		return n
	}

	log.Fatal(token.kind)
	return node{}
}
