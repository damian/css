package main

import (
	"bytes"
	"fmt"
	// "io/ioutil"
)

type TokenType int

type Token struct {
	Type  TokenType
	Value string
}

func (t Token) String() string {
	switch t.Type {
	case TokenEOF:
		return "EOF"
	default:
		return fmt.Sprintf("%s", t.Value)
	}
}

const (
	TokenEOF TokenType = iota
	TokenWhitespace
	TokenColon
	TokenSemicolon
	TokenLeftCurly
	TokenRightCurly
	TokenString
	TokenComma
	TokenHash
	TokenDelim
)

func main() {
	str := "#foo { background-color: red; }"
	// dat, _ := ioutil.ReadFile("./test.css")
	// str := string(dat)
	b := bytes.NewBufferString(str)
	scanner := NewScanner(b)
	for {
		token := scanner.Scan()
		fmt.Println(token)
		if token.Type == TokenEOF {
			break
		}
	}
}
