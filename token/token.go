package token

type TokenType string

type Token struct {
	Type    TokenType // 词法单元类型
	Literal string    // 词法单元字面量
}

/* 词法单元类型 */
const (
	ILLEGAL = "ILLEGAL" // 非法单元
	EOF     = "EOF"     // 文件结束

	// 标识符 + 字面量
	IDENT  = "IDENT"
	INT    = "INT"
	STRING = "STRING"

	// 运算符
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"
	EQ       = "=="
	NOT_EQ   = "!="

	// 分隔符
	COMMA     = ","
	SEMICOLON = ";"

	// 括号
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// 关键字
	LET      = "LET"
	FUNCTION = "FUNCTION"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"

	// 数组相关
	LBRACKET = "["
	RBRACKET = "]"

	// 哈希表相关
	COLON = ":"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

/*
检查字符串是否是关键字
*/
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		// 如果ident是关键字，则返回关键字
		return tok
	}
	// 否则，返回标识符
	return IDENT
}
