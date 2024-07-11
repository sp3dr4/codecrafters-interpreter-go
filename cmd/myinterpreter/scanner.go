package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type Scanner struct {
	Source  []byte
	Tokens  []Token
	Errors  []string
	start   int
	current int
	line    int
}

func NewScanner(source []byte) *Scanner {
	return &Scanner{
		Source:  source,
		Tokens:  make([]Token, 0),
		start:   0,
		current: 0,
		line:    1,
	}
}

func (s *Scanner) IsAtEnd() bool {
	return s.current >= len(s.Source)
}

func IsDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func (s *Scanner) Peek() byte {
	if s.IsAtEnd() {
		var bnil byte
		return bnil
	}
	return s.Source[s.current]
}

func (s *Scanner) PeekNext() byte {
	if s.current+1 >= len(s.Source) {
		var bnil byte
		return bnil
	}
	return s.Source[s.current+1]
}

func (s *Scanner) Match(expected byte) bool {
	if s.IsAtEnd() {
		return false
	}
	if s.Source[s.current] != expected {
		return false
	}
	s.current += 1
	return true
}

func (s *Scanner) Advance() byte {
	v := s.Source[s.current]
	s.current += 1
	return v
}

func (s *Scanner) CurrentSubstr() string {
	return string(s.Source[s.start:s.current])
}

func (s *Scanner) AddToken(t TokenType, literal any) {
	text := s.CurrentSubstr()
	s.Tokens = append(s.Tokens, Token{t, text, literal, s.line})
}

func (s *Scanner) AddString() {
	for s.Peek() != '"' && !s.IsAtEnd() {
		if s.Peek() == '\n' {
			s.line += 1
		}
		s.Advance()
	}

	if s.IsAtEnd() {
		s.Errors = append(s.Errors, fmt.Sprintf("[line %d] Error: Unterminated string.\n", s.line))
	} else {
		s.Advance()
		val := string(s.Source[s.start+1 : s.current-1])
		s.AddToken(String, val)
	}
}

func (s *Scanner) AddNumber() {
	for IsDigit(s.Peek()) {
		s.Advance()
	}

	fmt.Fprintln(os.Stderr, "read left part of number literal", s.start, s.current, s.CurrentSubstr())

	if s.Peek() == '.' && IsDigit(s.PeekNext()) {
		fmt.Fprintln(os.Stderr, "saw dot followed by digit", s.start, s.current, s.CurrentSubstr())
		s.Advance()
		for IsDigit(s.Peek()) {
			fmt.Fprintln(os.Stderr, "\treading decimal digits", s.start, s.current, s.CurrentSubstr(), string(s.Peek()), IsDigit(s.Peek()))
			s.Advance()
		}
	}

	raw := s.CurrentSubstr()
	n, err := strconv.ParseFloat(raw, 64)
	if err != nil {
		log.Fatalf("unable to parse number literal %v: %v", raw, err)
	}
	fmt.Fprintf(os.Stderr, "raw: %v -> float: %v\n", raw, n)
	s.AddToken(Number, n)
}

func (s *Scanner) ScanToken() {
	c := s.Advance()
	switch c {
	case '(':
		s.AddToken(LeftParen, nil)
	case ')':
		s.AddToken(RightParen, nil)
	case '{':
		s.AddToken(LeftBrace, nil)
	case '}':
		s.AddToken(RightBrace, nil)
	case ',':
		s.AddToken(Comma, nil)
	case '.':
		s.AddToken(Dot, nil)
	case '-':
		s.AddToken(Minus, nil)
	case '+':
		s.AddToken(Plus, nil)
	case ';':
		s.AddToken(Semicolon, nil)
	case '*':
		s.AddToken(Star, nil)
	case '!':
		ttype := Bang
		if s.Match('=') {
			ttype = BangEqual
		}
		s.AddToken(ttype, nil)
	case '=':
		ttype := Equal
		if s.Match('=') {
			ttype = EqualEqual
		}
		s.AddToken(ttype, nil)
	case '<':
		ttype := Less
		if s.Match('=') {
			ttype = LessEqual
		}
		s.AddToken(ttype, nil)
	case '>':
		ttype := Greater
		if s.Match('=') {
			ttype = GreaterEqual
		}
		s.AddToken(ttype, nil)
	case '/':
		if s.Match('/') {
			for s.Peek() != '\n' && !s.IsAtEnd() {
				s.Advance()
			}
		} else {
			s.AddToken(Slash, nil)
		}
	case ' ':
	case '\r':
	case '\t':
	case '\n':
		s.line += 1
	case '"':
		s.AddString()
	default:
		if IsDigit(c) {
			s.AddNumber()
		} else {
			text := string(s.Source[s.start:s.current])
			msg := fmt.Sprintf("[line %d] Error: Unexpected character: %v\n", s.line, text)
			s.Errors = append(s.Errors, msg)
		}
	}
}

func (s *Scanner) ScanTokens() []Token {
	for i := 0; !s.IsAtEnd(); i++ {
		s.start = s.current
		s.ScanToken()
	}
	s.Tokens = append(s.Tokens, Token{EOF, "", nil, s.line})
	return s.Tokens
}
