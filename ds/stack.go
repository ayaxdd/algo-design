package ds

type Stack struct {
	elems []any
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(item any) {
	s.elems = append(s.elems, item)
}

func (s *Stack) Pop() (any, bool) {
	if s.IsEmpty() {
		return nil, false
	}

	lastIndex := len(s.elems) - 1
	item := s.elems[lastIndex]
	s.elems = s.elems[:lastIndex]

	return item, true
}

func (s *Stack) Peek() (any, bool) {
	if s.IsEmpty() {
		return nil, false
	}

	lastIndex := len(s.elems) - 1
	item := s.elems[lastIndex]

	return item, true
}

func (s *Stack) IsEmpty() bool {
	return len(s.elems) == 0
}
