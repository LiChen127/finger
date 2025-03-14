package token

/* 关键字 */
var keywords = map[string]TokenType{
	// 变量声明
	"let":   LET,
	"const": CONST,

	// 函数声明
	"fn":    FUNCTION,
	"return": RETURN,

	// 控制流
	"for": FOR,
	"while": WHILE,
	"do":      DO,
	"break":   BREAK,
	"continue": CONTINUE,
	"switch":  SWITCH,
	"case":    CASE,
	"default": DEFAULT,
	"if":      IF,
	"else":    ELSE,

	// 异常处理
	"try":    TRY,
	"catch":  CATCH,
	"finally": FINALLY,
	"throw":  THROW,

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
	"async":  ASYNC,
	"await":  AWAIT,
	"Promise": PROMISE,
	"then":   THEN,

	// 内置函数
	"print":  PRINT,
	"len":    LEN,

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
