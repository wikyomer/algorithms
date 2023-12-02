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
				if GetPriority(string(char)) <= GetPriority(RPnotation.Pick()) {
					for !RPnotation.IsEmpty() {
						expression = append(expression, RPnotation.Pop())
					}
				}
				RPnotation.Push(string(char))
			}
		}
	}
	for !RPnotation.IsEmpty() {
		expression = append(expression, RPnotation.Pop())
	}

	//for _, char := range expression {
	//	if unicode.IsDigit(rune(char)) {
	//		calcStack.Push(char)
	//	} else {
	//		op2 := calcStack.Pop()
	//		op1 := calcStack.Pop()
	//		switch char {
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
		fallthrough
	case "-":
		return 1
	case "*":
		fallthrough
	case "/":
		return 2
	}
	return -1
}
