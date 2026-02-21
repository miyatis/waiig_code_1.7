package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// 2つのトークンを読み込む。curToken と peekToken をセットするため。
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() *ast.Program {
	return nil
}
