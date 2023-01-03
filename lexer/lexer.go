package lexer

import (
	"fmt"
	"strings"
)

type TokenType string

const (
	CODE   TokenType = "code"
	EOF    TokenType = "eof"
	COLON  TokenType = "colon"
	COMMA  TokenType = "comma"
	NUMBER TokenType = "number"
	STRING TokenType = "string"
	EXTM3U TokenType = "EXT3MU"
	EXTINF TokenType = "EXTINF"
)

type Lexer struct {
	S      []rune
	Cursor int
}

type Token struct {
	Type  TokenType
	Value string
}

func (l Lexer) RunLexer() ([]Token, error) {
	var tokens = make([]Token, 0)
	token, err := l.getNextToken()
	if err != nil {
		return []Token{}, err
	}
	for token.Type != EOF {
		tokens = append(tokens, token)
		token, err = l.getNextToken()
		if err != nil {
			return []Token{}, err
		}
	}
	return tokens, nil
}

func (l *Lexer) getNextToken() (Token, error) {
	str := l.S
	if l.Cursor >= len(str) {
		return Token{EOF, "EOF"}, nil
	}
	// skips whitespaces and newline
	for string(str[l.Cursor]) == " " || string(str[l.Cursor]) == "\n" {
		l.Cursor++
		if l.Cursor == len(str) {
			return Token{EOF, ""}, nil
		}
	}
	// recognises colon
	if string(str[l.Cursor]) == ":" {
		l.Cursor++
		return Token{COLON, ":"}, nil
	}

	// recognises comma
	if string(str[l.Cursor]) == "," {
		l.Cursor++
		return Token{COMMA, ","}, nil
	}

	// recognises numeric and float literals
	if idx := strings.IndexRune("0123456789", str[l.Cursor]); idx != -1 {
		num := ""
		float := 0
		for l.Cursor < len(str) && string(str[l.Cursor]) != "\n" && string(str[l.Cursor]) != ":" && string(str[l.Cursor]) != " " && string(str[l.Cursor]) != "," {
			idx = strings.IndexRune("0123456789", str[l.Cursor])
			if idx == -1 {
				if string(str[l.Cursor]) == "." && len(num) > 0 && float == 0 {
					float = 1
					num += string(str[l.Cursor])
					l.Cursor++
					continue
				}
				return Token{}, fmt.Errorf("expected numeric literal, got %s", string(str[l.Cursor]))
			}
			num += string(str[l.Cursor])
			l.Cursor++
		}
		return Token{NUMBER, num}, nil
	}

	// recogises codes eg: #EXTINF, etc
	if string(str[l.Cursor]) == "#" {
		s := ""
		for l.Cursor < len(str) && string(str[l.Cursor]) != "\n" && string(str[l.Cursor]) != ":" && string(str[l.Cursor]) != " " {
			s += string(str[l.Cursor])
			l.Cursor++
		}
		if s == "#EXTM3U" {
			return Token{EXTM3U, s}, nil
		}
		if s == "#EXTINF" {
			return Token{EXTINF, s}, nil
		}
		return Token{CODE, s}, nil
	}

	// recognises string literals
	s := ""
	for l.Cursor < len(str) && string(str[l.Cursor]) != "\n" && string(str[l.Cursor]) != ":" && string(str[l.Cursor]) != " " {
		s += string(str[l.Cursor])
		l.Cursor++
	}
	return Token{STRING, s}, nil
}
