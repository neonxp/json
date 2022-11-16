package parser

type stateFunc func(*lexer) stateFunc

type stateStack []stateFunc

func (ss *stateStack) Push(s stateFunc) {
	*ss = append(*ss, s)
}

func (ss *stateStack) Pop() (s stateFunc) {
	if len(*ss) == 0 {
		return nil
	}
	*ss, s = (*ss)[:len(*ss)-1], (*ss)[len(*ss)-1]
	return s
}
