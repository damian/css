package main

import (
  "fmt"
  "io/ioutil"
)

type tokenType int

type Token struct {
  typ   tokenType
  value string
}

func (token Token) String() string {
  switch token.typ {
  case tokenEOF:
    return "EOF"
  case tokenError:
    return token.value
  }
  return fmt.Sprintf("%q", token.value)
}

const (
  tokenComment tokenType = iota
  tokenNewline
  tokenWhitespace
  tokenError
  tokenEOF
)

func main() {
  stylesheet := "test.css"
  file, err := ioutil.ReadFile(stylesheet)
  if err != nil {
    fmt.Println("Error")
  }
  strFile := string(file)
  lexer := NewLexer(strFile)
  for {
    token := lexer.Next()
    if token.typ == tokenEOF {
      break
    }
  }
}
