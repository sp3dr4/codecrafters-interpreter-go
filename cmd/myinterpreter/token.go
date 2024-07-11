package main

import "fmt"

type TokenType int

const (
	LeftParen TokenType = iota
	RightParen
	LeftBrace
	RightBrace
	Comma
	Dot
	Minus
	Plus
	Semicolon
	Slash
	Star
	Bang
	BangEqual
	Equal
	EqualEqual
	Greater
	GreaterEqual
	Less
	LessEqual
	EOF
)

var tokenTypeName = map[TokenType]string{
	LeftParen:    "LEFT_PAREN",
	RightParen:   "RIGHT_PAREN",
	LeftBrace:    "LEFT_BRACE",
	RightBrace:   "RIGHT_BRACE",
	Comma:        "COMMA",
	Dot:          "DOT",
	Minus:        "MINUS",
	Plus:         "PLUS",
	Semicolon:    "SEMICOLON",
	Slash:        "SLASH",
	Star:         "STAR",
	Bang:         "BANG",
	BangEqual:    "BANG_EQUAL",
	Equal:        "EQUAL",
	EqualEqual:   "EQUAL_EQUAL",
	Greater:      "GREATER",
	GreaterEqual: "GREATER_EQUAL",
	Less:         "LESS",
	LessEqual:    "LESS_EQUAL",
	EOF:          "EOF",
}

func (tt TokenType) String() string {
	return tokenTypeName[tt]
}

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal any
	Line    int
}

func (t *Token) String() string {
	l := "null" // t.Literal
	return fmt.Sprintf("%s %s %v", t.Type.String(), t.Lexeme, l)
}
