package main

import "fmt"

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

func (s *Scanner) Advance() byte {
	v := s.Source[s.current]
	s.current += 1
	return v
}

func (s *Scanner) AddToken(t TokenType, literal any) {
	// fmt.Fprintf(os.Stderr, "[AddToken] start:%v | current:%v\n", s.start, s.current)
	text := string(s.Source[s.start:s.current])
	s.Tokens = append(s.Tokens, Token{t, text, literal, s.line})
}

func (s *Scanner) AddError() {
	text := string(s.Source[s.start:s.current])
	msg := fmt.Sprintf("[line %d] Error: Unexpected character: %v\n", s.line, text)
	s.Errors = append(s.Errors, msg)
}

func (s *Scanner) ScanToken() {
	switch s.Advance() {
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
	default:
		s.AddError()
	}
}

func (s *Scanner) ScanTokens() []Token {
	for i := 0; !s.IsAtEnd(); i++ {
		// fmt.Fprintf(os.Stderr, "[ScanTokens] i:%v | IsAtEnd:%v\n", i, s.IsAtEnd())
		s.start = s.current
		s.ScanToken()
	}
	s.Tokens = append(s.Tokens, Token{EOF, "", nil, s.line})
	return s.Tokens
}
