package calc

import (
	"algorithms/structs"
)

type TokenType string

const (
	OP  TokenType = "OP"
	NUM TokenType = "NUM"
)

type OpType string

const (
	ADD OpType = "ADD"
	SUB OpType = "SUB"
	MUL OpType = "MUL"
	DIV OpType = "DIV"
)

type Token struct {
	tokenType TokenType
	value     int
	op        OpType
}

type TokenQueue = structs.Queue[Token]

func Calculation(input string) (result int) {
	infix := parse(input)
	rpn := shuntingYard(infix)
	result = evaluate(rpn)
	return result
}
