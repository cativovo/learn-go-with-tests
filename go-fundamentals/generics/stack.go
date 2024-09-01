package generics

type node[T any] struct {
	value T
	prev  *node[T]
}

type Stack[T any] struct {
	head   *node[T]
	length int
}

func (s *Stack[T]) Push(v T) {
	n := node[T]{
		value: v,
	}
	s.length++
	if s.head == nil {
		s.head = &n
		return
	}

	n.prev = s.head
	s.head = &n
}

func (s *Stack[T]) Pop() (T, bool) {
	var defaultV T
	if s.head == nil {
		return defaultV, false
	}

	n := s.head
	s.head = n.prev
	s.length--

	return n.value, true
}

func (s *Stack[T]) IsEmpty() bool {
	return s.length == 0
}

func (s *Stack[T]) Length() int {
	return s.length
}
