package main

import (
	"fmt"
	"strings"
)

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
	String
	Number
	EOF
	Identifier
	// keywords
	And
	Class
	Else
	False
	For
	Fun
	If
	Nil
	Or
	Print
	Return
	Super
	This
	True
	Var
	While
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
	String:       "STRING",
	Number:       "NUMBER",
	EOF:          "EOF",
	Identifier:   "IDENTIFIER",
	And:          "AND",
	Class:        "CLASS",
	Else:         "ELSE",
	False:        "FALSE",
	For:          "FOR",
	Fun:          "FUN",
	If:           "IF",
	Nil:          "NIL",
	Or:           "OR",
	Print:        "PRINT",
	Return:       "RETURN",
	Super:        "SUPER",
	This:         "THIS",
	True:         "TRUE",
	Var:          "VAR",
	While:        "WHILE",
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
	var l string
	switch v := t.Literal.(type) {
	case nil:
		l = "null"
	case string:
		l = v
	case float64:
		l = fmt.Sprintf("%v", v)
		if !strings.Contains(l, ".") {
			l = fmt.Sprintf("%s.0", l)
		}
	default:
		l = fmt.Sprintf("%v", v)
	}
	return fmt.Sprintf("%s %s %s", t.Type.String(), t.Lexeme, l)
}
