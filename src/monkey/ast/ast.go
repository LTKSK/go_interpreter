package ast

import "monkey/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

type Identifier struct {
	Token token.Token
	Value string
}

// let <ident> = <expr>;
type LetStatement struct {
	Token token.Token
	Name  *Identifier // let x = 5 の x
	Value Expression  // let x = 5 の 5
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// return <expr>;
type ReturnStatement struct {
	Token       token.Token // the 'return' token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
