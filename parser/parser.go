package parser

import (
	"errors"
	"fmt"
	"hls-parser/lexer"
	"strconv"
)

type Segment struct {
	Duration float32
	RelUrl   string
}

type Parser struct {
	Tokens   []lexer.Token
	Cursor   int
	Segments *[]Segment
}

func (p Parser) lookAhead() (lexer.Token, error) {
	if p.Cursor == len(p.Tokens) {
		return lexer.Token{}, errors.New("unexpected end of input")
	}
	return p.Tokens[p.Cursor], nil
}

func (p *Parser) eat(expType lexer.TokenType) (string, error) {
	tok, err := p.lookAhead()
	if err != nil {
		return "", err
	}
	if tok.Type == expType {
		p.Cursor++
		return tok.Value, nil
	}
	return "", fmt.Errorf("expected %s, but got %v", string(expType), tok.Type)
}

func (p *Parser) prodC() error {
	_, err := p.lookAhead()
	// end production if end of input
	if err != nil {
		return nil
	}
	_, err = p.eat(lexer.EXTINF)
	if err != nil {
		return err
	}
	_, err = p.eat(lexer.COLON)
	if err != nil {
		return err
	}
	num, err := p.eat(lexer.NUMBER)
	if err != nil {
		return err
	}
	_, err = p.eat(lexer.COMMA)
	if err != nil {
		return err
	}
	rel, err := p.eat(lexer.STRING)
	if err != nil {
		return err
	}
	dur, err := strconv.ParseFloat(num, 32)
	if err != nil {
		return fmt.Errorf("could not parse string %s to float", num)
	}
	*p.Segments = append(*p.Segments, Segment{RelUrl: rel, Duration: float32(dur)})
	err = p.prodC()
	return err
}

// currently EXTINF is compulsory
func (p *Parser) prodB() error {
	tok, err := p.lookAhead()
	// end production if end of input
	if err != nil {
		return nil
	}
	if tok.Type == lexer.CODE {
		_, err = p.eat(lexer.CODE)
		if err != nil {
			return err
		}
		_, err = p.eat(lexer.COLON)
		if err != nil {
			return err
		}
		tok, err := p.lookAhead()
		if err != nil {
			return err
		}
		if tok.Type == lexer.NUMBER {
			p.eat(lexer.NUMBER)
		} else if tok.Type == lexer.STRING {
			p.eat(lexer.STRING)
		} else {
			return fmt.Errorf("expected number or string got %s", string(tok.Type))
		}
		err = p.prodB()
		return err
	} else if tok.Type == lexer.EXTINF {
		err := p.prodC()
		return err
	} else {
		return errors.New("invalid header")
	}
}

func (p *Parser) Parse() error {
	_, err := p.eat(lexer.EXTM3U)
	if err != nil {
		return err
	}
	err = p.prodB()
	return err
}
