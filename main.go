package main

import (
  "fmt"
  "io"
  "unicode/utf8"
)

type TokenType int

type Token struct {
  Type TokenType
  Value string
}

func (t Token) String() string {
  switch t.Type {
  case TokenEOF:
    return "EOF"
  default:
    return fmt.Sprintf("<%s>", t.Value)
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
)

type Lexer struct {
  s string
  i int64
}

func (l *Lexer) Emit(t TokenType, v string) *Token {
  return &Token{t, v}
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

func (l *Lexer) Lex() *Token {
  r, _, err := l.Next()
  if err == io.EOF {
    return l.Emit(TokenEOF, "")
  }
  switch r {
  case ' ', '\t', '\n':
    return l.Emit(TokenWhitespace, string(r))
  case '{':
    return l.Emit(TokenLeftCurly, string(r))
  case '}':
    return l.Emit(TokenRightCurly, string(r))
  case ':':
    return l.Emit(TokenColon, string(r))
  case ';':
    return l.Emit(TokenSemicolon, string(r))
  default:
    return l.Emit(TokenString, string(r))
  }
}

func NewLexer(s string) *Lexer { return &Lexer{s, 0} }

func main() {
  str := "body { background-color: red; }"
  lexer := NewLexer(str)
  for {
    token := lexer.Lex()
    fmt.Println(token)
    if token.Type == TokenEOF {
      break
    }
  }
}
