package lexer

import (
	"finger/token"
)

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
	窥视第二个字符之后的字符
*/
func (l *Lexer) peekNextChar() byte {
	if l.readPosition + 1 >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition + 1]
}

/*
	检查当前正在查看的字符，根据字符返回相应的词法单元。
*/
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	/* 运算符的处理 */
	// = | == | ===
	case '=':
		if l.peekChar() == '=' {
			if l.peekNextChar() == '=' {
				// 处理 严格相等 
				ch := l.ch
				l.readChar()
				l.readChar()
				literal := string(ch) + "=="
				tok = token.Token{Type: token.EQ, Literal: literal}
			} else {
			  // 处理 == 
				ch := l.ch
				l.readChar()
				literal := string(ch) + string(l.ch)
				tok = token.Token{Type: token.EQ, Literal: literal}
			}
		} else {
			// 处理 模糊相等
			tok = newToken(token.ASSIGN, l.ch)
		}
	// ! | != | !==
	case '!':
		if l.peekChar() == '=' {
			if l.peekNextChar() == '=' {
				// 处理 !==
				ch := l.ch
				// 读取后面的两个字符
				l.readChar()
				l.readChar()
				literal := string(ch) + "!=" // 即 !==
				tok = token.Token{Type: token.NOT_EQ, Literal: literal}
			} else {
				// 处理 != 
				ch := l.ch
				// 读取下一个字符
				l.readChar()
				literal := string(ch) + string(l.ch) // 即 !=
				tok = token.Token{Type: token.NOT_EQ, Literal: literal}
			}
		} else {
			// 否则，返回BANG
			tok = newToken(token.BANG, l.ch)
		}
	// 处理 + | += | ++
	case '+':
		if l.peekChar() == '=' {
			// 处理 +=
			ch := l.ch
			l.readChar()
			literal := string(ch) + "="
			tok = token.Token{Type: token.PLUS_EQ, Literal: literal}
		} else if l.peekChar() == '+' {
			// 处理 ++
			ch := l.ch
			l.readChar()
			literal := string(ch) + "+"
			tok = token.Token{Type: token.INCREMENT, Literal: literal}
		} else {
			// 否则，返回+
			tok = newToken(token.PLUS, l.ch)
		}
	// 处理 - | -= | -- | ->
	case '-':
		if l.peekChar() == '=' {
			// 处理 -=
			ch := l.ch
			l.readChar()
			literal := string(ch) + "="
			tok = token.Token{Type: token.MINUS_EQ, Literal: literal}
		} else if l.peekChar() == '-' {
			// 处理 --
			ch := l.ch
			l.readChar()
			literal := string(ch) + "-"
			tok = token.Token{Type: token.DECREMENT, Literal: literal}
		} else if l.peekChar() == '>' {
			// 处理 ->
			ch := l.ch
			l.readChar()
			literal := string(ch) + ">"
			tok = token.Token{Type: token.ARROW, Literal: literal}
		} else {
			// 否则，返回-
			tok = newToken(token.MINUS, l.ch)
		}
	// 处理 * | *=
	case '*':
		if l.peekChar() == '=' {
			// 处理 *=
			ch := l.ch
			l.readChar()
			literal := string(ch) + "="
			tok = token.Token{Type: token.ASTERISK_EQ, Literal: literal}
		} else {
			// 否则，返回*
			tok = newToken(token.ASTERISK, l.ch)
		}
	// 处理 / | /= | //
	case '/':
		if l.peekChar() == '/' {
			// 处理 //
			ch := l.ch
			l.readChar()
			literal := string(ch) + "/"
			tok = token.Token{Type: token.COMMENT, Literal: literal}
		} else if l.peekChar() == '=' {
			// 处理 /=
			ch := l.ch
			l.readChar()
			literal := string(ch) + "="
			tok = token.Token{Type: token.SLASH_EQ, Literal: literal}
		} else {
			// 否则，返回/
			tok = newToken(token.SLASH, l.ch)
		}
	// 处理 % | %=
	case '%':
		if l.peekChar() == '=' {
			// 处理 %=
			ch := l.ch
			l.readChar()
			literal := string(ch) + "="
			tok = token.Token{Type: token.MODULO_EQ, Literal: literal}
		} else {
			// 否则，返回%
			tok = newToken(token.MODULO, l.ch)
		}
	// 逻辑运算符
	// && 
	case '&':
		if l.peekChar() == '&' {
			// 处理 &&
			ch := l.ch
			l.readChar()
			literal := string(ch) + "&"
			tok = token.Token{Type: token.AND, Literal: literal}
		} else {
			// 否则，返回&
			tok = newToken(token.BIT_AND, l.ch)
		}
	// || 
	case '|':
		if l.peekChar() == '|' {
			// 处理 ||
			ch := l.ch
			l.readChar()
			literal := string(ch) + "|"
			tok = token.Token{Type: token.OR, Literal: literal}
		} else {
			// 否则，返回|
			tok = newToken(token.BIT_OR, l.ch)
		}
	// 位运算符
	// 处理 ^
	case '^':
		tok = newToken(token.BIT_XOR, l.ch)
	// 处理 ~
	case '~':
		tok = newToken(token.BIT_NOT, l.ch)
	// 处理 < | << | <=
	case '<':
		if l.peekChar() == '<' {	
			// 处理 <<
			ch := l.ch
			l.readChar()
			literal := string(ch) + "<"
			tok = token.Token{Type: token.BIT_SHIFT_LEFT, Literal: literal}
		} else if l.peekChar() == '=' {
			// 处理 <=
			ch := l.ch
			l.readChar()
			literal := string(ch) + "="
			tok = token.Token{Type: token.LTE, Literal: literal}
		} else {
			// 否则，返回<
			tok = newToken(token.LT, l.ch)
		}
	// 处理 > | >> | >=
	case '>':
		if l.peekChar() == '>' {
			// 处理 >>
			ch := l.ch
			l.readChar()
			literal := string(ch) + ">"
			tok = token.Token{Type: token.BIT_SHIFT_RIGHT, Literal: literal}
		} else if l.peekChar() == '=' {
			// 处理 >=
			ch := l.ch
			l.readChar()
			literal := string(ch) + "="
			tok = token.Token{Type: token.GTE, Literal: literal}
		} else {
			// 否则，返回>
			tok = newToken(token.GT, l.ch)
		}
	// 处理操作符
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case ':':
		tok = newToken(token.COLON, l.ch)
	// . | ...
	case '.':
		// tok = newToken(token.DOT, l.ch)
		if l.peekChar() == '.' && l.peekNextChar() == '.' {
			// 处理 ...
			ch := l.ch
			l.readChar()
			l.readChar()
			literal := string(ch) + ".."
			tok = token.Token{Type: token.SPREAD, Literal: literal}
		} else {
			// 否则，返回.
			tok = newToken(token.DOT, l.ch)
		}
	// ? | ?. | ??
	case '?':
		if l.peekChar() == '.' {
			// 处理 ?.
			ch := l.ch
			l.readChar()
			literal := string(ch) + "."
			tok = token.Token{Type: token.OPTIONAL_CHAIN, Literal: literal}
		} else if l.peekChar() == '?' {
			// 处理 ??
			ch := l.ch
			l.readChar()
			literal := string(ch) + "?"
			tok = token.Token{Type: token.NULLISH, Literal: literal}
		} else {
			// 否则，返回?
			tok = newToken(token.QUESTION, l.ch)
		}
	// 括号
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '[':
		tok = newToken(token.LBRACKET, l.ch)
	case ']':
		tok = newToken(token.RBRACKET, l.ch)
	// 处理字符串
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
		return tok
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.NUMBER
			tok.Literal = l.readNumber()
			return tok
		} else if l.ch == 0 {
			tok.Type = token.EOF
			tok.Literal = ""
		} else {
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

/*
	读取一个标识符
*/
func (l *Lexer) readIdentifier() string {
	position := l.position

	if l.ch == '_' {
		l.readChar()
		if l.ch == '_' {
			// 处理 __xxxxx
			return l.readSpecialIdentifier(position)
		}
	}

	// 普通标识符处理
	for isLetter(l.ch) || isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

/*
	读取一个特殊标识符
*/
func (l *Lexer) readSpecialIdentifier(startPosition int) string {
	for isLetter(l.ch) || l.ch == '_' || isDigit(l.ch) {
		l.readChar()
	}
	word := l.input[startPosition:l.position]

	switch word {
	case "__proto__":
		return word
	default:
		return word
	}	
}
