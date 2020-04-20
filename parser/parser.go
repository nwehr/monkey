package parser

import (
	"fmt"
	"github.com/nwehr/monkey/ast"
	"github.com/nwehr/monkey/lexer"
	"github.com/nwehr/monkey/token"
)

type Parser struct {
	lex *lexer.Lexer

	currToken token.Token
	peekToken token.Token
}

func New(lex *lexer.Lexer) *Parser {
	p := &Parser{lex: lex}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) ParseProgram() (ast.Program, error) {
	program := ast.Program{}

	for p.currToken.Type != token.EOF {
		stmt, err := p.parseStatement()
		if err != nil {
			return program, err
		}
		program.Statements = append(program.Statements, stmt)
		p.nextToken()
	}

	return program, nil
}

func (p *Parser) parseStatement() (ast.Statement, error) {
	switch p.currToken.Type {
	case token.LET:
		return p.parseLetStatement()
	}

	return nil, fmt.Errorf("Parse error; Unknown token \"%s\"", p.currToken.Literal)
}

func (p *Parser) parseLetStatement() (ast.LetStatement, error) {
	letStmt := ast.LetStatement{Token: p.currToken}

	if p.peekToken.Type != token.IDENT {
		return letStmt, fmt.Errorf("Parse error")
	}

	letStmt.Name = &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}

	p.nextToken()

	if p.peekToken.Type != token.ASSIGN {
		return letStmt, fmt.Errorf("Parse error; Expecting assignment operator")
	}

	p.nextToken()

	for p.currToken.Type != token.SEMICOLON {
		p.nextToken()
	}

	return letStmt, nil
}

func (p *Parser) nextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.lex.NextToken()
}
