package main

import (
	"algorithms/structs"
	"unicode"
)

func calculation(infixNotation string) (result int) {
	RPnotation := structs.Stack[string]{}
	var expression []string

	for _, char := range infixNotation {
		if unicode.IsDigit(char) {
			expression = append(expression, string(char))
		} else {
			if RPnotation.IsEmpty() {
				RPnotation.Push(string(char))
			} else {
				if GetPriority(string(char)) >= GetPriority(RPnotation.Pick()) {
					expression = append(expression, RPnotation.Pop())
				}
				RPnotation.Push(string(char))
			}
		}
	}
	for !RPnotation.IsEmpty() {
		expression = append(expression, RPnotation.Pop())
	}

	//for i := range expression {
	//	if unicode.IsDigit(rune(expression[i])) {
	//		calcStack.Push(expression[i])
	//	} else {
	//		op2 := calcStack.Pop()
	//		op1 := calcStack.Pop()
	//		switch expression[i] {
	//		case '+':
	//			interRes := op1 + op2
	//			calcStack.Push(interRes)
	//		case '-':
	//			interRes := op1 - op2
	//			calcStack.Push(interRes)
	//		case '*':
	//			interRes := op1 * op2
	//			calcStack.Push(interRes)
	//		case '/':
	//			interRes := op1 / op2
	//			calcStack.Push(interRes)
	//		}
	//	}
	//}

	//result = calcStack.Pop()
	return result
}

func GetPriority(operation string) (priority int) {
	switch operation {
	case "+":
	case "-":
		return 1
	case "*":
	case "/":
		return 2
	}
	return -1
}
