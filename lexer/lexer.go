package lexer

import (
	"github.com/nwehr/monkey/token"
)

type Lexer struct {
	input string
	pos   int
}

func New(input string) *Lexer {
	return &Lexer{input: input}
}

func (lex *Lexer) NextToken() token.Token {
	lex.skipWhitespace()
	ch := lex.nextChar()

	switch ch {
	case '=':
		if lex.peekChar() == '=' {
			lex.nextChar()
			return token.Token{token.EQ, "=="}
		}
		return token.Token{token.ASSIGN, string(ch)}
	case '+':
		return token.Token{token.PLUS, string(ch)}
	case '-':
		return token.Token{token.MINUS, string(ch)}
	case '*':
		return token.Token{token.ASTERISK, string(ch)}
	case '/':
		return token.Token{token.SLASH, string(ch)}
	case '!':
		if lex.peekChar() == '=' {
			lex.nextChar()
			return token.Token{token.NOT_EQ, "!="}
		}
		return token.Token{token.BANG, string(ch)}
	case '<':
		return token.Token{token.LT, string(ch)}
	case '>':
		return token.Token{token.GT, string(ch)}
	case ',':
		return token.Token{token.COMMA, string(ch)}
	case ';':
		return token.Token{token.SEMICOLON, string(ch)}
	case '(':
		return token.Token{token.LPAREN, string(ch)}
	case ')':
		return token.Token{token.RPAREN, string(ch)}
	case '{':
		return token.Token{token.LBRACE, string(ch)}
	case '}':
		return token.Token{token.RBRACE, string(ch)}
	case 0:
		return token.Token{token.EOF, ""}
	}

	if isLetter(ch) {
		literal := lex.readIdent()
		return token.Token{token.LookupIdent(literal), literal}
	}

	if isDigit(ch) {
		return token.Token{token.INT, lex.readNumber()}
	}

	return token.Token{token.ILLEGAL, string(ch)}
}

func (lex *Lexer) readIdent() string {
	start := lex.pos - 1

	for isLetter(lex.peekChar()) {
		lex.nextChar()
	}

	return lex.input[start:lex.pos]
}

func (lex *Lexer) readNumber() string {
	start := lex.pos - 1

	for isDigit(lex.peekChar()) {
		lex.nextChar()
	}

	return lex.input[start:lex.pos]
}

func (lex *Lexer) nextChar() byte {
	pos := lex.pos
	lex.pos++

	if pos >= len(lex.input) {
		return 0
	}

	return lex.input[pos]
}

func (lex *Lexer) peekChar() byte {
	if lex.pos >= len(lex.input) {
		return 0
	}

	return lex.input[lex.pos]
}

func (lex *Lexer) skipWhitespace() {
	for lex.peekChar() == ' ' || lex.peekChar() == '\t' || lex.peekChar() == '\n' || lex.peekChar() == '\r' {
		lex.nextChar()
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
