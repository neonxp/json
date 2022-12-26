package lexer

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

const eof rune = -1

type Lexem struct {
	Type  lexType // Type of Lexem.
	Value string  // Value of Lexem.
	Start int     // Start position at input string.
	End   int     // End position at input string.
}

//go:generate stringer -type=lexType
type lexType int

const (
	LEOF lexType = iota
	LError
	LObjectStart
	LObjectEnd
	LObjectKey
	LObjectValue
	LArrayStart
	LArrayEnd
	LString
	LNumber
	LBoolean
	LNull
)

// Lexer holds current scanner state.
type Lexer struct {
	Input  string     // Input string.
	Start  int        // Start position of current lexem.
	Pos    int        // Pos at input string.
	Output chan Lexem // Lexems channel.
	width  int        // Width of last rune.
	states stateStack // Stack of states to realize PrevState.
}

// newLexer returns new scanner for input string.
func NewLexer(input string) *Lexer {
	return &Lexer{
		Input:  input,
		Start:  0,
		Pos:    0,
		Output: make(chan Lexem, 2),
		width:  0,
	}
}

// Run lexing.
func (l *Lexer) Run(init stateFunc) {
	for state := init; state != nil; {
		state = state(l)
	}
	close(l.Output)
}

// PopState returns previous state function.
func (l *Lexer) PopState() stateFunc {
	return l.states.Pop()
}

// PushState pushes state before going deeper states.
func (l *Lexer) PushState(s stateFunc) {
	l.states.Push(s)
}

// Emit current lexem to output.
func (l *Lexer) Emit(typ lexType) {
	l.Output <- Lexem{
		Type:  typ,
		Value: l.Input[l.Start:l.Pos],
		Start: l.Start,
		End:   l.Pos,
	}
	l.Start = l.Pos
}

// Errorf produces error lexem and stops scanning.
func (l *Lexer) Errorf(format string, args ...interface{}) stateFunc {
	l.Output <- Lexem{
		Type:  LError,
		Value: fmt.Sprintf(format, args...),
		Start: l.Start,
		End:   l.Pos,
	}
	return nil
}

// Next rune from input.
func (l *Lexer) Next() (r rune) {
	if int(l.Pos) >= len(l.Input) {
		l.width = 0
		return eof
	}
	r, l.width = utf8.DecodeRuneInString(l.Input[l.Pos:])
	l.Pos += l.width
	return r
}

// Back move position to previos rune.
func (l *Lexer) Back() {
	l.Pos -= l.width
}

// Ignore previosly buffered text.
func (l *Lexer) Ignore() {
	l.Start = l.Pos
	l.width = 0
}

// Peek rune at current position without moving position.
func (l *Lexer) Peek() (r rune) {
	r = l.Next()
	l.Back()
	return r
}

// Accept any rune from valid string. Returns true if Next rune was in valid string.
func (l *Lexer) Accept(valid string) bool {
	if strings.ContainsRune(valid, l.Next()) {
		return true
	}
	l.Back()
	return false
}

// AcceptString returns true if given string was at position.
func (l *Lexer) AcceptString(s string, caseInsentive bool) bool {
	input := l.Input[l.Start:]
	if caseInsentive {
		input = strings.ToLower(input)
		s = strings.ToLower(s)
	}
	if strings.HasPrefix(input, s) {
		l.width = 0
		l.Pos += len(s)
		return true
	}
	return false
}

// AcceptAnyOf substrings. Retuns true if any of substrings was found.
func (l *Lexer) AcceptAnyOf(s []string, caseInsentive bool) bool {
	for _, substring := range s {
		if l.AcceptString(substring, caseInsentive) {
			return true
		}
	}
	return false
}

// AcceptWhile passing symbols from input while they at `valid` string.
func (l *Lexer) AcceptWhile(valid string) bool {
	isValid := false
	for l.Accept(valid) {
		isValid = true
	}
	return isValid
}

// AcceptWhileNot passing symbols from input while they NOT in `invalid` string.
func (l *Lexer) AcceptWhileNot(invalid string) bool {
	isValid := false
	for !strings.ContainsRune(invalid, l.Next()) {
		isValid = true
	}
	l.Back()
	return isValid
}

// AtStart returns true if current lexem not empty
func (l *Lexer) AtStart() bool {
	return l.Pos == l.Start
}
