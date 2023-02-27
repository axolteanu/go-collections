package stack

type Node struct {
	value any
	next  *Node
}

type Stack struct {
	head *Node
}

func New() *Stack {
	return new(Stack)
}

func (s *Stack) Push(value any) {
	node := Node{value, nil}
	if s.head != nil {
		node.next = s.head
	}
	s.head = &node
}

func (s *Stack) Pop() any {
	if s.head == nil {
		return nil
	}
	ret := s.head.value
	s.head = s.head.next
	return ret
}

func (s *Stack) Peek() any {
	if s.IsEmpty() {
		return nil
	} else {
		return s.head.value
	}
}

func (s *Stack) IsEmpty() bool {
	return s.head == nil
}
