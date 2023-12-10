package calc

import "algorithms/structs"

func shuntingYard(infix TokenQueue) (rpn TokenQueue) {
	sortingStack := structs.Stack[Token]{}

	for !infix.IsEmpty() {
		if infix.Pick().tokenType == NUM {
			rpn.Push(infix.Pop())
		} else {
			if !sortingStack.IsEmpty() {
				if GetPriority(infix.Pick()) <= GetPriority(sortingStack.Pick()) {
					if GetPriority(infix.Pick()) == GetPriority(sortingStack.Pick()) {
						rpn.Push(sortingStack.Pop())
					} else {
						for !sortingStack.IsEmpty() {
							rpn.Push(sortingStack.Pop())
						}
					}
				}
			}
			sortingStack.Push(infix.Pop())
		}
	}
	for !sortingStack.IsEmpty() {
		rpn.Push(sortingStack.Pop())
	}

	rpn.Print()

	return rpn
}

func GetPriority(operation Token) (priority int) {
	switch operation.op {
	case ADD:
		fallthrough
	case SUB:
		return 1
	case MUL:
		fallthrough
	case DIV:
		return 2
	}
	return -1
}
