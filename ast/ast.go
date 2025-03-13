package ast

import (
	"bytes"
	"finger/token"
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