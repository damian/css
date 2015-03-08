package main

import (
  "fmt"
  "io"
  "unicode/utf8"
)

type Token int

const (
  TokenEOF Token = iota
  TokenWhitespace
  TokenColon
  TokenSemicolon
  TokenLeftCurly
  TokenRightCurly
  TokenString
)

type Lexer struct {
  s string
  i int64
}

func (l *Lexer) Next() (ch rune, size int, err error) {
  if l.i >= int64(len(l.s)) {
    return 0, 0, io.EOF
  }
  ch, size = utf8.DecodeRuneInString(l.s[l.i:])
  l.i += int64(size)
  return
}

func (l *Lexer) String() string {
  return fmt.Sprintf("str: %s, ind: %d", l.s, l.i)
}

func (l *Lexer) Lex() Token {
  for {
    r, _, err := l.Next()
    if err == io.EOF {
      fmt.Println("EOF-TOKEN")
      return TokenWhitespace
    }
    switch r {
    case ' ', '\t', '\n':
      fmt.Println("WHITESPACE-TOKEN")
      return TokenWhitespace
    case '{':
      fmt.Println("{-TOKEN")
      return TokenLeftCurly
    case '}':
      fmt.Println("}-TOKEN")
      return TokenRightCurly
    case ':':
      fmt.Println("COLON-TOKEN")
      return TokenColon
    case ';':
      fmt.Println("SEMICOLON-TOKEN")
      return TokenSemicolon
    default:
      fmt.Println("OTHER-TOKEN")
      return TokenString
    }
  }
}

func NewLexer(s string) *Lexer { return &Lexer{s, 0} }

func main() {
  str := "body { background-color: red; }"
  lexer := NewLexer(str)
  lexer.Lex()
}
