package lexer

import (
	"leiter-jakab/monkey/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() (tok token.Token) {
	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peakChar() == '=' {
			tok = l.makeTwoCharToken(token.EQ)
		} else {
			tok = token.FromByte(token.ASSIGN, l.ch)
		}
	case '+':
		tok = token.FromByte(token.PLUS, l.ch)
	case '-':
		tok = token.FromByte(token.MINUS, l.ch)
	case '/':
		tok = token.FromByte(token.SLASH, l.ch)
	case '*':
		tok = token.FromByte(token.ASTERISK, l.ch)
	case '<':
		tok = token.FromByte(token.LT, l.ch)
	case '>':
		tok = token.FromByte(token.GT, l.ch)
	case '!':
		if l.peakChar() == '=' {
			tok = l.makeTwoCharToken(token.NEQ)
		} else {
			tok = token.FromByte(token.BANG, l.ch)
		}
	case '(':
		tok = token.FromByte(token.LPAREN, l.ch)
	case ')':
		tok = token.FromByte(token.RPAREN, l.ch)
	case '{':
		tok = token.FromByte(token.LBRACE, l.ch)
	case '}':
		tok = token.FromByte(token.RBRACE, l.ch)
	case ',':
		tok = token.FromByte(token.COMMA, l.ch)
	case ';':
		tok = token.FromByte(token.SEMICOLON, l.ch)
	case 0:
		tok = token.FromString(token.EOF, "")
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readWord(isLetter)
			tok.Type = token.LookUpIdent(tok.Literal)
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readWord(isDigit)
		} else {
			tok = token.FromByte(token.ILLEGAL, l.ch)
		}
		return
	}

	l.readChar()
	return
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) peakChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) makeTwoCharToken(tt token.TokenType) token.Token {
	ch := l.ch
	l.readChar()
	return token.FromString(tt, string(ch)+string(l.ch))
}

func (l *Lexer) readWord(f func(byte) bool) string {
	position := l.position
	for f(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for isWhitespace(l.ch) {
		l.readChar()
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
