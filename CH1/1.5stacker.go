package CH1

import (
	"errors"
	"log"
)

type Stack []interface{}

func (stack Stack) Len() int {
	return len(stack)
}
func (stack Stack) Cap() int {
	return cap(stack)
}
func (stack *Stack) Push(x interface{}) {
	*stack = append(*stack, x)
	log.Printf("入栈%v", x)
}
func (stack Stack) Top() (interface{}, error) {
	if len(stack) == 0 {
		return nil, errors.New("空栈")

	}
	return stack[len(stack)-1], nil
}
func (stack *Stack) Pop() (interface{}, error) {
	theStack := *stack
	if len(theStack) == 0 {
		return nil, errors.New("空栈")
	}
	x := theStack[len(theStack)-1]
	*stack = theStack[:len(theStack)-1]
	return x, nil

}
