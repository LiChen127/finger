package lexer

import "finger/token"

type Lexer struct {
	input        string
	position     int  // 所输入字符串中的当前位置(指向当前字符)
	readPosition int  // 当前读取的下一个位置(指向当前字符的下一个字符)
	ch           byte // 当前正在查看的字符
}

/*
	创建一个词法分析器
*/
func New(input string) *Lexer {
	// 创建一个词法分析器
	l := &Lexer{input: input}
	// 读取下一个字符
	l.readChar()
	return l
}

/*
	读取下一个字符，并前移其在input中的位置
	ch = 0 意味着NIL字符,EOF, 只支持ASCII字符
*/
func (l *Lexer) readChar() {
	// 如果读取位置超过了输入字符串的长度，则将字符设置为0(表示EOF)
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		// 否则，将当前字符设置为输入字符串中readPosition位置的字符
		l.ch = l.input[l.readPosition]
	}
	// 更新读取位置
	l.position = l.readPosition
	l.readPosition++
}

/*
	窥视输入中的下一个字符，不会移动输入中的指针位置
*/
func (l *Lexer) peekChar() byte {
	// 如果读取位置超过了输入字符串的长度，则将字符设置为0(表示EOF)
	if l.readPosition >= len(l.input) {
		return 0
	}
	// 否则，返回输入字符串中readPosition位置的字符
	return l.input[l.readPosition]
}

/*
	检查当前正在查看的字符，根据字符返回相应的词法单元。
*/
func (l *Lexer) NextToken() token.Token {
	// 创建一个token
	var tok token.Token

	// 跳过空白字符
	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			// 如果下一个字符是=，则返回EQ == 
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			// 否则，返回ASSIGN
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '!':
		if l.peekChar() == '=' {
			// 如果下一个字符是=，则返回NOT_EQ !=
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			// 否则，返回BANG
			tok = newToken(token.BANG, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
		return tok
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	case '[':
		tok = newToken(token.LBRACKET, l.ch)
	case ']':
		tok = newToken(token.RBRACKET, l.ch)
	case ':':
		tok = newToken(token.COLON, l.ch)
	default:
		// 只要不是可识别的字符，就检查是否是标识符
		if isLetter(l.ch) {
			// 读入一个标识符
			tok.Literal = l.readIdentifier()
			// 根据标识符返回相应的词法单元
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			// 读入一个数字
			tok.Literal = l.readNumber()
			// 返回数字
			return tok
		} else {
			// 否则，返回非法字符
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

/*
	跳过空白字符
*/
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

/*
	读入一个数字, 只能读取整数目前，忽略了浮点数、十六进制数、八进制数
	@todo: 需要扩展
*/
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

/*
	检查字符是否是数字
*/
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

/*
	读入一个标识符并前移词法分析器的扫描位置，知道遇到非字母字符
*/
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}


/*
	检查字符是否是字母, finger处理器可处理的语言格式
*/
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

/*
	创建一个token
*/
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

/*
	读入一个字符串, 并返回字符串
*/
func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 { 
			break
		}
	}
	str := l.input[position:l.position]
	l.readChar()
	return str
}