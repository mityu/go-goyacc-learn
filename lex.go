package main

import "text/scanner"

type Token struct {
	token   int
	literal string
	// TODO: file, line, column
}

type Lexer struct {
	scanner.Scanner
	result Expr
}

func (l *Lexer) Lex(lval *yySymType) int {
	token := int(l.Scan())
	if token == scanner.Int {
		token = NUMBER
	}
	lval.token = Token{token: token, literal: l.TokenText()}
	return token
}

func (l *Lexer) Error(e string) {
	panic(e)
}
