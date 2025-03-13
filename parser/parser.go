package parser

/*
	语法分析器
	采用自上而下的递归下降分析法
*/

import (
	"finger/ast"
	"finger/lexer"
	"finger/token"
	"fmt"
	"strconv"
)

const (
	_ int = iota
	LOWEST
	EQUALS // ==
	LESSGREATER // > or <
	SUM // +
	PRODUCT // *
	PREFIX // -X or !X
	CALL // myFunction(X)
)


// 优先级表
var precedences = map[token.TokenType]int {
	token.EQ: EQUALS,
	token.NOT_EQ: EQUALS,
	token.LT: LESSGREATER,
	token.GT: LESSGREATER,
	token.PLUS: SUM,
	token.MINUS: SUM,
	token.SLASH: PRODUCT,
	token.ASTERISK: PRODUCT,	
}

type Parser struct {
	l *lexer.Lexer // 指向词法分析器示例的指针

	curToken token.Token // 类似词法分析中的position, 指向当前正在解析的词法单元
	peekToken token.Token // 类似词法分析中的readPosition, 指向当前正在解析的词法单元的下一个词法单元
	errors []string // 错误信息, 是切片，每个错误语句都报错，而不是遇到一个错误就退出

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns map[token.TokenType]infixParseFn
}

/*
	普拉特语法分析器: 自上而下的递归下降分析法
	主要思想: 将解析函数(语义代码)与词法单元类型相关联。
	每当遇到某个词法单元类型时，调用相关联的解析函数来解析对应的表达式，最后返回生成的AST节点
	每个词法单元类型最多可以关联两个解析函数，这取决于词法单元的位置，是位于前缀位置还是中缀位置。
*/

// 定义两种类型的函数: 前缀和中缀解析函数
type (
	prefixParseFn func() ast.Expression
	infixParseFn func(ast.Expression) ast.Expression
)

func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

/*
	创建一个新的语法分析器
*/
func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l: l,
		errors: []string{},
	}

	p.nextToken()
	p.nextToken()

	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.registerPrefix(token.IDENT, p.parseIdentifier)
	p.registerPrefix(token.INT, p.parseIntegerLiteral)
	// 前缀表达式解析器
	p.registerPrefix(token.BANG, p.parsePrefixExpression)
	p.registerPrefix(token.MINUS, p.parsePrefixExpression)
	// 中缀表达式解析器
	p.registerInfix(token.PLUS, p.parseInfixExpression)
	p.registerInfix(token.MINUS, p.parseInfixExpression)
	p.registerInfix(token.SLASH, p.parseInfixExpression)
	p.registerInfix(token.ASTERISK, p.parseInfixExpression)
	p.registerInfix(token.EQ, p.parseInfixExpression)
	p.registerInfix(token.NOT_EQ, p.parseInfixExpression)
	p.registerInfix(token.LT, p.parseInfixExpression)
	p.registerInfix(token.GT, p.parseInfixExpression)


	return p
}

/*
	标识符解析器
*/
func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}


/*
	返回错误信息
*/
func (p *Parser) Errors() []string {
	return p.errors
}

/*
	添加错误信息
*/
func (p *Parser) peekErrors(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

/*
	获取下一个词法单元，前移curToken和peekToken
*/
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

/*
	检查peekToken是否是预期的词法单元
*/
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true 
	}
	p.peekErrors(t)
	return false 
}

/*
	检查curToken是否是预期的词法单元
*/
func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

/*
	检查peekToken是否是预期的词法单元
*/
func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

/*
	AST解析器
*/
func (p *Parser) parseProgram() *ast.Program {
	// 构造AST的根节点
	program := &ast.Program{}
	program.Statements = []ast.Statement{}
	// 遍历输入的词法单元, 直到遇见EOF
	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

/*
	语句解析器
*/
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
		case token.LET:
			return p.parseLetStatement()
		case token.RETURN:
			return p.parseReturnStatement()
		default:
			return p.parseExpressionStatement()
	}
}

/*
	let语句解析器
*/
func (p *Parser) parseLetStatement() *ast.LetStatement {
	// 创建一个let语句节点
	stmt := &ast.LetStatement{Token: p.curToken}

	// 判断下一个是不是期望的词法单元, 即标识符
	if !p.expectPeek(token.IDENT) {
		return nil
	}
	// 前移curToken, 并设置let语句节点的标识符
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	// 判断下一个是不是期望的词法单元, 即赋值符号
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// @todo: 跳过对表达式的处理, 直到遇见分号
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

/*
	return语句解析器
*/
func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	p.nextToken()

	// @todo: 跳过对表达式的处理，直到遇到分号
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

/*
	表达式语句解析器
*/
func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}
	
	stmt.Expression = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

/*
	表达式解析器
*/
func (p *Parser) parseExpression(precedence int) ast.Expression {
	// 获取当前词法单元的类型
	prefix := p.prefixParseFns[p.curToken.Type]

	if prefix == nil {
		p.noPrefixParseFnError(p.curToken.Type)
		return nil
	}

	leftExp := prefix()

	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}

		p.nextToken()
		leftExp = infix(leftExp)	
	}

	return leftExp	
}

func (p *Parser) noPrefixParseFnError(t token.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}


/*
	整数解析器
*/
func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.curToken}
	// 将字符串转换为int64
	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)

	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	lit.Value = value

	return lit
}

/*
	前缀表达式解析器
*/
func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token: p.curToken,
		Operator: p.curToken.Literal,
	}

	p.nextToken()

	expression.Right = p.parseExpression(PREFIX)

	return expression
}

/*
	中缀表达式解析器
*/
func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expression := &ast.InfixExpression{
		Token: p.curToken,
		Operator: p.curToken.Literal,
		Left: left,
	}

	precedence := p.curPrecedence()
	p.nextToken()
	expression.Right = p.parseExpression(precedence)

	return expression
}

/*
	根据peekToken的类型返回优先级
*/
func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}
	
	return LOWEST
}

/*
	根据curToken的类型返回优先级
*/
func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.curToken.Type]; ok {
		return p
	}

	return LOWEST
}
