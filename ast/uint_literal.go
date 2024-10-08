package ast

import (
	"fmt"

	"github.com/HannesKimara/cddlc/token"
)

// IntegerLiteral represents the AST Node for an integer literal i.e 3
type UintLiteral struct {
	Pos     token.Position
	Token   token.Token
	Literal uint64
}

func (ul *UintLiteral) Start() token.Position {
	return ul.Pos
}

func (ul *UintLiteral) End() token.Position {
	return ul.Pos.To(len(fmt.Sprintf("%d", ul.Literal)))
}

func (ul *UintLiteral) String() string {
	s := fmt.Sprintf("%s - %s", ul.Start(), ul.End())

	return s
}
