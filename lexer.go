package main

type Lexer struct {
  input string    // the string being scanned.
  pos   int       // current position in the input.
  start int       // start position of this item.
  width int       // width of the last rune read from input.
}

func NewLexer(str string) *Lexer {
  // TODO: Remove unnecessary whitespace from string
  return &Lexer{ input: str, pos: 0 }
}

func (l *lexer) next() Token {
}
