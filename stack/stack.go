package stack

type Node struct {
	value any
	next  *Node
}

type Stack struct {
	head *Node
}

func (s *Stack) Push(element any) {
	node := Node{element, nil}
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

func (s *Stack) IsEmpty() bool {
	return s.head == nil
}
