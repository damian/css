package main

import (
	"bufio"
	// "fmt"
	"io"
	"unicode"
)

type Scanner struct {
	r *bufio.Reader
}

func (s *Scanner) Emit(t TokenType, v string) *Token {
	return &Token{t, v}
}

var eof = rune(0)

func (s *Scanner) Read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

// func (s *Scanner) String() string {
// 	return fmt.Sprintf("str: %s, ind: %d", s.s, s.i)
// }

func (s *Scanner) Scan() *Token {
	ch := s.Read()

	if isWhitespace(ch) {
		return s.scanWhitespace()
	}

	switch ch {
	case eof:
		return s.Emit(TokenEOF, "")
	case ',':
		return s.Emit(TokenComma, string(ch))
	case '{':
		return s.Emit(TokenLeftCurly, string(ch))
	case '}':
		return s.Emit(TokenRightCurly, string(ch))
	case ':':
		return s.Emit(TokenColon, string(ch))
	case ';':
		return s.Emit(TokenSemicolon, string(ch))
	case '#':
		// This is an identifier(ID)
		if isIdentStart(s) || isValidEscape(s) {
			// TODO: DN - Need to consume a name here
			return s.Emit(TokenHash, string(ch))
		} else {
			return s.Emit(TokenDelim, string(ch))
		}
	default:
		return s.Emit(TokenString, string(ch))
	}
}

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func isIdentStart(s *Scanner) bool {
	_, codePointTwo := s.Read(), s.Read()
	s.Unread()
	s.Unread()

	if unicode.IsLetter(codePointTwo) || unicode.IsNumber(codePointTwo) {
		return true
	}

	str := string(codePointTwo)
	if str == "_" || str == "-" {
		return true
	}

	// Is this none ASCII
	return uint32(codePointTwo) > '\x7F'
}

func isValidEscape(s *Scanner) bool {
	_, codePointTwo, codePointThree := s.Read(), s.Read(), s.Read()
	s.Unread()
	s.Unread()
	s.Unread()

	if uint32(codePointTwo) != '\\' {
		return false
	}

	if uint32(codePointThree) == '\n' {
		return false
	}

	return true
}

func (s *Scanner) scanWhitespace() *Token {
	var ch rune
	for isWhitespace(s.peek()) {
		ch = s.Read()
	}
	return s.Emit(TokenWhitespace, string(ch))
}

func (s *Scanner) Unread() error {
	return s.r.UnreadRune()
}

func (s *Scanner) peek() rune {
	ch := s.Read()
	s.Unread()
	return ch
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}
