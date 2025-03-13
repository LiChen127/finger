package ast

import (
	"bytes"
	"finger/token"
	"strings"
)

/**
let <标识符> = <表达式>;
*/

/*
	抽象语法树节点
*/
type Node interface {
	TokenLiteral() string
	String() string
}

/*
	节点
*/
type Statement interface {
	Node
	statementNode()
}

/*
	表达式
*/
type Expression interface {
	Node
	expressionNode()
}

/*
	程序
*/
type Program struct {
	// 切片，存储多个实现Statements的AST节点
	Statements []Statement
}

/*
	获取程序的第一个语句的token
*/
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}

	return ""
}

/*
	返回程序的string表示
*/
func (p *Program) String() string {
	// 创建缓冲区
	var out bytes.Buffer

	for _, s := range p.Statements {
		// 将每条语句的string表示写入缓冲区
		out.WriteString(s.String())
	}

	return out.String()
}

type LetStatement struct {
	Token token.Token // token.LET词法单元
	Name *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())	
	}

	out.WriteString(";")

	return out.String()
}


type ReturnStatement struct {
	Token token.Token // token.RETURN词法单元
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	
	out.WriteString(";")

	return out.String()
}

type Identifier struct {
	Token token.Token // token.IDENT词法单元
	Value string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) String() string {
	return i.Value
}

/*
	表达式语句
*/
type ExpressionStatement struct {
	Token token.Token // 该表达式中的第一个词法单元
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}

func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil { 
		return es.Expression.String()
	}

	return ""
}

type IntegerLiteral struct {
	Token token.Token // token.INT词法单元
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}

func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}

func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}
	

/*
	前缀表达式
*/
type PrefixExpression struct {
	Token token.Token // token.BANG or token.MINUS
	Operator string
	Right Expression
}

func (pe *PrefixExpression) expressionNode() {}

func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")
	
	return out.String()
}


/*
	中缀表达式
*/
type InfixExpression struct {
	Token token.Token // token.PLUS or token.MINUS or token.SLASH or token.ASTERISK
	Left Expression
	Operator string
	Right Expression
}

func (ie *InfixExpression) expressionNode() {}

func (ie *InfixExpression) TokenLiteral() string {
	return ie.Token.Literal
}

func (ie *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()
}

/*
	布尔字面量
*/
type Boolean struct {
	Token token.Token // token.TRUE or token.FALSE
	Value bool
}

func (b *Boolean) expressionNode() {}

func (b *Boolean) TokenLiteral() string {
	return b.Token.Literal
}

func (b *Boolean) String() string {
	return b.Token.Literal
}


/*
	if语句
*/
type IfExpression struct {
	Token token.Token // token.IF词法单元
	Condition Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode() {}

func (ie *IfExpression) TokenLiteral() string {
	return ie.Token.Literal
}

func (ie *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ie.Alternative.String())
	}

	return out.String()
}


/*
	块语句
*/
type BlockStatement struct {
	Token token.Token // token.LBRACE词法单元
	Statements []Statement
}

func (bs *BlockStatement) statementNode() {}

func (bs *BlockStatement) TokenLiteral() string {
	return bs.Token.Literal
}

func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

/*
	函数声明
*/
type FunctionLiteral struct {
	Token token.Token // token.FUNCTION词法单元
	Parameters []*Identifier
	Body *BlockStatement
}

func (fl *FunctionLiteral) expressionNode() {}

func (fl *FunctionLiteral) TokenLiteral() string {
	return fl.Token.Literal
}

func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}

	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(fl.Body.String())
	out.WriteString("\n}")

	return out.String()
}

/*
	调用表达式
*/

type CallExpression struct {
	Token token.Token // token.LPAREN词法单元
	Function Expression
	Arguments []Expression
}

func (ce *CallExpression) expressionNode() {}

func (ce *CallExpression) TokenLiteral() string {
	return ce.Token.Literal
}

func (ce *CallExpression) String() string {
	var out bytes.Buffer

	args := []string{}

	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}