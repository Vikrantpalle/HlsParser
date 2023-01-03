package main

import (
	"fmt"
	"hls-parser/lexer"
	"hls-parser/parser"
	"os"
)

func GetSegments(url string) ([]parser.Segment, error) {
	var segments = make([]parser.Segment, 0)
	dat, err := os.ReadFile(url)
	if err != nil {
		return segments, fmt.Errorf("file error: %v", err)
	}
	var l = lexer.Lexer{S: []rune(string(dat)), Cursor: 0}
	tokens, err := l.RunLexer()
	if err != nil {
		return segments, fmt.Errorf("lexer error %v", err)
	}
	var p = parser.Parser{Tokens: tokens, Cursor: 0, Segments: &segments}
	err = p.Parse()
	if err != nil {
		return segments, fmt.Errorf("parsing error %v", err)
	}
	return segments, nil
}
