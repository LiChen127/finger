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
	IDENT   = "IDENT"   // 标识符

	/* 运算符 */

	// 赋值运算符
	ASSIGN      = "="
	PLUS_EQ     = "+="
	MINUS_EQ    = "-="
	ASTERISK_EQ = "*="
	SLASH_EQ    = "/="
	MODULO_EQ   = "%="

	// 算术运算符
	PLUS      = "+"
	MINUS     = "-"
	ASTERISK  = "*"
	SLASH     = "/"
	MODULO    = "%"
	INCREMENT = "++"
	DECREMENT = "--"

	// 逻辑运算符
	BANG          = "!"
	AND           = "&&"
	OR            = "||"
	EQ            = "=="
	NOT_EQ        = "!="
	STRICT_EQ     = "===" // 严格相等
	STRICT_NOT_EQ = "!==" // 严格不相等

	// 位运算符
	BIT_AND         = "&"
	BIT_OR          = "|"
	BIT_XOR         = "^"
	BIT_NOT         = "~"
	BIT_SHIFT_LEFT  = "<<"
	BIT_SHIFT_RIGHT = ">>"

	// 比较运算符
	LT  = "<"
	GT  = ">"
	LTE = "<="
	GTE = ">="

	// 操作符
	COMMA          = ","
	SEMICOLON      = ";"
	COLON          = ":"
	DOT            = "."
	QUESTION       = "?"
	NULLISH        = "??"
	SPREAD         = "..."
	OPTIONAL_CHAIN = "?."

	// 括号
	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"

	// 函数式编程关键字
	FUNCTION = "fn"
	ARROW    = "->"
	RETURN   = "return"

	// 变量声明
	LET   = "let"
	CONST = "const"

	// 基础值
	TRUE      = "true"
	FALSE     = "false"
	NULL      = "null"
	UNDEFINED = "undefined"

	// 字面量
	STRING  = "string"
	NUMBER  = "number"
	BOOLEAN = "boolean"
	ARRAY   = "array"
	OBJECT  = "object"
	REGEX   = "regex"
	DATE    = "date"
	MAP     = "Map"
	SET     = "Set"

	// 控制流支持
	FOR      = "for"
	WHILE    = "while"
	DO       = "do"
	BREAK    = "break"
	CONTINUE = "continue"
	SWITCH   = "switch"
	CASE     = "case"
	DEFAULT  = "default"
	IF       = "if"
	ELSE     = "else"

	// 函数式特性
	MAPFn   = "map"
	REDUCE  = "reduce"
	FILTER  = "filter"
	FOREACH = "foreach"
	CONCAT  = "concat"
	SLICE   = "slice"
	SPLIT   = "split"
	JOIN    = "join"

	// 原型系统
	PROTO  = "__proto__"
	CREATE = "create"
	TYPEOF = "typeof"
	IN     = "in"

	// 异步支持
	ASYNC   = "async"
	AWAIT   = "await"
	PROMISE = "Promise"
	THEN    = "then"

	// 模块系统
	IMPORT = "import"
	EXPORT = "export"
	FROM   = "from"

	// 错误处理
	TRY     = "try"
	CATCH   = "catch"
	FINALLY = "finally"
	THROW   = "throw"

	// 内置函数
	PRINT = "print"
	LEN   = "len"

	// 注释
	COMMENT = "//"
	// @todo: 支持多行注释
	// MULTI_LINE  = "/*"
	// MULTI_LINE2 = "*/"
	// DOC_COMMENT = "/**"
	// DOC_LINE    = "*"
	// DOC_LINE2   = "*/"
)

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

/* 关键字 */
var keywords = map[string]TokenType{
	// 变量声明
	"let":   LET,
	"const": CONST,

	// 函数声明
	"fn":     FUNCTION,
	"return": RETURN,

	// 控制流
	"for":      FOR,
	"while":    WHILE,
	"do":       DO,
	"break":    BREAK,
	"continue": CONTINUE,
	"switch":   SWITCH,
	"case":     CASE,
	"default":  DEFAULT,
	"if":       IF,
	"else":     ELSE,

	// 异常处理
	"try":     TRY,
	"catch":   CATCH,
	"finally": FINALLY,
	"throw":   THROW,

	// 模块系统
	"import": IMPORT,
	"export": EXPORT,
	"from":   FROM,

	// 原型系统
	"typeof":    TYPEOF,
	"__proto__": PROTO,
	"create":    CREATE,
	"in":        IN,

	// 异步支持
	"async":   ASYNC,
	"await":   AWAIT,
	"Promise": PROMISE,
	"then":    THEN,

	// 内置函数
	"print": PRINT,
	"len":   LEN,

	// 基础值
	"true":      TRUE,
	"false":     FALSE,
	"null":      NULL,
	"undefined": UNDEFINED,
	"boolean":   BOOLEAN,
	"number":    NUMBER,
	"string":    STRING,
	"array":     ARRAY,
	"object":    OBJECT,
	"regex":     REGEX,
	"date":      DATE,
	"Map":       MAP,
	"Set":       SET,

	// 函数式
	"map":     MAPFn,
	"reduce":  REDUCE,
	"filter":  FILTER,
	"foreach": FOREACH,
	"concat":  CONCAT,
	"slice":   SLICE,
	"split":   SPLIT,
	"join":    JOIN,
}
