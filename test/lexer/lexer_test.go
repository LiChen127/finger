package test

import (
	"finger/lexer"
	"finger/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `
	let x = 5
	let y = 10
	const z = "hello";
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "x"},
		{token.ASSIGN, "="},
		{token.NUMBER, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "y"},
		{token.ASSIGN, "="},
		{token.NUMBER, "10"},	
		{token.SEMICOLON, ";"},
		{token.CONST, "const"},
		{token.IDENT, "z"},
		{token.ASSIGN, "="},
		{token.STRING, "hello"},
		{token.SEMICOLON, ";"},
	}

	l := lexer.New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}