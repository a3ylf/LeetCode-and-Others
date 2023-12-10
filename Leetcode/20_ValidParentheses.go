package Leetcode

type Stack []interface{}

func (s *Stack) Push(item interface{}) {
	*s = append(*s, item)
}

func (s *Stack) Pop() interface{} {
	if len(*s) == 0 {
		return nil
	}

	index := len(*s) - 1
	item := (*s)[index]
	*s = (*s)[:index]
	return item
}

func (s *Stack) Peek() interface{} {
	if len(*s) == 0 {
		return nil
	}

	return (*s)[len(*s)-1]
}

func isValid(s string) bool {
	stack := Stack{}
	for _, x := range s {
		if x == '[' {
			stack.Push(']')
		} else if x == '{' {
			stack.Push('}')
		} else if x == '(' {
			stack.Push(')')
		} else if stack.Peek() == x {
			stack.Pop()
		} else {
			return false
		}
	}
	return stack.Peek() == nil
}
