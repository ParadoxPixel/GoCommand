package components

type Syntax struct {
	Usage string
	Arguments []Argument
}

func (s *Syntax) AddArgument(arg Argument) *Syntax {
	s.Arguments = append(s.Arguments, arg)
	return s
}