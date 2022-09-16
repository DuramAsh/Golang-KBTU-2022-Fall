package main

import "fmt"

type DS interface {
	peek() (int, error)
	pop() error
	isEmpty() bool
	print()
	push(val int)
}

type Incrementer interface {
	increment(steps int)
	getStack() *Stack
}

type Node struct {
	Value int
	Next  *Node
}

type Stack struct {
	top *Node
}

type IncrementStack struct {
	stack *Stack
}

func (is *IncrementStack) getStack() *Stack {
	return is.stack
}

func (is *IncrementStack) increment(steps int) {
	temp := is.stack.top
	for i := 0; i < steps; i++ {
		if temp == nil {
			break
		}
		temp.Value++
		temp = temp.Next
	}
}

func main() {
	var st Incrementer = &IncrementStack{
		stack: &Stack{},
	}
	st.getStack().push(1)
	st.getStack().push(2)
	st.getStack().push(3)
	st.getStack().push(4)
	st.getStack().push(5)
	st.increment(2)
	for i := 0; i < 6; i++ {
		val, err := st.getStack().peek()
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(val)
		}
		err = st.getStack().pop()
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	st.getStack().print()
}

func (s *Stack) push(val int) {
	newNode := &Node{Value: val}
	newNode.Next = s.top
	s.top = newNode
}

func (s *Stack) isEmpty() bool {
	return s.top == nil
}

func (s *Stack) pop() error {
	if s.isEmpty() {
		return fmt.Errorf("error: the stack is empty")
	}
	s.top = s.top.Next
	return nil
}

func (s *Stack) peek() (int, error) {
	if s.isEmpty() {
		return 0, fmt.Errorf("error: the stack is empty")
	}
	return s.top.Value, nil
}

func (s *Stack) print() {
	if s.isEmpty() {
		fmt.Println("stack is empty, nothing to print")
		return
	}
	temp := s.top
	for temp != nil {
		fmt.Print(temp.Value, " ")
		temp = temp.Next
	}
	fmt.Println()
}
