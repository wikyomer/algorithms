package main

import (
	"algorithms/calc"
	"fmt"
)

func main() {
	fmt.Print(calc.Calculation("23+14*4/8-5"))
	//fmt.Print(calc.Calculation("23+14*456/98*7*64/5-1"))
	//expression, err := bufio.NewReader(os.Stdin).ReadString('\n')
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Print(calc.Calculation(expression))
}
