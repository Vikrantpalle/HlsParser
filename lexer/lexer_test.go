package lexer_test

import (
	"os"
	"reflect"
	"testing"

	"github.com/vikrantpalle/hlsParser/lexer"
)

func TestNumberToken(t *testing.T) {
	var exp = lexer.Token{lexer.NUMBER, "44"}
	var tokens = make([]lexer.Token, 0)
	var l = lexer.Lexer{S: []rune(" 44 "), Cursor: 0, Tokens: &tokens}
	err := l.RunLexer()
	if err != nil {
		t.Errorf("%v", err)
	}
	if len(tokens) == 1 && reflect.DeepEqual(exp, tokens[0]) {
		return
	}
	t.Errorf("Expected %v but got: %v", exp, tokens)
}

func TestColonToken(t *testing.T) {
	var exp = lexer.Token{lexer.COLON, ":"}
	var tokens = make([]lexer.Token, 0)
	var l = lexer.Lexer{S: []rune(" : "), Cursor: 0, Tokens: &tokens}
	err := l.RunLexer()
	if err != nil {
		t.Errorf("%v", err)
	}
	if len(tokens) == 1 && reflect.DeepEqual(exp, tokens[0]) {
		return
	}
	t.Errorf("Expected %v but got: %v", exp, tokens)
}

func TestCommaToken(t *testing.T) {
	var exp = lexer.Token{lexer.COMMA, ","}
	var tokens = make([]lexer.Token, 0)
	var l = lexer.Lexer{S: []rune(" , "), Cursor: 0, Tokens: &tokens}
	err := l.RunLexer()
	if err != nil {
		t.Errorf("%v", err)
	}
	if len(tokens) == 1 && reflect.DeepEqual(exp, tokens[0]) {
		return
	}
	t.Errorf("Expected %v but got: %v", exp, tokens)
}

func TestStringToken(t *testing.T) {
	var exp = lexer.Token{lexer.STRING, "abcd"}
	var tokens = make([]lexer.Token, 0)
	var l = lexer.Lexer{S: []rune(" abcd "), Cursor: 0, Tokens: &tokens}
	err := l.RunLexer()
	if err != nil {
		t.Errorf("%v", err)
	}
	if len(tokens) == 1 && reflect.DeepEqual(exp, tokens[0]) {
		return
	}
	t.Errorf("Expected %v but got: %v", exp, tokens)
}

func TestExtm3uToken(t *testing.T) {
	var exp = lexer.Token{lexer.EXTM3U, "#EXTM3U"}
	var tokens = make([]lexer.Token, 0)
	var l = lexer.Lexer{S: []rune(" #EXTM3U "), Cursor: 0, Tokens: &tokens}
	err := l.RunLexer()
	if err != nil {
		t.Errorf("%v", err)
	}
	if len(tokens) == 1 && reflect.DeepEqual(exp, tokens[0]) {
		return
	}
	t.Errorf("Expected %v but got: %v", exp, tokens)
}

func TestExtinfToken(t *testing.T) {
	var exp = lexer.Token{lexer.EXTINF, "#EXTINF"}
	var tokens = make([]lexer.Token, 0)
	var l = lexer.Lexer{S: []rune(" #EXTINF "), Cursor: 0, Tokens: &tokens}
	err := l.RunLexer()
	if err != nil {
		t.Errorf("%v", err)
	}
	if len(tokens) == 1 && reflect.DeepEqual(exp, tokens[0]) {
		return
	}
	t.Errorf("Expected %v but got: %v", exp, tokens)
}

func TestCodeToken(t *testing.T) {
	var exp = lexer.Token{lexer.CODE, "#EXT-X-VERSION"}
	var tokens = make([]lexer.Token, 0)
	var l = lexer.Lexer{S: []rune(" #EXT-X-VERSION "), Cursor: 0, Tokens: &tokens}
	err := l.RunLexer()
	if err != nil {
		t.Errorf("%v", err)
	}
	if len(tokens) == 1 && reflect.DeepEqual(exp, tokens[0]) {
		return
	}
	t.Errorf("Expected %v but got: %v", exp, tokens)
}

func BenchmarkRunLexer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dat, _ := os.ReadFile("../file.m3u8")
		var tokens = make([]lexer.Token, 0)
		var l = lexer.Lexer{S: []rune(string(dat)), Cursor: 0, Tokens: &tokens}
		l.RunLexer()
	}
}
