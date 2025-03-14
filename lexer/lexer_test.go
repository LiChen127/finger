package lexer

import (
	"finger/token"
	"testing"
)

func TestSpecialIdentifiers(t *testing.T) {
	input := `
	let _privateVar = 123;
	let __temp = 456;
	let obj = {
		__proto__: baseObj
	};
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "_privateVar"},
		{token.ASSIGN, "="},
		{token.NUMBER, "123"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "__temp"},
		{token.ASSIGN, "="},
		{token.NUMBER, "456"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "obj"},
		{token.ASSIGN, "="},
		{token.LBRACE, "{"},
		{token.PROTO, "__proto__"},
		{token.COLON, ":"},
		{token.IDENT, "baseObj"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
	}

	runTokenTest(t, input, tests)
}

func TestOperators(t *testing.T) {
	input := `
	a += 1;
	b -= 2;
	c *= 3;
	d /= 4;
	e %= 5;
	f++;
	g--;
	h << 2;
	i >> 3;
	j && k;
	l || m;
	!n;
	~o;
	p ^ q;
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.IDENT, "a"},
		{token.PLUS_EQ, "+="},
		{token.NUMBER, "1"},
		{token.SEMICOLON, ";"},
		
		{token.IDENT, "b"},
		{token.MINUS_EQ, "-="},
		{token.NUMBER, "2"},
		{token.SEMICOLON, ";"},
		
		{token.IDENT, "c"},
		{token.ASTERISK_EQ, "*="},
		{token.NUMBER, "3"},
		{token.SEMICOLON, ";"},
		
		{token.IDENT, "d"},
		{token.SLASH_EQ, "/="},
		{token.NUMBER, "4"},
		{token.SEMICOLON, ";"},
		
		{token.IDENT, "e"},
		{token.MODULO_EQ, "%="},
		{token.NUMBER, "5"},
		{token.SEMICOLON, ";"},
		
		{token.IDENT, "f"},
		{token.INCREMENT, "++"},
		{token.SEMICOLON, ";"},
		
		{token.IDENT, "g"},
		{token.DECREMENT, "--"},
		{token.SEMICOLON, ";"},
		
		{token.IDENT, "h"},
		{token.BIT_SHIFT_LEFT, "<<"},
		{token.NUMBER, "2"},
		{token.SEMICOLON, ";"},
		
		{token.IDENT, "i"},
		{token.BIT_SHIFT_RIGHT, ">>"},
		{token.NUMBER, "3"},
		{token.SEMICOLON, ";"},
		
		{token.IDENT, "j"},
		{token.AND, "&&"},
		{token.IDENT, "k"},
		{token.SEMICOLON, ";"},
		
		{token.IDENT, "l"},
		{token.OR, "||"},
		{token.IDENT, "m"},
		{token.SEMICOLON, ";"},
		
		{token.BANG, "!"},
		{token.IDENT, "n"},
		{token.SEMICOLON, ";"},
		
		{token.BIT_NOT, "~"},
		{token.IDENT, "o"},
		{token.SEMICOLON, ";"},
		
		{token.IDENT, "p"},
		{token.BIT_XOR, "^"},
		{token.IDENT, "q"},
		{token.SEMICOLON, ";"},
	}

	runTokenTest(t, input, tests)
}

func TestModernFeatures(t *testing.T) {
	input := `
	async fn getData() {
		let result = await fetch();
		let value = result?.data ?? "default";
		let rest = obj;
	}
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASYNC, "async"},
		{token.FUNCTION, "fn"},
		{token.IDENT, "getData"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.AWAIT, "await"},
		{token.IDENT, "fetch"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "value"},
		{token.ASSIGN, "="},
		{token.IDENT, "result"},
		{token.OPTIONAL_CHAIN, "?."},
		{token.IDENT, "data"},
		{token.NULLISH, "??"},
		{token.STRING, "default"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "rest"},
		{token.ASSIGN, "="},
		{token.IDENT, "obj"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
	}

	runTokenTest(t, input, tests)
}

// 辅助函数
func runTokenTest(t *testing.T, input string, tests []struct {
	expectedType    token.TokenType
	expectedLiteral string
}) {
	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
} 