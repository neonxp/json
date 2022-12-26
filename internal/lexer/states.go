package lexer

func InitJson(l *Lexer) stateFunc {
	ignoreWhiteSpace(l)
	switch {
	case l.Accept("{"):
		l.Emit(LObjectStart)
		return stateInObject
	case l.Accept("["):
		l.Emit(LArrayStart)
	case l.Peek() == eof:
		return nil
	}
	return l.Errorf("Unknown token: %s", string(l.Peek()))
}

func stateInObject(l *Lexer) stateFunc {
	// we in object, so we expect field keys and values
	ignoreWhiteSpace(l)
	if l.Accept("}") {
		l.Emit(LObjectEnd)
		// If meet close object return to previous state (including initial)
		return l.PopState()
	}
	ignoreWhiteSpace(l)
	l.Accept(",")
	ignoreWhiteSpace(l)
	if !scanQuotedString(l, '"') {
		return l.Errorf("Unknown token: %s", string(l.Peek()))
	}
	l.Emit(LObjectKey)
	ignoreWhiteSpace(l)
	if !l.Accept(":") {
		return l.Errorf("Expected ':'")
	}
	ignoreWhiteSpace(l)
	l.Emit(LObjectValue)
	switch {
	case scanQuotedString(l, '"'):
		l.Emit(LString)
		ignoreWhiteSpace(l)
		l.Accept(",")
		l.Ignore()
		ignoreWhiteSpace(l)
		return stateInObject
	case scanNumber(l):
		l.Emit(LNumber)
		ignoreWhiteSpace(l)
		l.Accept(",")
		l.Ignore()
		ignoreWhiteSpace(l)
		return stateInObject
	case l.AcceptAnyOf([]string{"true", "false"}, true):
		l.Emit(LBoolean)
		ignoreWhiteSpace(l)
		l.Accept(",")
		l.Ignore()
		ignoreWhiteSpace(l)
		return stateInObject
	case l.AcceptString("null", true):
		l.Emit(LNull)
		ignoreWhiteSpace(l)
		l.Accept(",")
		l.Ignore()
		ignoreWhiteSpace(l)
		return stateInObject
	case l.Accept("{"):
		l.Emit(LObjectStart)
		l.PushState(stateInObject)
		return stateInObject
	case l.Accept("["):
		l.Emit(LArrayStart)
		l.PushState(stateInObject)
		return stateInArray
	}
	return l.Errorf("Unknown token: %s", string(l.Peek()))
}

func stateInArray(l *Lexer) stateFunc {
	ignoreWhiteSpace(l)
	l.Accept(",")
	ignoreWhiteSpace(l)
	switch {
	case scanQuotedString(l, '"'):
		l.Emit(LString)
	case scanNumber(l):
		l.Emit(LNumber)
	case l.AcceptAnyOf([]string{"true", "false"}, true):
		l.Emit(LBoolean)
	case l.AcceptString("null", true):
		l.Emit(LNull)
	case l.Accept("{"):
		l.Emit(LObjectStart)
		l.PushState(stateInArray)
		return stateInObject
	case l.Accept("["):
		l.Emit(LArrayStart)
		l.PushState(stateInArray)
		return stateInArray
	case l.Accept("]"):
		l.Emit(LArrayEnd)
		return l.PopState()
	}
	return stateInArray
}

func ignoreWhiteSpace(l *Lexer) {
	l.AcceptWhile(" \n\t") // ignore whitespaces
	l.Ignore()
}
