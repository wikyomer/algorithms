package calc

import "algorithms/structs"

func evaluate(rpn TokenQueue) (result int) {
	calcStack := structs.Stack[Token]{}

	for !rpn.IsEmpty() {
		if rpn.Pick().tokenType == NUM {
			calcStack.Push(rpn.Pop())
		} else {
			operand2 := calcStack.Pop()
			operand1 := calcStack.Pop()

			switch rpn.Pick().op {
			case ADD:
				calcStack.Push(Token{tokenType: NUM, value: operand1.value + operand2.value})
			case SUB:
				calcStack.Push(Token{tokenType: NUM, value: operand1.value - operand2.value})
			case MUL:
				calcStack.Push(Token{tokenType: NUM, value: operand1.value * operand2.value})
			case DIV:
				calcStack.Push(Token{tokenType: NUM, value: operand1.value / operand2.value})
			}

			rpn.Pop()
		}
	}

	result = calcStack.Pop().value
	return result
}
