package calc

import (
	"strconv"
	"unicode"
)

func parse(input string) (infix TokenQueue) {
	var buffer string

	for _, elem := range input {
		if unicode.IsDigit(elem) {
			buffer += string(elem)
		} else {
			infix.Push(tokenFromBuffer(&buffer))
			infix.Push(parseOp(elem))
		}
	}
	infix.Push(tokenFromBuffer(&buffer))

	//infix.Print()

	return infix
}

func tokenFromBuffer(buffer *string) Token {
	value, _ := strconv.Atoi(*buffer)
	token := Token{tokenType: NUM, value: value}
	*buffer = ""
	return token
}

func parseOp(char rune) Token {
	token := Token{tokenType: OP}
	switch char {
	case '+':
		token.op = ADD
	case '-':
		token.op = SUB
	case '*':
		token.op = MUL
	case '/':
		token.op = DIV

	}
	return token
}
